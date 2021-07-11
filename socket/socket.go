package mdgwsocket

import (
	"fmt"
	"io"
	"net"
)

const (
	SOCK_WRITE_ERR        = -1
	SOCKET_WRITE_PART_ERR = -2
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

//发起TCP连接
func ConnGateWay(raddr *net.TCPAddr) (*net.TCPConn, error) {
	//发起tcp连接
	rconn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		fmt.Printf("Remote addr connection failed:%s\n", err)
		return rconn, err
	}

	fmt.Printf("conncet remote addr success:%s\n", raddr)

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

// //判断是一个完整的消息，返回消息的长度
// func IsFullMessage(b *bytes.Buffer) int {
// 	//检查长度
// 	var msglen int
// 	if b.Len() < msg.MSGHEADER_LEN {
// 		msglen = 0
// 	} else if b.Len() == msg.MSGHEADER_LEN {
// 		return b.Len()
// 	} else { //收到的消息长度大于消息头部
// 		var testByte []byte
// 		copy(testByte, b.Bytes()[:msg.MSGHEADER_LEN])

// 		//获取消息头部
// 		msgHeader := &msg.MsgHeader{}
// 		msg.GetMsgHeader(msgHeader, testByte, msg.MSGHEADER_LEN)
// 		//获取消息体长度，然后检查buffer中的字节序列的长度是否大于等于消息的长度
// 		msglen := msgHeader.BodyLength + msg.MSGHEADER_LEN + msg.MSGTAIL_LEN

// 		if msglen <= uint32(b.Len()) {
// 			return b.Len()
// 		}
// 	}

// 	return msglen
// }

// //从socket读取字节存放到buffer中，
// //在读取的时候判断接收到的字节是否到达一个数据包，然后返回处理
// func ReadFromSock(rconn io.ReadWriteCloser, b *bytes.Buffer) int {
// 	//根据协议，验证消息小于1024
// 	readbuf := make([]byte, 1024)
// 	readnum, err := rconn.Read(readbuf)
// 	if err != nil {
// 		fmt.Println("ReadFromSock the read mdgw error:", err)
// 		return -1
// 	}
// 	fmt.Println("ReadFromSock number is:", readnum)

// 	//判断buffer中的字节是否构成了完成的消息
// 	b.Write(readbuf[:readnum])

// 	msgReady := IsFullMessage(b)

// 	return msgReady
// }
