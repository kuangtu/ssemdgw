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
	defer rconn.Close()
	return 0
}

//基于协议进行验证
func (mdgw *MdgwSession) VerifyMDGW() bool {
	defer mdgw.Rconn.Close()
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
	readnum := sock.ReadFromSock(mdgw.Rconn, mdgw.RecvBuf)

	//验证登录结果
	if readnum != -1 {
		fmt.Println("read from socket number is:", readnum)
	}

	return true

}
