func getRow(rowIndex int) []int {
    nC := 1
    temp := []int{nC}
    for i:=1;i<=rowIndex;i++{
        
        temp = append(temp, (nC * (rowIndex - i+1))/i)
        nC = (nC * (rowIndex - i+1))/i
    }
    return temp
}