package mdgwutils

//计算校验和
func CalCheckSum(buffer []byte, u32Len uint32) uint32 {
	var u32Cnt uint32
	var chkSum uint8
	for u32Cnt = 0; u32Cnt < u32Len; u32Cnt++ {
		chkSum += uint8(buffer[u32Cnt])
	}

	return uint32(chkSum)
}
