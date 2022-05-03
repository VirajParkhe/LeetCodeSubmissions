func getCount(boxes string, isLeft bool) int{
    count := 0
    l := len(boxes)
    for i, val := range boxes {
        if val == '1' {
            if isLeft {
                count += l - i 
            }else {
                count += i + 1
            }
        }
    }
    return count
}

func minOperations(boxes string) []int {
    l := len(boxes)
    ret := []int{}
    for i:=0;i<l;i++{
        lc, rc := 0,0
        lc = getCount(boxes[:i], true)
        if i != l-1{
            rc = getCount(boxes[i+1:], false)
        }
        ret = append(ret, lc+rc)
    }
    return ret
}