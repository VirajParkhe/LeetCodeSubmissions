func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m := make([]int, 26)
    for i:=0;i<len(s);i++{
        m[int(s[i]) - int('a')] +=1 
    }
    for i:=0;i<len(t);i++{
        m[int(t[i]) - int('a')] -=1
    }
    for i:=0;i<26;i++{
        if m[i]!=0 {
            return false
        }
    }
    return true
}