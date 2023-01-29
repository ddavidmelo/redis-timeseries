package redis

import (
	"redis-timeseries/internal/config"
	"time"

	"github.com/gomodule/redigo/redis"

	redis_timeseries_go "github.com/RedisTimeSeries/redistimeseries-go"
	log "github.com/sirupsen/logrus"
)

var database *DBHandler

type DBHandler struct {
	RedisTS *redis_timeseries_go.Client
}

// ConnectDB opens a connection to the database
func ConnectDB() {
	log.Info("connecting to Redis")

	host := config.GetRedisConfig().Host + ":" + config.GetRedisConfig().Port
	password := config.GetRedisConfig().Password
	pool := &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", host, redis.DialPassword(password))
	}}

	client := redis_timeseries_go.NewClientFromPool(pool, "ts-client")

	for {
		if err := client.Pool.Get().Err(); err != nil {
			log.Info("redis: database error, will retry in 5s")
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	database = &DBHandler{client}
	log.Info("connected to Redis")

}

func DB() *DBHandler {
	return database
}
