package message

const(
	LoginMesType			="LoginMes"
	LoginResMesType			="LoginResMes"
	RegisterMesType			="RegisterMes"
	RegisterResMesType		="RegisterResMes"
	NotifyUserStatusMesType	="NotifyUserStatusMes"
	SmsMesType				="SmsMes"
)

const(
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId int `json:"userId"` //用户Id
	UserPwd string `json:"userPwd"` //用户密码
	UserName string `json:"userName"` //用户名
}

type LoginResMes struct {
	Code int `json:"code"`
	UserId []int 
	Error string `json:"error"`
}
type RegisterMes struct {
	User User `json:"user"`
}
type RegisterResMes struct {
	Code int `json:"code"`
	Error string `json:"error"`
}
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}
type SmsMes struct {
	Content string `json:"content"`
	User
}