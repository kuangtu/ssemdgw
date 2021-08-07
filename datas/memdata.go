package datas

type MemData strcut {
	IndexMemHq map[IndexHq]string
	StockMemHq map[StkHq]string

}

var MemHqData MemData

//内存中的行情通过hash保存

//内存指数行情

//内存竞价行情

//初始化内存行情
func InitMemHq() {
	MemHqData.IndexMemHq =make(map[IndexHq]string)
	MemHqData.StockMemHq = make(map[StkHq]string)
}

func