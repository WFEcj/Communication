package process

import(
	"fmt"
	"code/chatroom/comon/message"
	"encoding/json"
)
func outputGroupMes(mes *message.Message)  {
	var smsMes message.SmsMes 
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("Unmarshal err",err)
		return
	}
	info := fmt.Sprintf("用户%d:\t对大家说:%s",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}