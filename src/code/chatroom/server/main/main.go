package main

import(
	"fmt"
	"net"
	"code/chatroom/server/model"
	"time"
)
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func main()  {
	initPool("localhost:6379",16,0,300 * time.Second)
	initUserDao()
	fmt.Println("服务器开始监听8889端口.....")
	listen , err := net.Listen("tcp","localhost:8889")
	if err != nil {
		fmt.Println("Listen err=",err)
	}
	defer listen.Close()
	for{
		conn , err := listen.Accept()
		if err != nil {
			fmt.Println("连接失败....")
		}
		go process1(conn)
	}
}
func process1(conn net.Conn) {
	defer conn.Close()
	processor := &Processor {
		Conn : conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端与服务器远程协助错误...")
		return
	}
}



