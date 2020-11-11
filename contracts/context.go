package contracts

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/Decem-Technology/service-helper/bootstrap"
	"github.com/Decem-Technology/service-helper/helpers/wrappers"
	"github.com/Decem-Technology/service-helper/utils"
	"github.com/Decem-Technology/service-helper/validators"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v7"
	microError "github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/thoas/go-funk"
)

type (
	// AppContext helper context
	AppContext struct {
		context.Context
		claims *utils.CustomClaims
	}
)

const (
	BASIC_SCHEMA   string = "Basic "
	BEARER_SCHEMA  string = "Bearer "
	msAudienceName string = "micro_service"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("date", validators.DateValidation)
	validate.RegisterValidation("datetime", validators.DatetimeValidation)
	validate.RegisterValidation("date_range", validators.DateRangeValidation)
	validate.RegisterValidation("required_if", validators.RequiredIf)
	validate.RegisterValidation("required_if_not", validators.RequiredIfNotEqual)
}

// RegisterValidation register validation
func RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return validate.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

// Meta get headers
func (ctx *AppContext) Meta() metadata.Metadata {
	md, ok := metadata.FromContext(ctx.Context)
	if !ok {
		md = metadata.Metadata{}
	}
	return md
}

// Token get header access token
func (ctx *AppContext) Token() (string, error) {
	if token, ok := ctx.Meta()["Authorization"]; ok {
		token = strings.TrimSpace(token)
		if token == "" || len(token) < 8 || (token[:7] != BEARER_SCHEMA) {
			return "", errors.New("Token is invalid")
		}
		token = strings.TrimSpace(token[7:])
		return token, nil
	}
	return "", errors.New("Token not found")
}

// Header get header
func (ctx *AppContext) Header(v string) (string, error) {
	if header, ok := ctx.Meta()[v]; ok {
		return header, nil
	}
	return "", errors.New(fmt.Sprintf("%s not found", v))
}

// VerifyToken verify bearer token
func (ctx *AppContext) VerifyToken(tokenTypes ...string) (*string, *utils.CustomClaims, error) {
	namespace, _ := ctx.Header("Micro-Namespace")
	contentType, _ := ctx.Header("Content-Type")
	if _, err := ctx.Header("Method"); err != nil && namespace == "go.micro" && contentType == "application/grpc+proto" {
		id := "micro_service"
		claims := utils.CustomClaims{
			StandardClaims: jwt.StandardClaims{
				Audience: msAudienceName,
				Subject:  id,
			},
		}
		ctx.claims = &claims
		return &id, &claims, nil
	}
	oauth := new(bootstrap.OAuth)

	if token, err := ctx.Token(); err == nil {
		claims, err := oauth.VerifyJWT(token)
		if err != nil {
			return nil, nil, err
		}
		ctx.claims = claims
		match := false
		if len(tokenTypes) > 0 {
			if tokenTypes[0] != "*" {
				for _, v := range tokenTypes {
					if claims.Audience == v {
						match = true
						break
					}
				}
			} else {
				match = true
			}
		}
		if match == false {
			return nil, nil, errors.New("token type is invalid")
		}

		redisKey := fmt.Sprintf("tkid_%s", claims.Id)
		val, err := contract.redis.DB().Get(redisKey).Result()
		if err != nil && err != redis.Nil {
			return nil, nil, err
		} else if val == "revoked" {
			return nil, nil, errors.New("Token has been revoked")
		}

		return &claims.Subject, claims, nil
	}

	return nil, nil, errors.New("cannot verify token")
}

// VerifyPermission check token permission
func (ctx *AppContext) VerifyPermission(permissions ...string) (bool, error) {
	if ctx.claims == nil {
		return false, errors.New("Token not found")
	}
	if ctx.claims.Audience == msAudienceName {
		return true, nil
	}
	if len(permissions) > 0 {
		if permissions[0] == "*" {
			return true, nil
		}
		for _, v := range permissions {
			for _, p := range ctx.claims.Permissions {
				if v == p {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

// GetClaims get jwt claims
func (ctx *AppContext) GetClaims() *utils.CustomClaims {
	return ctx.claims
}

// VerifyBasicAuth verify basic auth
func (ctx *AppContext) VerifyBasicAuth() error {
	authHeader, err := ctx.Header("X-Authorization")
	if err != nil {
		return err
	}
	str, err := base64.StdEncoding.DecodeString(authHeader[len(BASIC_SCHEMA):])
	if err != nil {
		return errors.New("Base64 encoding issue")
	}
	creds := strings.Split(string(str), ":")
	if len(creds) != 2 {
		return errors.New("Username or Password is invalid")
	}
	username := creds[0]
	password := creds[1]
	if username == os.Getenv("BASIC_AUTH_USERNAME") && password == os.Getenv("BASIC_AUTH_PASSWORD") {
		return nil
	} else {
		return errors.New("Username or Password is invalid")
	}
}

// RequestKey get request key
func (ctx *AppContext) RequestKey(item interface{}) []string {
	data := ctx.StructToMap(item, "json")
	return funk.Keys(data).([]string)
}

// ModelData wrap request key contain model column name
func (ctx *AppContext) ModelData(model interface{}, req map[string]interface{}) map[string]interface{} {
	modelKeys := ctx.GetModelColumns(model)
	data := map[string]interface{}{}
	for i, v := range req {
		if funk.Contains(modelKeys, i) {
			data[i] = v
		}
	}
	return data
}

// StructToMap map struct value
func (ctx *AppContext) StructToMap(item interface{}, tag string) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get(tag)
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = ctx.StructToMap(field, tag)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}

// GetModelColumns get column name from struct model (struct field must define column option in tag)
func (ctx *AppContext) GetModelColumns(item interface{}) []string {
	res := []string{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("gorm")
		if tag != "" && tag != "-" {
			column := ""
			tagOpts := strings.Split(tag, ";")
			for _, vtag := range tagOpts {
				if strings.HasPrefix(vtag, "column:") {
					column = strings.TrimPrefix(vtag, "column:")
					break
				}
			}
			if column != "" && v.Field(i).Type.Kind() != reflect.Struct {
				res = append(res, column)
			}
		}
	}
	return res
}

// Validate validate request data and print response if have error
func (ctx *AppContext) Validate(s interface{}, prefixNamespace ...string) error {

	err := validate.Struct(s)
	if validate, ok := err.(*validator.InvalidValidationError); ok {
		panic(validate)
	}

	var errMsg string
	//Validation errors occurred
	errorsFields := make(map[string]*microError.ValidateErrors)
	errorLists := make([]string, 0)
	//Use reflector to reverse engineer struct
	if err != nil {
		for _, validate := range err.(validator.ValidationErrors) {
			//If json tag doesn't exist, use lower case of name
			namespace := strings.Split(validate.StructNamespace(), ".")
			namespace = namespace[1 : len(namespace)-1]
			namespace = append(prefixNamespace, namespace...)
			nodePrefix := ""
			if len(namespace) > 0 {
				subNamespace := []string{}
				for _, v := range namespace {
					subNamespace = append(subNamespace, wrappers.SnakeCase(v))
				}
				nodePrefix = strings.Join(subNamespace, ".") + "."
			}
			name := wrappers.SnakeCase(validate.StructField())
			fieldName := fmt.Sprintf("%s%s", nodePrefix, name)
			keyName := strings.Replace(fmt.Sprintf("%s", name), "_", " ", -1)
			value := validate.Param()
			if _, ok := errorsFields[fieldName]; !ok {
				errorsFields[fieldName] = &microError.ValidateErrors{
					FieldName: fieldName,
				}
			}
			var msg string
			switch validate.Tag() {
			case "required":
				msg = fmt.Sprintf("The %s is required", keyName)
				break
			case "required_if":
				param := strings.Split(value, `:`)
				paramField := wrappers.SnakeCase(strings.Replace(fmt.Sprintf("%s", param[0]), "_", " ", -1))
				paramValue := param[1]
				msg = fmt.Sprintf("The %v is required when %v is %v", keyName, paramField, paramValue)
				break
			case "required_if_not":
				param := strings.Split(value, `:`)
				paramField := wrappers.SnakeCase(strings.Replace(fmt.Sprintf("%s", param[0]), "_", " ", -1))
				paramValue := param[1]
				msg = fmt.Sprintf("The %v is required when %v is not equal %v", keyName, paramField, paramValue)
				break
			case "required_without":
				msg = fmt.Sprintf("The %v is required when %v is empty", keyName, value)
				break
			case "email":
				msg = fmt.Sprintf("The %s should be a valid email", keyName)
				break
			case "eq":
				msg = fmt.Sprintf("The %v should equal with %v", keyName, value)
			case "gt":
				msg = fmt.Sprintf("The %v must greater than %v", keyName, value)
			case "dive":
				msg = fmt.Sprintf("The %v must be an array %v", keyName, value)
			case "oneof":
				param := strings.Join(strings.Split(value, ` `), ", ")
				msg = fmt.Sprintf("The %v does not exist in %v", keyName, param)
			case "eqfield":
				msg = fmt.Sprintf("The %s should be equal to %s", keyName, validate.Param())
				break
			case "numeric":
				msg = fmt.Sprintf("The %s should be a valid numeric", keyName)
				break
			default:
				msg = fmt.Sprintf("The %s is invalid", keyName)
				break
			}
			errorsFields[fieldName].Message = append(errorsFields[fieldName].Message, msg)
			errorLists = append(errorLists, msg)
			errMsg += fmt.Sprintf("%s\n", msg)
		}
		return microError.NewValidateError(errorsFields)
	} else {
		return nil
	}
}
