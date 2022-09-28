// auth: kunlun
// date: 2019-01-10
// description: redis操作工具类
package common

import (
	"config"
	"github.com/gomodule/redigo/redis"
	"log"
)

//
func RedisSet(key, value string) bool {
	conn := config.RedisPool.Get()
	//释放资源
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		log.Panic("set error: ", err)

		return false
	}
	return true
}

func RedisSetWithExpire(key, value string, num int) bool {
	conn := config.RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		log.Panic("set with expire error: ", err)
		return false
	}
	_, err1 := conn.Do("EXPIRE", key, num*60)
	if err1 != nil {
		log.Panic("set with expire error: ", err1)
		return false
	}
	return true
}

func RedisGetString(key string) (value string, err error) {
	conn := config.RedisPool.Get()
	defer conn.Close()
	exist, _ := redis.Bool(conn.Do("EXISTS", key))
	if !exist {
		return "", nil
	}
	v, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Panicf("get string error: %v", err)
		return "", err
	}
	return v, nil
}

func RedisGetStringMap(key string) (value map[string]string, err error) {
	conn := config.RedisPool.Get()
	defer conn.Close()
	v, err := redis.StringMap(conn.Do("GET", key))
	if err != nil {
		log.Panic("get string map  error: ", err)
		return nil, err
	}
	return v, nil
}
