package util

import(
	"fmt"
	"encoding/binary"
	"encoding/json"
	"code/chatroom/comon/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [8094]byte
}
func (this *Transfer) ReadPkg() (mes message.Message , err error) {
	_,err = this.Conn.Read(this.Buf[0:4])
	if err != nil {
		fmt.Println("conn Read err",err)
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	n ,err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return 
	}
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json Unmarshal err:",err)
		return
	}
	return 
}
func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen)
	n,err := this.Conn.Write(this.Buf[:4])
	if n!= 4 || err != nil {
		fmt.Println("conn write err,",err)
		return 
	}
	// fmt.Printf("长度为%d,内容为：%s",len(data),string(data))
	n,err = this.Conn.Write(data)
	if n != int(pkgLen) && err != nil {
		fmt.Println("conn write err",err)
		return 
	}
	return
}