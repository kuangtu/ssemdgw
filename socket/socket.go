packet mdgwsocket

import (
    "io"
    "net"
)

//socket管理结构体
type MdgwSock struct {
    recvBytes   uint64
    //本地和远程服务端地址
    laddr, raddr *net.TCPAddr
    lconn, rconn io.ReadWriteCloser
    //日志
    Log Logger
}

//创建MdgwSock管理对象
func New(raddr *net.TCPAddr) *MdgwSock {
    return &MdgwSock {
        recvBytes: 0,
        raddr: raddr,
        //日志变量
        //Log: NullLogger {}
    }
}

//发起TCP连接，连接MDGW行情网关
func (mdgw *MdgwSock) ConnMDGW() int{
    mdgw.rconn, err : = net.DialTCP("tcp", nil, mdgw.raddr)
    //连接行情网关失败，等待重新连接TODO
    if err != nil {
        mdgw.Log.Warn("Remote addr connection failed:%s", err)
        return
    }
    
    //log连接成功
    mdgw.Log.Info("conncet remote addr success:%s", mdgw.raddr)
    
    return 0
}

//基于协议进行验证
func (mdgw *MdgwSock) VerifyMDGW() bool {
    //创建验证消息
    
    //发送验证消息
    
    //接收验证消息

}

//socket监听

//socket读取


