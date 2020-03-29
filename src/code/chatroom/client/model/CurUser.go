package model

import(
	"net"
	"code/chatroom/comon/message"

)

type CurUser struct {
	Conn net.Conn
	message.User
}