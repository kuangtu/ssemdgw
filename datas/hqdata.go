package datas

import (
	msg "ssevss/message"
)

//指数行情结构体
type IndexHq struct {
	SecurityID    [msg.SecurityID_LEN]byte
	Symbol        [msg.Symbol_LEN]byte
	LastPrc       uint64
	OpenPrc       uint64
	HighPrc       uint64
	LowPrc        uint64
	ClsPrc        uint64
	PreClsPrc     uint64
	TotalVolume   uint64
	TotalTurnover uint64
	TrdPahase     [msg.TradingPhase_LEN]byte
	LastTime      uint32
}

//竞价股票行情结构体
type StkHq struct {
	SecurityID    [msg.SecurityID_LEN]byte
	Symbol        [msg.Symbol_LEN]byte
	LastPrc       uint64
	OpenPrc       uint64
	HighPrc       uint64
	LowPrc        uint64
	ClsPrc        uint64
	PreClsPrc     uint64
	TotalVolume   uint64
	TotalTurnover uint64
	Iopv          uint64 //etf净值
	TrdPahase     [msg.TradingPhase_LEN]byte
	LastTime      uint32
}
