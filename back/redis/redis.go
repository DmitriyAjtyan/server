package redis

import (
	"fmt"
	config "server/back/config"
	"strconv"
	"time"

	redis "github.com/gomodule/redigo/redis"
)

// Options is data type for redis configuration data
type Options struct {
	Host      string
	Port      int
	Password  string
	DBName    int
	MaxIdle   int
	MaxActive int
	Timeout   time.Duration
}

// Connect is a function for connection with redis
func Connect() *redis.Pool {
	Opt := &Options{
		Host:      config.ConfigData.Redis.Host,
		Port:      config.ConfigData.Redis.Port,
		Password:  config.ConfigData.Redis.Password,
		DBName:    config.ConfigData.Redis.DBName,
		MaxIdle:   config.ConfigData.Redis.MaxIdle,
		MaxActive: config.ConfigData.Redis.MaxActive,
		Timeout:   config.ConfigData.Redis.IdleTimeout,
	}
	return &redis.Pool{
		MaxIdle:   Opt.MaxIdle,
		MaxActive: Opt.MaxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf(Opt.Host+":"+strconv.Itoa(Opt.Port)), redis.DialDatabase(Opt.DBName))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// Ping function check redis connection
func Ping(c redis.Conn) error {
	s, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)

	return nil
}
