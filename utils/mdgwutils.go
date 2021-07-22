package mdgwutils

import (
	"strconv"
	"time"
)

//计算校验和
func CalCheckSum(buffer []byte, u32Len uint32) uint32 {
	var u32Cnt uint32
	var chkSum uint8
	for u32Cnt = 0; u32Cnt < u32Len; u32Cnt++ {
		chkSum += uint8(buffer[u32Cnt])
	}

	return uint32(chkSum)
}

//获取当前的系统时间,返回格式YYYYMMDDHHmmSSsss
func GetCurTime() uint64 {
	curTime := time.Now()
	//通过字符串Jan 2 15:04:05 2006 MST格式输出
	timeStr := curTime.Format("20060102150405000")
	timeU64, _ := strconv.ParseUint(timeStr, 10, 64)

	return timeU64

}

//测试时消除“定义后未使用”告警
func UNUSED(x ...interface{}) {

}
