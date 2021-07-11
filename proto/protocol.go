package proto

import (
	sess "ssevss/session"
	sock "ssevss/socket"
)

const (
	LOGINMDGW_OK  = 0
	CONN_FAILED   = -1
	VERIFY_FAILED = -2
)

//连接网关进行验证
func LoginMdgw(mdgwAddr string) int {

	//socket连接网关
	raddr := sock.NewSockAddr(mdgwAddr)
	mdgwSess := sess.NewMdgwSession(raddr)

	//连接
	rconn, err := sock.ConnGateWay(mdgwSess.Raddr)

	//socket 连接失败
	if err != nil {
		return CONN_FAILED
	}

	//设置rconn
	mdgwSess.SetSessionConn(rconn)

	return LOGINMDGW_OK
}

func VerifyMDGW() {

}
