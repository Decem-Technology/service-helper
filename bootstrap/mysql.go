package bootstrap

import (
	"fmt"
	"os"
	"strconv"
	"time"

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

type MySQLConn struct {
	mysql.Config
	ConnectionName  string // empty is default
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxIdleTime *time.Duration
	ConnMaxLifetime *time.Duration
}

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
	} else {
		maxConn := 2
		if v := os.Getenv("MYSQL_MAX_CONN"); v != "" {
			if newConn, err := strconv.Atoi(v); err == nil {
				maxConn = newConn
			}
		}
		maxLifetime, _ := time.ParseDuration("4m")
		if v := os.Getenv("MYSQL_MAX_LIFETIME"); v != "" {
			maxLifetime, _ = time.ParseDuration(os.Getenv("MYSQL_MAX_LIFETIME"))
		}
		// if v := conf.MaxIdleConns; v > 0 {
		// 	c.SetMaxIdleConns(v)
		// }
		if v := maxConn; v > 0 {
			c.SetMaxOpenConns(v)
		}
		// if v := conf.ConnMaxIdleTime; v != nil {
		// 	c.SetConnMaxIdleTime(*v)
		// }
		c.SetConnMaxLifetime(maxLifetime)

	}
	c.SetMaxIdleConns(2)
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
