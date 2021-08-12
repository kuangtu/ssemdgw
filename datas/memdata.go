package datas

type MemData struct {
	IndexMemHq map[IndexHq]string
	StockMemHq map[StkHq]string
}

var MemHqData MemData

//初始化内存行情
func InitMemHq() {
	MemHqData.IndexMemHq = make(map[IndexHq]string)
	MemHqData.StockMemHq = make(map[StkHq]string)
}

//更新指数行情
func (memHqData *MemData) UpdateIdxHq(idxHq IndexHq) {
	MemHqData.IndexMemHq[idxHq] = string(idxHq.SecurityID[:])

}

//更新证券行情
func (memHqData *MemData) UpdateStkHq(stkHq StkHq) {
	MemHqData.StockMemHq[stkHq] = string(stkHq.SecurityID[:])
}
