import "strings"
func minSteps(s string, t string) int {
    m := make([]int, 26)
    for _, r := range s {
        m[int(r)-int('a')] += 1 
    }
    count := 0
    for _, r := range t {
        if m[int(r)-int('a')] > 0{
            m[int(r)-int('a')] -= 1
        }else {
            count++
        } 
    }
    if count == 0 {
        for _, v := range m {
            if v > 0 {
                count += v
            }
        } 
    }
    
    return count
}