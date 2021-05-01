package sysconfig

import (
    _ "encoding/json"
)

type SysConf struct {
    GateWayIP string
    GateWayPort int
    LocalIP string
    LocalPort int
    BackDir string
}

