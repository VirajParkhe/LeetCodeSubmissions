func resolveStr(s string)string {
    str1 := []string{}
    for _, r := range s {
        val := string(r)
        if val == "#" {
            l := len(str1)
            if l != 0 {
                str1 = str1[:l - 1]
            }
        }else{
            str1 = append(str1, val)
        }
    }
    ret := ""
    for _, s := range str1 {
        ret += s
    }
    return ret
}
func backspaceCompare(s string, t string) bool {
    if resolveStr(s) ==  resolveStr(t) {
        return true
    }
    return false
}