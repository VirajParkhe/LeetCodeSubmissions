func reverseBits(num uint32) uint32 {
    count := 0
    var newNum uint32
    for  i:=31;i>=0;i-- {
        newNum += ((num&(1<<i))>>i)<<count
        count++
    }
    return uint32(newNum)
}