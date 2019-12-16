package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type RedisClient struct {
}

func NewRedisP() *RedisClient {
	return new(RedisClient)
}

var redisPool *redis.Pool

func init() {
	redisClient := NewRedisP()
	redisClient.SetRedisClient("127.0.0.1:6379", "", 20, 10, 50)
}

func (s *RedisClient) GetRedisClient() *redis.Pool {
	return redisPool
}

func (s *RedisClient) SetRedisClient(redisAddress, password string, maxActive, maxIdele, timeExpiretime int) {
	timeout := time.Duration(timeExpiretime)
	redisPool = &redis.Pool{
		MaxIdle:     maxIdele,
		MaxActive:   maxActive,
		IdleTimeout: timeout,
		Wait:        true,
		Dial: func() (conn redis.Conn, e error) {
			con, err := redis.Dial("tcp", redisAddress,
				redis.DialPassword(password),
				redis.DialConnectTimeout(timeout*time.Second),
				redis.DialReadTimeout(timeout*time.Second),
				redis.DialWriteTimeout(timeout*time.Second),
			)
			if err != nil {
				log.Printf("RedisClient.SetRedisClient err:%s", err.Error())
				return nil, err
			}
			return con, nil
		},
	}
}

func (s *RedisClient) SetRedisCacheMapByKey(keyName string, value interface{}, expiredSecond int32) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	_, err := redisConn.Do("set", keyName, value, "EX", expiredSecond)
	if err != nil {
		log.Printf("RedisClient.SetRedisCacheMapByKey redisConn.Do is err %s", err.Error())
	}
	return err
}

func (s *RedisClient) GetRedisStrByKey(keyName string) (string, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	return redis.String(redisConn.Do("get", keyName))
}

func (s *RedisClient) MGetRedisByHashByKeys(keyName string, args ...string) (map[string]string, []string, error) {
	var hasExistFields = make(map[string]string)
	var noExistFields []string
	redisConn := redisPool.Get()
	defer redisConn.Close()
	var argsInterface []interface{}
	argsInterface = append(argsInterface, keyName)
	for _, v := range args {
		argsInterface = append(argsInterface, v)
	}
	value, err := redis.Values(redisConn.Do("hmget", argsInterface...))
	if err != nil {
		return hasExistFields, noExistFields, err
	}

	for k, v := range value {
		if k <= len(args)-1 {
			key := args[k]
			if v != nil {
				tmpV := string(v.([]byte))
				hasExistFields[key] = tmpV
			} else {
				noExistFields = append(noExistFields, key)
			}
		}
	}
	return hasExistFields, noExistFields, nil
}

func (s *RedisClient) MGetRedisByKeys(args ...string) (map[string]string, map[string]string, error) {
	var hasExistFields = make(map[string]string)
	var noExistFields = make(map[string]string)
	redisConn := redisPool.Get()
	defer redisConn.Close()
	var argsInterface []interface{}
	for _, v := range args {
		argsInterface = append(argsInterface, v)
	}
	value, err := redis.Values(redisConn.Do("mget", argsInterface...))
	if err != nil {
		return hasExistFields, noExistFields, err
	}

	for k, v := range value {
		if k <= len(args)-1 {
			key := args[k]
			if v != nil {
				tmpV := string(v.([]byte))
				hasExistFields[key] = tmpV
			} else {
				noExistFields[key] = "" //用于上次使用
			}
		}
	}
	return hasExistFields, noExistFields, nil
}

func (s *RedisClient) MSetRedisKeyByHashKey(keyName string, args ...string) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	var argsInterface []interface{}
	argsInterface = append(argsInterface, keyName)
	for _, v := range args {
		argsInterface = append(argsInterface, v)
	}
	_, err := redisConn.Do("hmset", argsInterface...)
	return err
}

func (s *RedisClient) ScanByKey(keyName string) ([]string, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	existKeys := []string{}
	value, err := redis.Values(redisConn.Do("scan", "0", "MATCH", fmt.Sprintf("%s", keyName), "count", "100"))
	if err != nil {
		return existKeys, err
	}

	if len(value) <= 1 {
		return existKeys, nil
	}
	return redis.Strings(value[1], err)
}
