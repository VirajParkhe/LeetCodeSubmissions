
func convertToTitle(columnNumber int) string {
    ret := ""
    for ;;{
        if columnNumber < 27 {
            ret = string(columnNumber + int('A') - 1) + ret
            break
        }
        if columnNumber %26 == 0{
            ret = "Z" + ret
        }else{
           ret = string(columnNumber%26 + int('A') - 1) + ret 
        }
        v := columnNumber % 26 
        if v == 0 {
            v = 26
       }
        columnNumber = (columnNumber -v )/ 26
    }
    return ret
}

