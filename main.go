package main

import (
	"flag"
	"fmt"
	vssconf "ssevss/configs"
	sess "ssevss/session"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	confile = flag.String("f", "", "system config file path.")
)

func main() {
	var iRet int
	logrus.SetLevel(logrus.TraceLevel)
	flag.Parse()
	if *confile == "" {
		// fmt.Println("confile path empty")
		// os.Exit(1)
		logrus.Fatal("config file is missing.")
	}

	//读取配置文件，进行解析
	iRet = vssconf.ReadSysConf(*confile)

	if iRet == -1 {
		// fmt.Println("read config file failed, and exit")
		// os.Exit(-1)
		logrus.Fatal("read config file failed.")
	}

	//连接MDGW网关
	//（1）如果连接网关失败，等待配置时间后重复进行连接
	//（2）如果连接网关正常，验证失败，程序退出
	//（3）如果连接网关正常，验证正常，socket网关读取中断，
	// 断开连接，等待配置时间后重复进行连接

	for {
		var wait sync.WaitGroup
		//心跳消息设置，退出quit信号
		var ticker = time.NewTicker(time.Duration(vssconf.VssConf.HeaderBtInt))
		var quit = make(chan bool)
		//连接网关，并进行验证
		fmt.Println("start to login mdgw.")
		iRet = sess.LoginMdgw(vssconf.VssConf.Gatewayip)

		if iRet == sess.CONN_FAILED {
			fmt.Println("connect gateway failed, retry later")
			//确认清理缓存的数据等
			continue
		} else if iRet == sess.MDGWVERIFY_FAILED {
			fmt.Println("mdgwverify failed, connect manual.")
			break
		} else {
			fmt.Println("login ok")
			//开始进行接收和解析
			wait.Add(3)
			go sess.RecvMdgwMsg(&wait, quit)
			go sess.ProcMdgwMsg(&wait)
			go sess.SendHeartBtMsg(&wait, ticker, quit)
			//等待接收、解析、发送心跳消息goroutine退出
			wait.Wait()
		}

		time.Sleep(time.Duration(vssconf.VssConf.RetryTime) * time.Second)
	}

	// logingMsg, buf := msg.NewLoginMsg(3, 2)
	// fmt.Println("the msg send time is:", logingMsg.SendingTtime)
	// fmt.Println(string(buf.Bytes()))
	// //创建MdgwSession
	// raddr := sock.NewSockAddr(vssconf.VssConf.Gatewayip)
	// sess := sess.NewMdgwSession(raddr)

	// ret := sess.ConnMDGW()
	// fmt.Println("connect ret is:", ret)
	// if ret == -1 {
	// 	fmt.Println("connect mdgw failed:", ret)
	// }

	// res := sess.VerifyMDGW()
	// fmt.Println("verify res is:", res)

	// if res == false {
	// 	fmt.Println("verify mdgw failed.")

	// 	return
	// }

	// //启动定时发送心跳消息的goroutine

	// //验证通过之后，接收行情文件
	// sess.Rconn.Close()

	// // fmt.Println("the connMdgw ret is:", ret)

	// // var Header mdgwmsg.MsgHeader
	// // Header.SendingTtime = 1111
	// // fmt.Println("the send time is:", Header.SendingTtime)

	// // //连接MDGW行情网关
	// // //从配置sysconfig中获取Gatewayip
	// // raddr, err := net.ResolveTCPAddr("tcp", Sysconf.Gatewayip)

	// // if err != nil {
	// // 	logger.warn("Failed to resolve remote address:", err)
	// // 	os.Exit(1)
	// // }

}
