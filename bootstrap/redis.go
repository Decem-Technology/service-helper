package bootstrap

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
)

type (
	// RedisDB redis database management
	RedisDB struct {
	}
)

// dbRedis variable for define connection
var dbRedis redis.UniversalClient

// CreateRedisConnection make redis connection
func CreateRedisConnection() {
	database, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	db := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:       strings.Split(os.Getenv("REDIS_HOST"), ","),
		Password:    os.Getenv("REDIS_PASSWORD"),
		DB:          database,
		DialTimeout: time.Duration(15) * time.Second,
	})
	if err := db.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}

	fmt.Println("[Redis] connected")
	dbRedis = db
}

// DB get redis connection
func (c *RedisDB) DB() redis.UniversalClient {
	return dbRedis
}
