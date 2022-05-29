type node struct {
    s string
    ind int
}
func numOfPairs(nums []string, target string) int {
    m := map[int][]node{}
    for i:=0;i<len(nums);i++{
        m[len(nums[i])] = append( m[len(nums[i])], node{nums[i], i})
    }
    count := 0
    targetLen := len(target)
    for i:=0;i<len(nums);i++{
        if v, ok := m[targetLen - len(nums[i])]; ok {
            for _, s := range v {
                if nums[i] + s.s == target  && s.ind!=i{
                    count++
                }
            }
        }
    }
    return count
}