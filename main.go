package main

import (
    "fmt"
    "ssevss/sysconfig"
)

func main() {
    var sysconf = SysConf{"GateWayPort": 10}
    fmt.Println("the gateway port is:%s", sysconf.GateWayPort)
    //fmt.Println("ssevss:%d", i)

}