type uam struct {
    log map[int]bool
    am int
}

func findingUsersActiveMinutes(logs [][]int, k int) []int {
    m := map[int]uam{}
    for i:=0;i<len(logs);i++{
        if val, ok := m[logs[i][0]]; ok {
            if _, ok := val.log[logs[i][1]]; !ok{
                val.log[logs[i][1]] = true
                val.am += 1
                m[logs[i][0]] = val
            }
        }else{
            m[logs[i][0]] = uam{
                log: map[int]bool{
                    logs[i][1] : true,
                },
                am: 1,
            }
        }
    }
    ret := make([]int, k)
    for _, v := range m {
        ret[v.am - 1] += 1
    }
    return ret
}