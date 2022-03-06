package main

import "fmt"

const (

	//消息类型
	LOGINMSG_TYPE     = "S001"
	LOGOUTMSG_TYPE    = "S002"
	QUEUE_NOTICE_TYPE = "E001"
	HEARTBTMSG_TYPE   = "S003"
	MKTSTUSMSG_TYPE   = "M101"
	MKTSNAPMSG_TYPE   = "M102"
	//消息字符串填充
	SenderCompID     = "CSI"
	TargetCompID     = "SSE"
	AppVerID         = "1.00"
	MsgType_LEN      = 4
	SenderCompID_LEN = 32
	TargetCompID_LEN = 32
	AppVerID_LEN     = 8
	LOGOUTTXT_LEN    = 256
	//消息体头部长度
	MSGHEADER_LEN = 24
	//消息体长度
	LOGINMSG_BODY_LEN = 74
	MSGTAIL_LEN       = 4
	LOGINMSG_LEN      = MSGHEADER_LEN + LOGINMSG_BODY_LEN + MSGTAIL_LEN
	UINT64_LEN        = 8
	UINT32_LEN        = 4
	UINT16_LEN        = 2

	//市场状态消息
	TradingSessionID_LEN = 8
	//行情快照消息
	MDStreamID_LEN   = 5
	SecurityID_LEN   = 8
	Symbol_LEN       = 8
	TradingPhase_LEN = 8
	MDEntryType_LEN  = 2
)

type MsgHeader struct {
	MsgType      [MsgType_LEN]byte
	SendingTtime uint64
	MsgSeq       uint64
	BodyLength   uint32
}

//登录消息
type LoginMsg struct {
	MsgHeader
	SenderCompID [SenderCompID_LEN]byte
	TargetCompID [TargetCompID_LEN]byte
	HeartBtInt   uint16
	AppVerID     [AppVerID_LEN]byte
	MsgTail
}

type MsgTail struct {
	CheckSum uint32
}

func main() {
	var msg LoginMsg
	var setStr []byte
	setStr = []byte(SenderCompID)
	for i, c := range setStr {
		fmt.Println(i, c)
	}

	fmt.Printf("the msg TargetCom id is:%d", msg.HeartBtInt)
}
