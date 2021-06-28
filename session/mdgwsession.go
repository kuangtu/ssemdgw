package mdgwsession

import (
	"io"
	"net"
)

//MdgwSession结构体
type MdgwSession struct {
	recvBytes uint64
	//本地和远程服务端地址
	laddr, raddr *net.TCPAddr
	lconn, rconn io.ReadWriteCloser
	//日志
	Logger log
}

//创建MdgwSock管理对象
func NewMdgwSession(raddr *net.TCPAddr) *MdgwSession {
	return &MdgwSession{
		recvBytes: 0,
		raddr:     raddr,
		//日志变量
		//Log: NullLogger {}
	}
}

//发起TCP连接，连接MDGW行情网关
func (mdgw *MdgwSession) ConnMDGW() int {
	rconn, err := net.DialTCP("tcp", nil, mdgw.raddr)
	//连接行情网关失败，等待重新连接TODO
	if err != nil {
		mdgw.Logger.Warn("Remote addr connection failed:%s", err)
		return -1
	}

	//log连接成功
	mdgw.Logger.Info("conncet remote addr success:%s", mdgw.raddr)

	mdgw.rconn = rconn
	return 0
}

//基于协议进行验证
func (mdgw *MdgwSession) VerifyMDGW() bool {
	//创建验证消息

	//发送验证消息

	//接收验证消息

	return false

}

//socket监听

//socket读取
