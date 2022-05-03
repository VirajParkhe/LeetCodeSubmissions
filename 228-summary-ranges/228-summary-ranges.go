func summaryRanges(nums []int) []string {
    l := len(nums)
    if l == 0 {
        return nil
    }
    ret := []string{}
    temp := strconv.Itoa(nums[0])
    for i:=1;i<l;i++{
        if nums[i] != nums[i-1]+1 && nums[i]!=nums[i-1]{
            if !strings.Contains(temp, strconv.Itoa(nums[i-1])){
                temp += "->" + strconv.Itoa(nums[i-1])
            } 
            ret = append(ret, temp)
            temp = strconv.Itoa(nums[i])
        }
    }
    if !strings.Contains(temp, strconv.Itoa(nums[l-1])){
        temp += "->" + strconv.Itoa(nums[l-1])
    }
    ret = append(ret, temp)
    return ret
}