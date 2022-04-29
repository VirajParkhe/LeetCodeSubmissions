func getRow(rowIndex int) []int {
    nC := 1
    temp := []int{nC}
    for i:=1;i<=rowIndex;i++{
        nCr := (nC * (rowIndex - i+1))/i
        temp = append(temp, nCr)
        nC = nCr
    }
    return temp
}