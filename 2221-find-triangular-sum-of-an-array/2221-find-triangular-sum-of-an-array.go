func triangularSum(nums []int) int {
    end := len(nums)
    for;end!=1;{
        for i:=0;i<end-1;i++{
            nums[i] = (nums[i] + nums[i+1])%10
        }
        end--
    }
    return nums[0]
}