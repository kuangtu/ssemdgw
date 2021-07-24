package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type SysConf struct {
	Gatewayip    string `json:"GateWayIP:port"`
	Localip      string `json:"LocalIP:port"`
	Backdir      string `json:"BackDir"`
	SenderCompID string `json:"SenderCompID"`
	TargetCompID string `json:"TargetCompID"`
	HeaderBtInt  int    `json:"HeaderBtInt"`
	ApplVerID    string `json:"ApplVerID"`
	RetryTime    int    `json:"RetryTime"`
}

//全局变量，系统配置信息
var VssConf SysConf

//读取配置文件
func ReadSysConf(confPath string) int {
	fmt.Println("start read sys config file.")
	//打开文件进行读取
	jsonfile, err := os.Open(confPath)
	if err != nil {
		fmt.Println("open config file error:", err)

		return -1
	}
	defer jsonfile.Close()

	//读取json文件
	byteBuffer, err := ioutil.ReadAll(jsonfile)
	//fmt.Println(string(byteBuffer))
	err = json.Unmarshal(byteBuffer, &VssConf)
	if err != nil {
		fmt.Println("unmarshal json file error:", err)
		return -1
	}

	fmt.Println("the gatewayip is:", VssConf.Gatewayip)
	fmt.Println("the retry time is:", VssConf.RetryTime)

	return 0
}
