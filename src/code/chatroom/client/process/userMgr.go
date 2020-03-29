package process

import(
	"code/chatroom/comon/message"
	"fmt"
	"code/chatroom/client/model"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User,10)
var CurUser model.CurUser
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes)   {
	user , ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
func outputOnlineUser() {
	fmt.Println("当前用户在线列表：")
	for id,_ := range onlineUsers {
		fmt.Println("用户id：",id)
	}
}