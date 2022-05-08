var v = []string{"a", "e", "i", "o", "u"}
var count int
func construct(s string, n, start int){
    if n == 0 {
        count++
        return
    }
    for i:=start;i<5;i++{
        s+=v[i]
        construct(s, n-1, i)
    }
}

func countVowelStrings(n int) int {
    count = 0
    construct("", n, 0)
    return count
}