packet mdgwsocket

import (
    "io"
    "net"
)

//socket管理结构体
type MdgwSock struct {
    laddr, raddr *net.TCPAddr
    lconn, rconn io.ReadWriteCloser
}


//发送socket连接

//socket监听

//socket读取



