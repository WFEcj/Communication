package main

import(
	"fmt"
	"net"
	"code/chatroom/comon/message"
	"code/chatroom/server/util"
	"io"
	"code/chatroom/server/process"
)

type Processor struct{
	Conn net.Conn
}

func (this *Processor) process2() (err error) {
	for{
		tf := &util.Transfer {
			Conn : this.Conn,
		}
		mes,err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端推出服务，服务器也退出..")
				return err
			} else {
				fmt.Println("readPkg err",err)
				return err
			}
		}
		fmt.Println("msg=",mes)
		err = this.serverProcessMes(&mes)
		if err !=nil {
			return err
		}
	}
}
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录函数
		up := &process.UserProcess {
			Conn : this.Conn,
		}
		err = up.ServerProcessLogin(mes)
		if err != nil {
			return 
		}
	case message.RegisterMesType:
		//处理注册函数
		up := &process.UserProcess {
			Conn : this.Conn,
		}
		err = up.ServerProcessRegister(mes)
		if err != nil {
			return 
		}
	case message.SmsMesType:
		smsProcess := &process.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在....")	
	}
	return 
}