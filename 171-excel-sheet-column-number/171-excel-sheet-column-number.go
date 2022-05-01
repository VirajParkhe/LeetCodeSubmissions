
func titleToNumber(columnTitle string) int {
    ret := 0
    count:=0
    for i:=len(columnTitle) -1 ;i>=0;i--{
        val := int(columnTitle[i]) - int( 'A' ) + 1
        ret += val*int(math.Pow(float64(26), float64(count)))
        count++        
    }
    return ret
}