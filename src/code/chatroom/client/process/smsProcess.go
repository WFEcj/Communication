package process
import(
	"fmt"
	"code/chatroom/comon/message"
	"encoding/json"
	"code/chatroom/client/util"
)
type SmsProcess struct {

}
func (this *SmsProcess) SendGroupMes (content string)(err error)   {
	var mes message.Message 
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes 
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("序列化失败...")
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败...")
		return 
	}
	tf := &util.Transfer {
		Conn : CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送失败...")
		return
	}
	return
}

