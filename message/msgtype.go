package mdgwmsg

import (
	"bytes"
	"encoding/binary"
)

const (
	//消息字符串填充
	SenderCompID     = "CSI"
	TargetCompID     = "SSE"
	LOGINMSG_TYPE    = "S001"
	AppVerID         = "1.00"
	MsgType_LEN      = 4
	SenderCompID_LEN = 32
	TargetCompID_LEN = 32
	AppVerID_LEN     = 8
	//消息体头部长度
	MSGHEADER_LEN = 24
	//消息体长度
	LOGINMSG_BODY_LEN = 74
	LOGINMSG_TAIL_LEN = 4
	UINT64_LEN        = 8
	UINT32_LEN        = 4
	UINT16_LEN        = 2
)

//创建MDGW相关消息类型
type MDGWmsg struct {
}

type MsgHeader struct {
	MsgType      [4]byte
	SendingTtime uint64
	MsgSeq       uint64
	BodyLength   uint32
}

type MsgTail struct {
	CheckSum uint32
}

//登录消息
type LoginMsg struct {
	MsgHeader
	SenderCompID [32]byte
	TargetCompID [32]byte
	HeartBtInt   uint16
	AppVerID     [8]byte
	MsgTail
}

//注销消息
type LogoutMsg struct {
	Header        MsgHeader
	SessionStatus uint32
	Text          [256]byte
	MsgTail
}

//心跳消息
type HeartBtMsg struct {
	Header MsgHeader
	MsgTail
}

//市场状态消息
type MktStatusMsg struct {
	Header           MsgHeader
	SecurityType     uint8
	TradSesMode      uint8
	TradingSessionID [8]byte
	TotNoRelatedSym  uint32
	MsgTail
}

//行情快照
type SnapMsg struct {
	Header            MsgHeader
	SecurityType      uint8
	TradSesMode       uint8
	TradeDate         uint32
	LastUpdateTime    uint32
	MDStreamID        [5]byte
	SecurityID        [8]byte
	Symbol            [8]byte
	PreClosePx        uint64
	TotalVolumeTraded uint64
	NumTrades         uint64
	TotalValueTraded  uint64
	TradingPhaseCode  [8]byte
}

//指数行情快照
//根据条目个数需要进行扩展
type IndexSnap struct {
	SnapData    SnapMsg
	NoMDEntries uint16
	MDEntryType [2]byte
	MDEntryPx   uint64
}

//竞价行情快照
type BidSnap struct {
	SnapData          SnapMsg
	NoMDEntries       uint16
	MDEntryType       [2]byte
	MDEntryPx         uint64
	MDEntrySize       uint64
	MDEntryPositionNo uint8
}

func NewHeader(msgTytpe [4]byte, SendingTtime uint64, MsgSeq uint64, BodyLength uint32) (msgHeader *MsgHeader) {

}

//初始化
func initLoginMsg(loginMsg *LoginMsg) {

	//按照接口规范初始化char字符串类型，通过空格填充
	for i, _ := range loginMsg.SenderCompID {
		loginMsg.SenderCompID[i] = ' '
	}

	for i, _ := range loginMsg.TargetCompID {
		loginMsg.TargetCompID[i] = ' '
	}

	for i, _ := range loginMsg.AppVerID {
		loginMsg.AppVerID[i] = ' '
	}

}

//填充消息体
func setLoginMsgBody(loginMsg *LoginMsg) {
	//填充发送ID
	var setStr []byte
	setStr = []byte(SenderCompID)
	for i, c := range setStr {
		loginMsg.SenderCompID[i] = c
	}

	//填充目标ID
	setStr = []byte(TargetCompID)
	for i, c := range setStr {
		loginMsg.TargetCompID[i] = c
	}

	//填充APP
	setStr = []byte(AppVerID)
	for i, c := range setStr {
		loginMsg.AppVerID[i] = c
	}

	//填充心跳时间
	loginMsg.HeartBtInt = 1
}

func setLoginMsgHeader(loginMsg *LoginMsg, sendingTtime, msgSeq uint64) {
	//填充消息类型
	var setStr []byte
	setStr = []byte(LOGINMSG_TYPE)
	for i, c := range setStr {
		loginMsg.MsgType[i] = c
	}

	//填充消息序号
	loginMsg.MsgSeq = msgSeq
	msgSeq = msgSeq + 1

	//填充发送时间
	//获取当前时间
	loginMsg.SendingTtime = sendingTtime

	//消息体长度
	loginMsg.BodyLength = LOGINMSG_BODY_LEN
}

//计算登录消息校验和并填充MsgTail字段，返回字节数组buffer，后续进行发送
func calLoginMsgChkSum(loginMsg *LoginMsg) []byte {
	//将数据包中的字段放入到字节数组中，计算校验和
	buf := new(bytes.Buffer)
	//写入消息类型
	buf.Write(loginMsg.MsgType[:])
	//按照大端方式写入整数
	//写入发送时间
	binary.Write(buf, binary.BigEndian, loginMsg.SendingTtime)
	//写入消息序号
	binary.Write(buf, binary.BigEndian, loginMsg.MsgSeq)
	//写入消息体长度
	binary.Write(buf, binary.BigEndian, loginMsg.BodyLength)
	//写入发送ID
	buf.Write(loginMsg.SenderCompID[:])
	//写入目标ID
	buf.Write(loginMsg.TargetCompID[:])
	//写入心跳时间
	binary.Write(buf, binary.BigEndian, loginMsg.HeartBtInt)
	//写入版本信息
	buf.Write(loginMsg.AppVerID[:])

	//计算
	chksum := calCheckSum(buf, MSGHEADER_LEN+LOGINMSG_BODY_LEN)
	loginMsg.CheckSum = chksum
	return buf.Bytes()
}

//创建登录消息
func NewLoginMsg(sendingTtime, msgSeq uint64) *LoginMsg {
	loginMsg := &LoginMsg{}
	//初始化登录消息
	initLoginMsg(loginMsg)
	//填充消息体
	setLoginMsgBody(loginMsg)
	//填充消息头部
	setLoginMsgHeader(loginMsg, sendingTtime, msgSeq)
	//计算数据包校验和

	return loginMsg
}

//创建登录返回消息
func (mdgwmsg *MDGWmsg) NewLogoutMsg() *LogoutMsg {

}
