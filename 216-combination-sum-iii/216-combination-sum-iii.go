func findSum(nums []int, sum int)[][]int{
    i, j:=0, len(nums)-1
    ret := [][]int{}
    for ;i<j; {
        temp := nums[i] + nums[j]
        if temp == sum {
            ret = append(ret, []int{nums[i], nums[j]})
            i++
            j--
        }else if temp > sum {
            j--
        }else{
            i++
        }
    }
    return ret
}

func findSumFor(nums []int, n, k int) [][]int{
    if n <= 0 || len(nums) == 0  || k == 0 || (n < k ){
        return nil
    }
    if k == 2 {
        return findSum(nums,n)
    }
    ret := [][]int{}
    for i:=0;i<=len(nums)-k;i++{
        newSum := n - nums[i]
        if newSum == 0 && k == 1{
            return [][]int{[]int{nums[i]}}
        }
        temp := findSumFor(nums[i+1:], newSum, k-1)
        if temp != nil && len(temp) > 0 {
            for j:=0;j<len(temp);j++{
                ret = append(ret, append(temp[j], nums[i]))
            }
        }
    }
    return ret
}

func combinationSum3(k int, n int) [][]int {
    nums := []int{}
    for i:=0;i<9;i++{
        nums = append(nums, i+1)
    }
    return findSumFor(nums, n, k)
}