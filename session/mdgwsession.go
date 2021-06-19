package mdgwsession

import (
	_ "ssevss/message"
)

func NewHeader(msgTytpe [4]byte, SendingTtime uint64, MsgSeq uint64, BodyLength uint32) (msgHeader *MsgHeader) {

}

//创建登录消息
func NewLoginMsg(senderCompID, targetCompID string, heartBtInt int, applVerID [8]byte) (loginMsg *LoginMsg) {
	loginMsg = &LoginMsg{}
	//创建头部
}
