
func getMaximumXor(nums []int, maximumBit int) []int {
    mask := (1<<maximumBit) - 1
    xors := []int{}
    temp := 0
    for i:=0;i<len(nums);i++{
        xors = append(xors, temp^nums[i]^mask)
        temp = temp^nums[i]
    }
    ret := []int{}
    for i:=len(nums)-1;i>=0;i--{
        ret = append(ret, xors[i])
    }
    return ret
}