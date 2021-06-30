package mdgwsocket

import (
	"bytes"
	"fmt"
	"io"
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

func IsFullMessage(b *bytes.Buffer) bool {
	//检查长度
	fmt.Println("the read buffer len is:", b.Len())

	return true
}

//从socket读取字节存放到buffer中，
//在读取的时候判断接收到的字节是否到达一个数据包，然后返回处理
func ReadFromSock(rconn io.ReadWriteCloser, b *bytes.Buffer) int {
	readnum, err := rconn.Read(b.Bytes())
	if err != nil {
		fmt.Println("the read return number is:", readnum)
	}

	return readnum
}
