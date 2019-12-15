package main

import (
	"goLirary/nosql/redis"
	"log"
)

func main() {
	redisClient := redis.NewRedisP()
	err := redisClient.SetRedisCacheMapByKey("TEST2", "2", 6000)
	log.Println(err)
	log.Println(redisClient.GetRedisStrByKey("TEST2"))
}
