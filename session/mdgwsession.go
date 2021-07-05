package mdgwsession

import (
	"bytes"
	"fmt"
	"io"
	"net"

	msg "ssevss/message"

	sock "ssevss/socket"

	log "github.com/sirupsen/logrus"
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
}

//创建MdgwSock管理对象
func NewMdgwSession(raddr *net.TCPAddr) *MdgwSession {
	return &MdgwSession{
		RecvBuf:   new(bytes.Buffer),
		RecvBytes: 0,
		Raddr:     raddr,
		//日志变量
		//Log: NullLogger {}
	}
}

//发起TCP连接，连接MDGW行情网关
func (mdgw *MdgwSession) ConnMDGW() int {
	rconn, err := net.DialTCP("tcp", nil, mdgw.Raddr)
	//连接行情网关失败，等待重新连接TODO
	if err != nil {
		mdgw.Logger.Warn("Remote addr connection failed:%s", err)
		return -1
	}

	//log连接成功
	mdgw.Logger.Info("conncet remote addr success:%s", mdgw.Raddr)

	mdgw.Rconn = rconn
	return 0
}

//基于协议进行验证
func (mdgw *MdgwSession) VerifyMDGW() bool {
	var verifyRes bool = false
	//创建验证消息
	logingMsg, buf := msg.NewLoginMsg(3, 2)
	fmt.Printf("the long msg check sum is:%x", logingMsg.CheckSum)
	//发送验证消息
	sendnum, err := mdgw.Rconn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("socket send error:")
		return false
	}
	fmt.Println("sendnum is:", sendnum)
	//接收验证消息
	readbuf := make([]byte, msg.LOGINMSG_LEN)
	readnum, err := mdgw.Rconn.Read(readbuf)
	if err != nil {
		fmt.Println("Verify message, readfrom mdgw error:", err)
		return false
	}

	//通过字节流获取返回的消息
	fmt.Println("read number is:", readnum)
	verifyMsg := msg.GetMsgFromBytes(readbuf, readnum)
	//判断得到的消息类型
	switch v := verifyMsg.(type) {
	case *msg.LoginMsg:
		fmt.Println("verify mdgw get msg is loginMsg", v.MsgType)
		verifyRes = true
	case *msg.LogoutMsg:
		fmt.Println("verify mdgw get msg is logoutMsg", v.MsgType)
		verifyRes = false
	default:
		fmt.Printf("other msg type")
		verifyRes = false

	}

	return verifyRes

}

//验证通过之后，从MDGW接收流行情进行解析处理
func RecvAndProc(mdgw *MdgwSession) bool {

	//从MDGW行情网关接收行情到buffer
	for {
		//大于0的长度，表示接收到的字节组成了完成的消息（一个或者多个）
		//可以去读消息继续
		//0表示没有消息
		//-1表示读取失败
		nRead := sock.ReadFromSock(mdgw.Rconn, mdgw.RecvBuf)

		if nRead == 0 {
			continue
		} else if nRead == -1 {
			//读取错误，调出循环
			fmt.Println("read socket error")
			break
		} else { //包含了完整的消息，进行解析
			msg.ParseMsg(mdgw.RecvBuf.Next(nRead), nRead)
		}
	}

	return false
}
