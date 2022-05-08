import "strings"
func minSteps(s string, t string) int {
    m := map[rune]int{}
    for _, r := range s {
        if v, ok := m[r];ok{
            m[r] = v+1
        }else{
            m[r] = 1
        }
    }
    count := 0
    for _, r := range t {
        if v, ok := m[r]; ok {
            m[r] = v -1
            if v - 1 == 0{
                delete(m, r)
            }
        }else {
            count++
        } 
    }
    if count == 0 {
        for _, v := range m {
            count += v
        } 
    }
    
    return count
}