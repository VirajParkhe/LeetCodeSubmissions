
func getMaximumXor(nums []int, maximumBit int) []int {
    mask := (1<<maximumBit) - 1
    l := len(nums)
    xors := make([]int, l)
    temp := 0
    for i:=0;i<l;i++{
        xors[l-i-1] = temp^nums[i]^mask
        temp = temp^nums[i]
    }
    return xors
}