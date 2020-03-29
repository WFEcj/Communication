package process

import(
	"fmt"
	"code/chatroom/comon/message"
	"code/chatroom/server/util"
	"net"
	"encoding/json"
)
type SmsProcess struct {

}
func (this *SmsProcess) SendGroupMes (mes *message.Message)  {
	var smsMes message.SmsMes 
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("Unmarshal err",err)
		return
	}
	data ,err := json.Marshal(mes)
	if err !=nil {
		fmt.Println("Marshal err",err)
		return 
	}
	for id,up := range userMgr.onlinesUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data,up.Conn)
	}
}
func (this *SmsProcess) SendMesToEachOnlineUser (data []byte, conn net.Conn) {
	tf := &util.Transfer{
		Conn : conn, 
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("send msg err",err)
		return
	}
}