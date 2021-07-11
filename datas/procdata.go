package datas

import (
	"fmt"
	msg "ssevss/message"
	"sync"
)

//处理登录消息
func ProcLoginMsg(loginMsg *msg.LoginMsg) {

}

//处理注销消息
func ProcLogoutMsg(logoutMsg *msg.LogoutMsg) {

}

//处理市场状态消息
func ProcMtStatusMsg(mktStatusMs *msg.MktStatusMsg) {

}

//处理心跳消息
func ProcHeartBtMsg(heartBtMsg *msg.HeartBtMsg) {

}

//处理行情快照消息
func ProcHqSnapMsg(hqSnapMsg *msg.HqSnapMsg) {

}

//处理指数行情

//处理竞价行情
func ProcMdgwMsg(ch chan int, wait *sync.WaitGroup) bool {

	//等待执行
	<-ch
	fmt.Println("start ProcMdgw msg")
	wait.Done()
	return true
}
