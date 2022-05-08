func match(s, p string, l int) bool {
    m1 := map[string]string{}
    m2 := map[string]string{}
    for i:=0;i<l;i++{
        if (m1[string(s[i])] != string(p[i]) && m1[string(s[i])] !="") || (m2[string(p[i])] != string(s[i]) && m2[string(p[i])] !="")  {

                 return false
        
        }
        m1[string(s[i])] = string(p[i])
        m2[string(p[i])] = string(s[i])
    }
    return true
}

func findAndReplacePattern(words []string, pattern string) []string {
    l := len(pattern)
    ret := []string{}
    for i:=0;i<len(words);i++{
        if match(words[i], pattern, l) {
            ret = append(ret, words[i])
        }
    }
    return ret
}