func wordPattern(pattern string, s string) bool {
    m1 := map[string]string{}
    m2 := map[string]string{}
    splitted := strings.Split(s, " ")
    l1 := len(splitted)
    l2 := len(pattern)
    if l1 != l2 {
        return false
    }
    for i, word := range splitted {
        if (m1[string(pattern[i])] != word || m2[word] != string(pattern[i])) &&  ( m2[word]!="" || m1[string(pattern[i])] != "" ){
            return false
        }
        m1[string(pattern[i])] = word
        m2[word] = string(pattern[i])
    }
    return true
}