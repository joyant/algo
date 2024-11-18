package binary

/**
190. 颠倒二进制位

颠倒给定的 32 位无符号整数的二进制位。
*/

// 遍历，逐个取出最低位，成为 res 的最高位
func reverseBits(num uint32) uint32 {
    var res uint32
    for i := 0; i < 32; i++ {
        res = ((num & 1) << (31 - i)) | res
        num = num >> 1
    }
    return res
}
