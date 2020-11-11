module github.com/Decem-Technology/service-helper

go 1.14

require (
	cloud.google.com/go/firestore v1.1.1 // indirect
	firebase.google.com/go v3.13.0+incompatible
	github.com/aws/aws-sdk-go v1.34.13
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/elastic/go-elasticsearch/v7 v7.9.0
	github.com/go-playground/validator/v10 v10.3.0
	github.com/go-redis/redis/v7 v7.4.0
	github.com/golang/protobuf v1.4.1
	github.com/google/logger v1.1.0
	github.com/labstack/gommon v0.3.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/olivere/elastic/v7 v7.0.20
	github.com/robfig/cron/v3 v3.0.1
	github.com/thoas/go-funk v0.7.0
	go.mongodb.org/mongo-driver v1.4.1
	google.golang.org/api v0.14.0
	google.golang.org/protobuf v1.22.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.6
)

replace github.com/micro/go-micro/v2 v2.9.1 => github.com/Decem-Technology/nitro/v2 v2.9.1-ct.1
