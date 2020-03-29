package main

import (
	"fmt"
	"code/chatroom/client/process"
)

var(
	userId int
	userPwd string
	userName string
)

func main()  {
	var key int 
	loop := true
	for loop{
		fmt.Println("----------欢迎登录多人聊天系统----------")
		fmt.Println("\t1 登录聊天室")
		fmt.Println("\t2 注册用户")
		fmt.Println("\t3 退出系统")
		fmt.Println("\t请选择<1 - 3>:")
		fmt.Scanln(&key)
		switch  key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Printf("请输入用户的ID：")
			fmt.Scanln(&userId)
			fmt.Printf("请输入用户密码：")
			fmt.Scanln(&userPwd)
			up := &process.UserProcess{}
			err := up.Login(userId,userPwd)
			if err != nil {
				fmt.Println("login 函数出错");
				return 
			}
			//loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Printf("请输入用户的ID：")
			fmt.Scanln(&userId)
			fmt.Printf("请输入用户密码：")
			fmt.Scanln(&userPwd)
			fmt.Printf("请输入用户名称：")
			fmt.Scanln(&userName)
			up := &process.UserProcess{}
			err := up.Register(userId,userPwd,userName)
			if err != nil {
				fmt.Println("login 函数出错");
				return 
			}
			//loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("您的输入有误 请重新输入！")
		}
	}
}