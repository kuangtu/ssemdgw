package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	vssconf "ssevss/configs"
	msg "ssevss/message"
	sess "ssevss/session"
	sock "ssevss/socket"
)

var (
	confile = flag.String("f", "", "system config file path.")
)

func main() {
	fmt.Println("main function")
	flag.Parse()
	if *confile == "" {
		fmt.Println("confile path empty")
		os.Exit(1)
	}

	fmt.Println("the config file name is:", *confile)

	//打开文件进行读取
	jsonfile, err := os.Open(*confile)
	if err != nil {
		fmt.Println("open config file error:", err)
		os.Exit(1)
	}
	defer jsonfile.Close()

	//读取json文件
	byteBuffer, err := ioutil.ReadAll(jsonfile)
	fmt.Println(string(byteBuffer))

	err = json.Unmarshal(byteBuffer, &vssconf.VssConf)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("the gatewayip is:", vssconf.VssConf.Gatewayip)

	logingMsg, buf := msg.NewLoginMsg(3, 2)
	fmt.Println("the msg send time is:", logingMsg.SendingTtime)
	fmt.Println(string(buf.Bytes()))
	//创建MdgwSession
	raddr := sock.NewSockAddr(vssconf.VssConf.Gatewayip)
	sess := sess.NewMdgwSession(raddr)

	ret := sess.ConnMDGW()
	fmt.Println("connect ret is:", ret)
	if ret == -1 {
		fmt.Println("connect mdgw failed:", ret)
	}

	res := sess.VerifyMDGW()
	fmt.Println("verify res is:", res)

	if res == false {
		fmt.Println("verify mdgw failed.")

		return
	}

	//启动定时发送心跳消息的goroutine

	//验证通过之后，接收行情文件
	sess.Rconn.Close()

	// fmt.Println("the connMdgw ret is:", ret)

	// var Header mdgwmsg.MsgHeader
	// Header.SendingTtime = 1111
	// fmt.Println("the send time is:", Header.SendingTtime)

	// //连接MDGW行情网关
	// //从配置sysconfig中获取Gatewayip
	// raddr, err := net.ResolveTCPAddr("tcp", Sysconf.Gatewayip)

	// if err != nil {
	// 	logger.warn("Failed to resolve remote address:", err)
	// 	os.Exit(1)
	// }

}
