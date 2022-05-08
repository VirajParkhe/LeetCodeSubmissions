func matrixBlockSum(mat [][]int, k int) [][]int {
    l := len(mat)
    m := len(mat[0])
    ret := [][]int{}
    for i:=0;i<l;i++{
        temp := []int{}
        for j:=0;j<m;j++{
            val := 0
            for o := i-k;o<=i+k;o++{
                for p:= j-k;p<=j+k;p++{
                    if o>=0 && o<=l-1 && p>=0 && p<=m-1 {
                        val += mat[o][p]
                    }
                }
            }
            temp = append(temp, val)
            
        }
        ret = append(ret, temp)
    }
    return ret
}