// auth: kunlun
// date: 2019-01-10
// description: redis连接池
package config

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/configor"
	"time"
)

type Redis struct {
	Host      string //主机ip
	Port      int    //主机端口
	Password  string //密码
	DataBase  int    //数据库
	MaxActive int    //最大活跃数
	MaxIdle   int    //最大空闲数
}

var RedisPool *redis.Pool

// 创建redis连接池
func newPool(config *Redis) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: (60 * time.Second),
		Dial: func() (conn redis.Conn, e error) {
			server := config.Host + ":" + fmt.Sprintf("%d", config.Port)
			conn, e = redis.Dial("tcp", server)
			if e != nil {
				return nil, e
			}
			if config.Password != "" {
				if _, e = conn.Do("AUTH", config.Password); e != nil {
					conn.Close()
					return nil, e
				}
			}
			return conn, e
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func InitRedisPool() *redis.Pool {
	config := Redis{}
	configor.Load(&config, "redis.json")
	if (Redis{} == config) {
		return nil
	}
	RedisPool = newPool(&config)
	return RedisPool
}

func init() {
	InitRedisPool()
}
