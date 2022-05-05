import "sort"
func minPairSum(nums []int) int {
    sort.Ints(nums)
    j := len(nums)-1
    max := nums[0] + nums[j]
    for i:=0;i<j;i++{
        if nums[i] + nums[j] > max {
            max = nums[i] + nums[j]
        }
        j--
    }
    return max
}