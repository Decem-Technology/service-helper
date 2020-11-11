package bootstrap

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	// MySQL mysql database management
	MySQL struct {
	}
)

// dbMySQL variable for define connection
var dbMySQL *gorm.DB

// CreateMySQLConnection make connection
func CreateMySQLConnection() {
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "jhi_" + defaultTableName
	// }
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"),
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: connection,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Sprintf("[MySQL] connect database fail, error: %s", err))
	}
	fmt.Println("[MySQL] connected")

	c, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("[MySQL] connection poll error: %s", err))
	}
	c.SetMaxIdleConns(10)
	if debug, err := strconv.ParseBool(os.Getenv("APP_DEBUG")); err == nil {
		if debug {
			db = db.Debug()
		}
	}
	dbMySQL = db
}

// DB get mysql connection
func (c *MySQL) DB() *gorm.DB {
	return dbMySQL
}
