package main

import(
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string, maxIdle int,maxActice int,idleTimeout time.Duration)(err error)  {
	pool = &redis.Pool{
		MaxIdle : maxIdle,
		MaxActive : maxActice,
		IdleTimeout : idleTimeout,
		Dial : func()(redis.Conn, error){
			return redis.Dial("tcp",address)
		},
	}
	return
}