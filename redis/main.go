package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

var redisPool *redis.Pool

func main() {
	// Read .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	redisHost := os.Getenv("REDISHOST")
	redisPort := os.Getenv("REDISPORT")
	maxConnection, _ := strconv.Atoi(os.Getenv("MAX_CONNECTIONS"))
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	redisPool = redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", redisAddr)
	}, maxConnection)

	// get connection
	conn := redisPool.Get()
	defer conn.Close()

	// get value
	res := redisGet("test", conn)
	if res == "0" {
		// set initialvalue
		redisSet("test", "1", conn)
	} else {
		fmt.Println(res)
	}
	// set value
	redisSet("test", res, conn)
}

func redisSet(key string, value string, c redis.Conn) {
	tmp, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("set error")
	}
	tmp++
	w, err := c.Do("SET", key, strconv.Itoa(tmp))
	if err != nil {
		panic(err)
	}
	fmt.Println(w)
}

func redisGet(key string, c redis.Conn) string {
	r, err := redis.String(c.Do("GET", key))
	if err != nil {
		return "0"
	}
	return r
}
