type foo []rune

func (f foo)Less(i, j int)bool{
    return f[i] < f[j]
}

func (f foo)Len()int{
    return len(f)
}

func (f foo)Swap(i, j int){
    f[i], f[j] = f[j], f[i]
}

func groupAnagrams(strs []string) [][]string {
    m := map[string][]string{}
    for i:=0;i<len(strs);i++{
        runes := []rune(strs[i])
        sort.Sort(foo(runes))
        if val, ok := m[string(runes)]; ok{
            m[string(runes)] = append(val, strs[i])
        }else{
            m[string(runes)] = []string{strs[i]}
        }
    }
    ret := [][]string{}
    for _, val := range m{
        ret = append(ret, val)
    }
    return ret
}