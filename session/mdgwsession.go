package mdgwsession

import (
	"bytes"
	"fmt"
	"io"
	"net"
	msg "ssevss/message"
	sock "ssevss/socket"
	"sync"

	mdgwutils "ssevss/utils"

	log "github.com/sirupsen/logrus"
)

const (
	MDGWVERIFY_OK          = 1
	LOGINMDGW_OK           = 0
	CONN_FAILED            = -1
	MDGWVERIFY_FAILED      = -2
	SND_VERIFYMSG_FAILED   = -3
	SND_VERIFYMSG_PART_ERR = -4

	RCV_VERIFYMSG_FAILED   = -5
	RCV_VERIFYMSG_PART_ERR = -6
)

//MdgwSession结构体
type MdgwSession struct {
	RecvBuf   *bytes.Buffer
	RecvBytes uint64
	//本地和远程服务端地址
	Laddr, Raddr *net.TCPAddr
	Lconn, Rconn io.ReadWriteCloser
	//日志
	Logger log.Logger
	//解析消息的channel
	MsgQueue chan msg.MDGWMsg
	//向MDGW发送消息时mutex
	SendMutex sync.Mutex
}

var (
	vssSession MdgwSession
	msgMutex   sync.Mutex
	msgSeq     uint64
)

func InitSession() {
	vssSession = MdgwSession{
		RecvBuf:   new(bytes.Buffer),
		RecvBytes: 0,
		MsgQueue:  make(chan msg.MDGWMsg, 100),
		//日志变量
		//Log: NullLogger {}
	}
}

//设置Session对端地址
func (mdgw *MdgwSession) SetSessionRadd(rAddr *net.TCPAddr) {
	mdgw.Raddr = rAddr
}

//设置Session远程连接
func (mdgw *MdgwSession) SetSessionRconn(rConn *net.TCPConn) {
	mdgw.Rconn = rConn
}

//连接网关进行验证
func LoginMdgw(mdgwAddr string) int {
	var iRet int
	//socket连接网关
	raddr := sock.NewSockAddr(mdgwAddr)
	//连接
	rconn, err := sock.ConnGateWay(raddr)

	//socket 连接失败
	if err != nil {
		return CONN_FAILED
	}

	//设置连接信息
	InitSession()
	vssSession.SetSessionRadd(raddr)
	vssSession.SetSessionRconn(rconn)

	//验证消息，接收验证消息异常，也认为验证失败
	iRet = verifyMDGW()
	//验证通过
	if iRet == MDGWVERIFY_OK {
		return LOGINMDGW_OK
	}
	//登录失败
	return MDGWVERIFY_FAILED

}

//根据协议发送消息进行验证
func verifyMDGW() int {

	//创建验证消息
	logingMsg, buf := msg.NewLoginMsg(mdgwutils.GetCurTime(), getMsgSeq())
	fmt.Printf("the long msg check sum is:%x", logingMsg.CheckSum)

	//发送消息进行验证
	sendnum, err := sock.WriteSock(vssSession.Rconn, buf.Bytes(), buf.Len())

	if err != nil {
		fmt.Println("socket send verify message failed.")

		return SND_VERIFYMSG_FAILED
	}
	if sendnum != buf.Len() {
		fmt.Println("send verify message part")

		return SND_VERIFYMSG_PART_ERR
	}

	//接收消息解析验证结果
	readbuf := make([]byte, msg.LOGINMSG_LEN)

	readnum, err := sock.ReadSock(vssSession.Rconn, readbuf, msg.LOGINMSG_LEN)

	if err != nil {
		fmt.Println("socket read verify message failed.")
		return RCV_VERIFYMSG_FAILED
	}

	if readnum != msg.LOGINMSG_LEN {
		fmt.Println("recv verify message part")

		return RCV_VERIFYMSG_PART_ERR
	}

	//判断接收到的消息内容
	recvMsg := msg.GetMsgFromBytes(readbuf, readnum)
	//判断得到的消息类型
	switch v := recvMsg.(type) {
	case *msg.LoginMsg:
		fmt.Println("verify mdgw get msg is loginMsg", v.MsgType)
		return MDGWVERIFY_OK
	default:
		fmt.Printf("other msg type")
		return MDGWVERIFY_FAILED
	}

}

// // //发起TCP连接，连接MDGW行情网关
// func (mdgw *MdgwSession) ConnMDGW() int {
// 	rconn, err := net.DialTCP("tcp", nil, mdgw.Raddr)
// 	//连接行情网关失败，等待重新连接TODO
// 	if err != nil {
// 		mdgw.Logger.Warn("Remote addr connection failed:%s", err)
// 		return -1
// 	}

// 	//log连接成功
// 	mdgw.Logger.Info("conncet remote addr success:%s", mdgw.Raddr)

// 	mdgw.Rconn = rconn
// 	return 0
// }

// //基于协议进行验证
// func (mdgw *MdgwSession) VerifyMDGW() bool {
// 	var verifyRes bool = false
// 	//创建验证消息
// 	logingMsg, buf := msg.NewLoginMsg(3, 2)
// 	fmt.Printf("the long msg check sum is:%x", logingMsg.CheckSum)
// 	//发送验证消息
// sendnum, err := mdgw.Rconn.Write(buf.Bytes())
// 	if err != nil {
// 		fmt.Println("socket send error:")
// 		return false
// 	}
// 	fmt.Println("sendnum is:", sendnum)
// 	//接收验证消息
// 	readbuf := make([]byte, msg.LOGINMSG_LEN)
// 	readnum, err := mdgw.Rconn.Read(readbuf)
// 	if err != nil {
// 		fmt.Println("Verify message, readfrom mdgw error:", err)
// 		return false
// 	}

// 	//通过字节流获取返回的消息
// 	fmt.Println("read number is:", readnum)
// 	verifyMsg := msg.GetMsgFromBytes(readbuf, readnum)
// 	//判断得到的消息类型
// 	switch v := verifyMsg.(type) {
// 	case *msg.LoginMsg:
// 		fmt.Println("verify mdgw get msg is loginMsg", v.MsgType)
// 		verifyRes = true
// 	case *msg.LogoutMsg:
// 		fmt.Println("verify mdgw get msg is logoutMsg", v.MsgType)
// 		verifyRes = false
// 	default:
// 		fmt.Printf("other msg type")
// 		verifyRes = false

// 	}

// 	return verifyRes

// }

// //验证通过之后，从MDGW接收流行情进行解析处理
// func RecvAndProc(mdgw *MdgwSession) bool {

// 	//从MDGW行情网关接收行情到buffer
// 	for {
// 		//大于0的长度，表示接收到的字节组成了完成的消息（一个或者多个）
// 		//可以去读消息继续
// 		//0表示没有消息
// 		//-1表示读取失败
// 		nRead := sock.ReadFromSock(mdgw.Rconn, mdgw.RecvBuf)

// 		if nRead == 0 {
// 			continue
// 		} else if nRead == -1 {
// 			//读取错误，调出循环
// 			fmt.Println("read socket error")
// 			break
// 		} else { //包含了完整的消息，进行解析
// 			msg.ParseMsg(mdgw.RecvBuf.Next(nRead), nRead)
// 		}
// 	}

// 	return false
// }

//通过验证之后，从MDGW网关接收行情，以goroutine方式运行
func RecvMdgwMsg(wait *sync.WaitGroup) bool {

	//启动标志
	fmt.Println("recv mdgw msg")

	//开始进行接收
	for {
		//大于0的长度，表示接收到的字节组成了完成的消息（一个或者多个）
		//可以去读消息继续
		//0表示没有消息
		//-1表示读取失败
		nRead := sock.ReadFromSock(vssSession.Rconn, vssSession.RecvBuf)
		if nRead == 0 { //读到的长度为0?
			continue
		} else if nRead == -1 {
			//读取错误，调出循环
			fmt.Println("read socket error")
			break
		} else { //包含了完整的消息，进行解析
			msgLen := msg.IsFullMessage(vssSession.RecvBuf)

			if msgLen > 0 {
				getMsg := msg.GetMsgFromBytes(vssSession.RecvBuf.Next(msgLen), msgLen)
				//将获取的消息存放到队列中进行解析
				vssSession.MsgQueue <- getMsg
			}
		}
		//继续从socket中读取，然后放入到RecvBuf中
	}

	//发送退出消息到解析队列
	queNoticeMsg, _ := msg.NewQueueNoticeMsg()
	vssSession.MsgQueue <- queNoticeMsg
	wait.Done()
	return true
}

//处理竞价行情
func ProcMdgwMsg(wait *sync.WaitGroup) bool {

	//等待执行
	fmt.Println("start ProcMdgw msg")
	for {
		getMsg := <-vssSession.MsgQueue
		iRet := msg.ParseMsg(getMsg)
		if iRet == 0 {
			//退出控制消息
			break
		}

	}
	wait.Done()
	return true
}

//发送心跳消息
func SendHeartBtMsg(wait *sync.WaitGroup) bool {
	defer vssSession.SendMutex.Unlock()
	//创建消息
	heartBtMsg, buf := msg.NewHeartBtMsg(mdgwutils.GetCurTime(), getMsgSeq())
	vssSession.SendMutex.Lock()
	mdgwutils.UNUSED(heartBtMsg, buf)

	//通过socket发送心跳消息
	return true
}

//获取消息编号
func getMsgSeq() uint64 {
	defer msgMutex.Unlock()

	msgMutex.Lock()
	msgSeq += 1

	return msgSeq
}
