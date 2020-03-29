package process

import (
	"fmt"
	"code/chatroom/comon/message"
	"net"
	"encoding/json"
	"code/chatroom/client/util"
)

type UserProcess struct {

}

func (this *UserProcess) Login(userId int, userPwd string) (err error)  {
	conn ,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("连接失败....")
		return
	}
	defer conn.Close()
	
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	data ,err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("序列化失败....")
		return 
	}
	mes.Data = string(data)
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败.....")
		return 
	}
	tf := &util.Transfer{
		Conn : conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.Write err",err)
		return
	}
	mes ,err = tf.ReadPkg()
	if err != nil {
		fmt.Println("读取失败..")
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte (mes.Data),&loginResMes)
	if err != nil {
		fmt.Println("反序列化失败...")
		return 
	}
	if loginResMes.Code == 200 {
		//fmt.Println("登录成功")
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		for _,v := range loginResMes.UserId {
			if v == userId {
				continue
			}
			fmt.Println("用户Id：\t",v)
			user := &message.User{
				UserId : v,
				UserStatus : message.UserOnline,
			}
			onlineUsers[v] = user
		}
		go ServerProcessMes(conn)
		for {
			ShowMenu()
		}
		
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
func (this *UserProcess) Register(userId int,userPwd string ,userName string)(err error) {
	conn ,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("连接失败....")
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes 
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	data ,err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("序列化失败....")
		return 
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败...")
		return
	}
	tf := &util.Transfer{
		Conn : conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.Write err",err)
		return
	}
	mes ,err = tf.ReadPkg()
	if err != nil {
		fmt.Println("读取失败..")
		return
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte (mes.Data),&registerResMes)
	if err != nil {
		fmt.Println("反序列化失败...")
		return 
	}
	if registerResMes.Code == 200 {
		//fmt.Println("登录成功")
		fmt.Println("注册成功,请重新登录...")
		return
		
	} else {
		fmt.Println(registerResMes.Error)
		return
	}
}