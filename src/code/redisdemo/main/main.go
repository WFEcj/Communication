package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
func main()  {
	conn , err := redis.Dial("tcp","localhost:6379")
	if err != nil {
		fmt.Println("redis.Dial err",err)
		return 
	}
	defer conn.Close()
	_ ,err = conn.Do("set","name","tomjerry")
	if err != nil {
		fmt.Println("set err" ,err)
		return 
	}
	r,err := redis.String(conn.Do("get","name"))
	if err!= nil {
		fmt.Println("err",err)
		return
	}

	fmt.Println("name",r)
}