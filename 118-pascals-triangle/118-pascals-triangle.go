func generate(numRows int) [][]int {
    prev :=  []int{1,1}
    ret := [][]int{[]int{1}}
    if numRows == 1 {
        return ret
    }
    ret = append(ret, prev)
    if numRows == 2 {
        return ret
    }
    for i:=0;i<numRows-2;i++{
        temp := []int{1}
        j,k :=0,1
        for ;k<len(prev);{
            temp = append(temp, prev[j]+prev[k])
            j++
            k++
        }
        temp = append(temp,1)
        ret = append(ret, temp)
        prev = temp
    }
    return ret
}