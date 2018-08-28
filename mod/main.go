package main

import (
	"log"

	"github.com/go-redis/redis"
)

func main() {
	c := redis.Client{}
	log.Println(c)
}
