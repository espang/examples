// +build ignore
package main

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

func newPool(addr string, database int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("Select", database); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

var pool *redis.Pool

func KeyExists(key string) bool {
	conn := pool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", "key"))
	if err != nil {
		panic(err)
	}

	return exists
}

func SetEx(key, value string, ex int) {
	conn := pool.Get()
	defer conn.Close()
	conn.Do("SETEX", key, ex, value)
}

func SetBoolean(key string, value bool) {
	conn := pool.Get()
	defer conn.Close()

	conn.Do("SET", key, value)
}

func GetBoolean(key string) bool {
	conn := pool.Get()
	defer conn.Close()

	res, _ := redis.Bool(conn.Do("GET", key))
	return res
}

func main() {
	addr := "127.0.0.1:6379"
	database := 15
	pool = newPool(addr, database)

	key := "key"
	value := "value"

	log.Print("EXISTS key")
	exists := KeyExists(key)
	log.Printf("Result: %v", exists)

	log.Print("SETEX for 4 seconds")
	SetEx(key, value, 4)
	exists = KeyExists(key)
	log.Printf("Result: %v", exists)

	log.Print("Wait for 4 seconds")
	time.Sleep(5 * time.Second)
	exists = KeyExists(key)
	log.Printf("Result: %v", exists)

	log.Print("Set True/False")
	SetBoolean("false", false)
	SetBoolean("true", true)
	log.Printf("True: %v", GetBoolean("true"))
	log.Printf("False: %v", GetBoolean("false"))

	conn := pool.Get()
	reply, err := redis.Bool(conn.Do("SETNX", "false2", true))
	conn.Close()
	log.Printf("SETNX false: %v | err: %v", reply, err)
	log.Printf("False: %v", GetBoolean("false"))
}
