package process
import(
	"fmt"
	"os"
	"net"
	"code/chatroom/client/util"
	"code/chatroom/comon/message"
	"encoding/json"
)

func ShowMenu() {
	smsMes := &SmsProcess{}
	for{
		fmt.Println("----------恭喜xxx登录成功")
		fmt.Println("\t1. 显示在线列表")
		fmt.Println("\t2. 发送消息")
		fmt.Println("\t3. 信息列表")
		fmt.Println("\t4. 退出系统")
		fmt.Println("请输入<1 - 4>:")
		var key int
		var msg string
		fmt.Scanln(&key)
		switch key {
		case 1:
			//fmt.Println("查看在线列表。。。。")
			outputOnlineUser()
		case 2:
			//fmt.Println("发送消息")
			fmt.Println("你想和大伙说点什么：")
			fmt.Scanln(&msg)
			smsMes.SendGroupMes(msg)
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("请输入正确指令")
		}
	}
}
func ServerProcessMes(conn net.Conn){
	tf := &util.Transfer{
		Conn : conn,
	}
	for{
		fmt.Println("正在接收服务器的消息。。。")
		mes,err := tf.ReadPkg()
		if err != nil {
			fmt.Println("ServerProcessMes err",err)
			return
		}
		//fmt.Println("服务器说:",mes)
		fmt.Println(mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType :
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器返回了未知类型的消息...")
		}
	}

}