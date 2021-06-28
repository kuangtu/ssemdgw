package configs

type SysConf struct {
	Gatewayip    string `json:"GateWayIP:port"`
	Localip      string `json:"LocalIP:port"`
	Backdir      string `json:"BackDir"`
	SenderCompID string `json:"SenderCompID"`
	TargetCompID string `json:"TargetCompID"`
	HeaderBtInt  string `json:"HeaderBtInt"`
	ApplVerID    string `json:"ApplVerID"`
}

var VssConf SysConf
