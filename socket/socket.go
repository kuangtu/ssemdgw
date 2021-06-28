package mdgwsocket

import (
	"fmt"
	"net"
)

//socket地址解析
func NewSockAddr(addrStr string) *net.TCPAddr {
	//解析socket地址
	addr, err := net.ResolveTCPAddr("tcp", addrStr)
	if err != nil {
		fmt.Println("the addr is error:", addrStr)
	}

	return addr
}
