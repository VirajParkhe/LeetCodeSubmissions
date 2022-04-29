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
        k :=1
        for ;k<i+2;{
            temp = append(temp, ret[i+1][k-1]+ret[i+1][k])
            k++
        }
        temp = append(temp,1)
        ret = append(ret, temp)
    }
    return ret
}