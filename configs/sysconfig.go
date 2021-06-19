package configs

import (
	_ "encoding/json"
)

type SysConf struct {
	Gatewayip    string `json:"GateWayIP"`
	Gatewayport  int    `json:"GateWayPort"`
	Localip      string `json:"LocalIP"`
	Localport    int    `json:"LocalPort"`
	Backdir      string `json:"BackDir"`
	SenderCompID string `json:"SenderCompID"`
	TargetCompID string `json:"TargetCompID"`
	HeaderBtInt  string `json:"HeaderBtInt"`
	ApplVerID    string `json:"ApplVerID"`
}
