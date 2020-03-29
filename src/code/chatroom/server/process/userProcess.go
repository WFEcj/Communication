package process

import(
	"fmt"
	"code/chatroom/comon/message"
	"net"
	"encoding/json"
	"code/chatroom/server/util"
	"code/chatroom/server/model"
)

type UserProcess struct {
	Conn net.Conn
	UserId int
}

func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id , up := range userMgr.onlinesUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int ) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline
	data,err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("序列化失败...")
		return
	}
	mes.Data = string(data)
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败...")
		return
	}
	tf := &util.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送失败...")
		return
	}
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte (mes.Data) , &loginMes)
	if err != nil {
		fmt.Println("反序列化失败...")
		return
	}
	//返回的实例
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	//fmt.Println("user",user)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS{
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if (err == model.ERROR_USER_PWD){
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 503
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUser(loginMes.UserId)
		for id, _ := range userMgr.onlinesUsers {
			loginResMes.UserId = append(loginResMes.UserId,id)
		} 
		fmt.Println(user,"登录成功")
	}
	
	data ,err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化失败....")
		return 
	}
	resMes.Data = string(data)
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败.....")
		return 
	}
	//writePkg
	tf := &util.Transfer {
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte (mes.Data) , &registerMes)
	if err != nil {
		fmt.Println("反序列化失败...")
		return
	}
	//返回的实例
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil{
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
	} else {
		registerResMes.Code = 200
		
	}
	data ,err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("序列化失败....")
		return 
	}
	resMes.Data = string(data)
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败.....")
		return 
	}
	//writePkg
	tf := &util.Transfer {
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
