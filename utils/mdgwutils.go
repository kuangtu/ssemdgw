package mdgwutils

//计算校验和
func CalCheckSum(byte, uint32 u32Len) uint32{
	uint8 u8CheckSum = 0
	uint32 u32Cnt = 0

	for u32Cnt = 0; u32Cnt < u32Len; u32Cnt++ {
		CalCheckSum += (uint8)byte[u32Cnt]
	}

	return (uint32)u8ChkeckSum
}
