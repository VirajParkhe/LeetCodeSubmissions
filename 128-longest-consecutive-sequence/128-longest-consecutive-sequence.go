func longestConsecutive(nums []int) int {
    m := map[int]int{}
    mDone := map[int]int{}
    for i:=0;i<len(nums);i++{
        if _, ok := m[nums[i]-1];ok{
            m[nums[i]] = nums[i]-1
        }else{
            m[nums[i]] = nums[i]
        }
        if _, ok := m[nums[i]+1]; ok {
            m[nums[i]+1] = nums[i]
        }
    }
    max := 0
    for i:=0;i<len(nums);i++{
        if _, ok := mDone[nums[i]]; ok{
            continue
        }
        val := nums[i]
        //fmt.Println("Picking: ", val, nums[i], m)
        count := 1
        for ; val!=m[val]; {
            //fmt.Println("Checking: ", val, mDone)
            if v, ok := mDone[val]; ok{
                count += v -1
                mDone[nums[i]] = count
                //fmt.Println("Setting: ", mDone, nums[i], count, v)
                break
            }else {
                //fmt.Println("Moving: ", val, m)
                val = m[val]
                count++
            }
        }
        mDone[nums[i]] = count
        //fmt.Println("Done: ", mDone, nums[i])
        if count > max {
            //fmt.Println("Setting Max: ", max, count)
            max = count
        }
    }
    //fmt.Println(mDone, m)
    return max
}