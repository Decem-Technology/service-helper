package bootstrap

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	// PostgresDB PostgreSQL database management
	PostgresDB struct {
	}
)

// dbPostgresDB variable for define connection
var dbPostgresDB *gorm.DB

// CreatePostgreSQLConnection make connection
func CreatePostgreSQLConnection() {
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "jhi_" + defaultTableName
	// }
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DBNAME"),
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connection,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Sprintf("[postgres] failed to connect database: %s", err))
	}
	fmt.Println("[postgres] connected")

	c, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("[postgres] connection poll error: %s", err))
	}
	c.SetMaxIdleConns(10)
	if debug, err := strconv.ParseBool(os.Getenv("APP_DEBUG")); err == nil {
		if debug {
			db = db.Debug()
		}
	}
	dbPostgresDB = db
}

// DB get postgresql connection
func (c *PostgresDB) DB() *gorm.DB {
	return dbPostgresDB
}
