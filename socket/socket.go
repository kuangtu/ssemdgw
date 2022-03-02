package mdgwsocket

import (
	"bytes"
	"fmt"
	"io"
	"net"

	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
)

const (
	SOCK_WRITE_OK         = 0
	SOCK_WRITE_ERR        = -1
	SOCKET_WRITE_PART_ERR = -2
)

//socket地址解析
func NewSockAddr(addrStr string) *net.TCPAddr {
	//解析socket地址
	addr, err := net.ResolveTCPAddr("tcp", addrStr)
	if err != nil {
		// fmt.Println("the addr is error:", addrStr)
		logrus.Fatal("resolve tcp addr error.")
	}

	return addr
}

//发起TCP连接
func ConnGateWay(raddr *net.TCPAddr) (*net.TCPConn, error) {
	//发起tcp连接
	rconn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		// fmt.Printf("Remote addr connection failed:%s\n", err)
		log.Error("connect mdgw gateway failed:%s\n", err)
		return rconn, err
	}

	fmt.Printf("conncet remote addr success:%s\n", raddr)
	log.Info("connect mdgw gateway success:%s\n", raddr)

	return rconn, nil
}

//socket写入消息
func WriteSock(rconn io.ReadWriteCloser, buf []byte, len int) (int, error) {

	//写了一部分，然后socket断开了?
	sendnum, err := rconn.Write(buf)

	if err != nil {
		fmt.Println("write socket eror:", err)
		return sendnum, err
	}

	if sendnum != len {
		fmt.Println("write socket failed:", err)

		return sendnum, err
	}

	return sendnum, nil

}

//socket发送消息
func ReadSock(rconn io.ReadWriteCloser, buf []byte, len int) (int, error) {

	//读取特定长度的
	readnum, err := rconn.Read(buf)

	if err != nil {
		fmt.Println("read socket error:", err)
		return readnum, err
	}

	if readnum != len {
		fmt.Println("read socket failed:", err)

		return readnum, err
	}

	return readnum, nil

}

//从socket读取字节存放到buffer中，
//在读取的时候判断接收到的字节是否到达一个数据包，然后返回处理
func ReadFromSock(rconn io.ReadWriteCloser, b *bytes.Buffer) int {
	//根据协议，验证消息小于1024
	readbuf := make([]byte, 1024)
	readnum, err := rconn.Read(readbuf)
	if err != nil {
		fmt.Println("ReadFromSock the read mdgw error:", err)
		return -1
	}
	fmt.Println("ReadFromSock number is:", readnum)

	//判断buffer中的字节是否构成了完成的消息
	b.Write(readbuf[:readnum])

	return readnum

}
