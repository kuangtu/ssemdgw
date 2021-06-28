package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	vssconf "ssevss/configs"
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

	fmt.Println("the config file name is:", confile)

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

	//创建MdgwSession
	raddr := sock.NewSockAddr(vssconf.VssConf.Gatewayip)
	sess := sess.NewMdgwSession(&raddr)

	ret := sess.ConnMDGW()

	fmt.Println("the connMdgw ret is:", ret)
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
