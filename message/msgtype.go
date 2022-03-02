package mdgwmsg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"ssevss/datas"
	mdgwutils "ssevss/utils"
)

const (
	LOGINMSG_TYPE_INT = iota
	LOGOUTMSG_TYPE_INT
	HBMSG_TYPE_INT
	MKTSTATUS_TYPE_INT
	MKTHQ_TYPE_INT
	QUEUE_NOTICE_TYPE_INT
	OTHERMSG_TYPE_INT //可能收到了其他类型的消息
)

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

//创建MDGW消息结构体
type MDGWMsg interface {
	//获取消息类型
	GetMsgType() [MsgType_LEN]byte
}

type MsgHeader struct {
	MsgType      [MsgType_LEN]byte
	SendingTtime uint64
	MsgSeq       uint64
	BodyLength   uint32
}

type MsgTail struct {
	CheckSum uint32
}

//解析goroutine退出消息
//和心跳消息类似，消息类型不同
type QueueNoticeMsg struct {
	MsgHeader
	MsgTail
}

func (queueNoticeMsg *QueueNoticeMsg) GetMsgType() [msg.MsgType_LEN]byte {
	return queueNoticeMsg.MsgType
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

func (loginMsg *LoginMsg) GetMsgType() [MsgType_LEN]byte {
	return loginMsg.MsgType
}

//注销消息
type LogoutMsg struct {
	MsgHeader
	SessionStatus uint32
	Text          [LOGOUTTXT_LEN]byte
	MsgTail
}

//实现MDGWMsg接口
func (logoutMsg *LogoutMsg) GetMsgType() [MsgType_LEN]byte {
	return logoutMsg.MsgType
}

//心跳消息
type HeartBtMsg struct {
	MsgHeader
	MsgTail
}

//实现MDGWMsg接口
func (heartBtMsg *HeartBtMsg) GetMsgType() [MsgType_LEN]byte {
	return heartBtMsg.MsgType
}

//市场状态消息
type MktStatusMsg struct {
	MsgHeader
	SecurityType     uint8
	TradSesMode      uint8
	TradingSessionID [TradingSessionID_LEN]byte
	TotNoRelatedSym  uint32
	MsgTail
}

//实现MDGWMsg接口
func (mktStatusMsg *MktStatusMsg) GetMsgType() [MsgType_LEN]byte {
	return mktStatusMsg.MsgType
}

//行情快照，包含了扩展字段，根据类型读取扩展字段
type MktHqSnapMsg struct {
	MsgHeader
	SecurityType      uint8
	TradSesMode       uint8
	TradeDate         uint32
	LastUpdateTime    uint32
	MDStreamID        [MDStreamID_LEN]byte
	SecurityID        [SecurityID_LEN]byte
	Symbol            [Symbol_LEN]byte
	PreClosePx        uint64
	TotalVolumeTraded uint64
	NumTrades         uint64
	TotalValueTraded  uint64
	TradingPhaseCode  [TradingPhase_LEN]byte
}

//实现MDGWMsg接口
func (mktHqSnapMsg *MktHqSnapMsg) GetMsgType() [MsgType_LEN]byte {
	return mktHqSnapMsg.MsgType
}

//因为行情快照包含了扩展字段，非定长结构体
//需要基于消息头部中长度以及扩展字段解析
type MktHq struct {
	hqbufer *bytes.Buffer
}

func (MktHq *MktHq) GetMsgType() [MsgType_LEN]byte {
	return [MsgType_LEN]byte{'M', '1', '0', '2'}
}

//指数行情快照
//根据条目个数需要进行扩展
type IndexSnapExt struct {
	NoMDEntries uint16
	MDEntryType [MDEntryType_LEN]byte
	MDEntryPx   uint64
}

//竞价行情快照
type BidSnapExt struct {
	NoMDEntries       uint16
	MDEntryType       [2]byte
	MDEntryPx         uint64
	MDEntrySize       uint64
	MDEntryPositionNo uint8
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
func calLoginMsgChkSum(loginMsg *LoginMsg) *bytes.Buffer {
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
	chksum := mdgwutils.CalCheckSum(buf.Bytes(), MSGHEADER_LEN+LOGINMSG_BODY_LEN)
	loginMsg.CheckSum = chksum
	//写入校验和
	binary.Write(buf, binary.BigEndian, loginMsg.CheckSum)
	return buf
}

//创建登录消息,参数文件当前时间和消息编号
func NewLoginMsg(sendingTtime, msgSeq uint64) (*LoginMsg, *bytes.Buffer) {
	loginMsg := &LoginMsg{}
	//初始化登录消息
	initLoginMsg(loginMsg)
	//填充消息体
	setLoginMsgBody(loginMsg)
	//填充消息头部
	setLoginMsgHeader(loginMsg, sendingTtime, msgSeq)
	//计算数据包校验和，并填充校验值
	buf := calLoginMsgChkSum(loginMsg)
	return loginMsg, buf
}

func setQueNoticeMsgHeader(queueNoticeMsg *QueueNoticeMsg) {
	//填充消息类型
	var setStr []byte
	setStr = []byte(QUEUE_NOTICE_TYPE)
	for i, c := range setStr {
		queueNoticeMsg.MsgType[i] = c
	}

	//填充消息序号
	queueNoticeMsg.MsgSeq = 0
	//发送时间为0
	queueNoticeMsg.SendingTtime = 0
	//消息体长度为0
	queueNoticeMsg.BodyLength = 0

}

func calQueNoticeMsgChkSum(queueNoticeMsg *QueueNoticeMsg) *bytes.Buffer {
	//将数据包中的字段放入到字节数组中，计算校验和
	buf := new(bytes.Buffer)

	//写入消息类型
	buf.Write(queueNoticeMsg.MsgType[:])

	//按照大端方式写入发送时间
	binary.Write(buf, binary.BigEndian, queueNoticeMsg.SendingTtime)
	//按照大端方式写入发送序号
	binary.Write(buf, binary.BigEndian, queueNoticeMsg.MsgSeq)
	//写入消息体长度
	binary.Write(buf, binary.BigEndian, queueNoticeMsg.BodyLength)

	//计算校验和
	chksum := mdgwutils.CalCheckSum(buf.Bytes(), MSGHEADER_LEN)
	queueNoticeMsg.CheckSum = chksum
	//写入校验和
	binary.Write(buf, binary.BigEndian, chksum)

	return buf
}

//创建解析线程退出消息
func NewQueueNoticeMsg() (*QueueNoticeMsg, *bytes.Buffer) {
	queueNoticeMsg := &QueueNoticeMsg{}

	//填充消息头部
	setQueNoticeMsgHeader(queueNoticeMsg)
	//计算数据包校验和，并填充校验值
	buf := calQueNoticeMsgChkSum(queueNoticeMsg)

	return queueNoticeMsg, buf
}

func setHeartBtMsgHeader(heartBtMsg *HeartBtMsg, sendingTtime, msgSeq uint64) {
	//填充消息类型
	//填充消息类型
	var setStr []byte
	setStr = []byte(QUEUE_NOTICE_TYPE)
	for i, c := range setStr {
		heartBtMsg.MsgType[i] = c
	}

	//填充消息序号
	heartBtMsg.MsgSeq = msgSeq
	//发送时间为0
	heartBtMsg.SendingTtime = sendingTtime
	//消息体长度为0
	heartBtMsg.BodyLength = 0
}

func calHeartBtMsgChkSum(heartBtMsg *HeartBtMsg) *bytes.Buffer {
	//将数据包中的字段放入到字节数组中，计算校验和
	buf := new(bytes.Buffer)

	//写入消息类型
	buf.Write(heartBtMsg.MsgType[:])
	//按照大端方式写入发送时间
	binary.Write(buf, binary.BigEndian, heartBtMsg.SendingTtime)
	//按照大端方式写入发送序号
	binary.Write(buf, binary.BigEndian, heartBtMsg.MsgSeq)
	//写入消息体长度
	binary.Write(buf, binary.BigEndian, heartBtMsg.BodyLength)

	//计算校验和
	chksum := mdgwutils.CalCheckSum(buf.Bytes(), MSGHEADER_LEN)
	heartBtMsg.CheckSum = chksum
	//写入校验和
	binary.Write(buf, binary.BigEndian, chksum)

	return buf

}

//创建一个心跳消息
func NewHeartBtMsg(sendingTtime, msgSeq uint64) (*HeartBtMsg, *bytes.Buffer) {
	heartBtMsg := &HeartBtMsg{}

	//填充消息头部
	setHeartBtMsgHeader(heartBtMsg, sendingTtime, msgSeq)

	//计算数据包校验和，并填充校验值
	buf := calHeartBtMsgChkSum(heartBtMsg)

	return heartBtMsg, buf
}

func GetMsgHeader(msgHeader *MsgHeader, b []byte, len int) {
	//获取消息头部
	buf := bytes.NewReader(b)
	fmt.Println("mesage header buf:", buf.Len())

	err := binary.Read(buf, binary.BigEndian, msgHeader)

	if err != nil {
		fmt.Println("get message header failed:", err)
	}
}

//通过byte序列，获取一个消息
func GetMsgFromBytes(b []byte, msglen int) MDGWMsg {

	buf := bytes.NewReader(b)
	//根据消息类型，登录成功消息
	if bytes.Equal(b[:MsgType_LEN], []byte(LOGINMSG_TYPE)) {
		loginMsg := &LoginMsg{}
		fmt.Println("it's login msg")
		err := binary.Read(buf, binary.BigEndian, loginMsg)
		if err != nil {
			fmt.Println("read msg from bytes err:", err)
		}
		return loginMsg
	} else if bytes.Equal(b[:MsgType_LEN], []byte(LOGOUTMSG_TYPE)) {
		logoutMsg := &LogoutMsg{}
		fmt.Println("it's logout msg")
		err := binary.Read(buf, binary.BigEndian, logoutMsg)
		if err != nil {
			fmt.Println("read msg from bytes err:", err)
		}
		return logoutMsg
	} else if bytes.Equal(b[:MsgType_LEN], []byte(HEARTBTMSG_TYPE)) {
		htMsg := &HeartBtMsg{}
		fmt.Println("it's heartbeat msg")

		err := binary.Read(buf, binary.BigEndian, htMsg)
		if err != nil {
			fmt.Println("it's heartbeart msg")
		}
		return htMsg
	} else if bytes.Equal(b[:MsgType_LEN], []byte(MKTSTUSMSG_TYPE)) {
		mktStatusMsg := &MktStatusMsg{}

		fmt.Println("it's market status msg")

		err := binary.Read(buf, binary.BigEndian, mktStatusMsg)
		if err != nil {
			fmt.Println("it's market status msg")
		}
		return mktStatusMsg
	} else if bytes.Equal(b[:MsgType_LEN], []byte(MKTSNAPMSG_TYPE)) {
		fmt.Println("it's market hq message")
		mktHq := &MktHq{}
		mktHq.hqbufer = bytes.NewBuffer(b)

		return mktHq
	}

	return nil

}

//解析消息
func ParseMsg(mdgwMsg MDGWMsg) int {
	var iRet int
	switch v := mdgwMsg.(type) {
	case *LoginMsg:
		fmt.Println("verify mdgw get msg is loginMsg", v.MsgType)
		iRet = LOGINMSG_TYPE_INT
		datas.ProcLoginMsg(v)
	case *LogoutMsg:
		fmt.Println("verify mdgw get msg is logoutMsg", v.MsgType)
		iRet = LOGOUTMSG_TYPE_INT
		datas.ProcLogoutMsg(v)
	case *MktStatusMsg:
		fmt.Println("market status msg:", v.MsgType)
		iRet = MKTSTATUS_TYPE_INT
		datas.ProcMtStatusMsg(v)
	case *MktHqSnapMsg:
		fmt.Println("snap hq msg:", v.MsgType)
		iRet = MKTHQ_TYPE_INT
		datas.ProcHqSnapMsg(v)
	case *MktHq:
		fmt.Println("... mkt hq ")
		datas.ProcMktHqMsg(v)
	default:
		fmt.Printf("other msg type")
		iRet = OTHERMSG_TYPE_INT
	}

	return iRet
}

//判断是一个完整的消息，返回消息的长度
//等于0，是不足一个消息
//大于0，是一个完成的消息，（可能是业务消息，也可能控制消息）
func IsFullMessage(b *bytes.Buffer) int {
	//检查长度
	var msglen int
	if b.Len() < MSGHEADER_LEN {
		msglen = 0
	} else if b.Len() == MSGHEADER_LEN {
		//header消息，也可能是控制消息
		return b.Len()
	} else { //收到的消息长度大于消息头部
		var testByte []byte
		copy(testByte, b.Bytes()[:MSGHEADER_LEN])

		//获取消息头部
		msgHeader := &MsgHeader{}
		GetMsgHeader(msgHeader, testByte, MSGHEADER_LEN)
		//获取消息体长度，然后检查buffer中的字节序列的长度是否大于等于消息的长度
		msglen := msgHeader.BodyLength + MSGHEADER_LEN + MSGTAIL_LEN

		if int(msglen) <= b.Len() {
			return int(msglen)
		}
	}

	return msglen
}
