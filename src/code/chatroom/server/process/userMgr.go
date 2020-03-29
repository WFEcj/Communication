package process

import(
	"fmt"
)

var(
	userMgr *UserMgr
)
type UserMgr struct {
	onlinesUsers map[int]*UserProcess
}

func init()  {
	userMgr = &UserMgr{
		onlinesUsers : make(map[int]*UserProcess, 1024),
	}
}
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlinesUsers[up.UserId] = up
}
func (this *UserMgr) DelOnlineUser(userId int ) {
	delete(this.onlinesUsers,userId)
}
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlinesUsers
}

func (this *UserMgr) GetOnlineUserById (userId int ) (up *UserProcess, err error) {
	up , ok := this.onlinesUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d不存在...",userId)
		return 
	}
	return
}