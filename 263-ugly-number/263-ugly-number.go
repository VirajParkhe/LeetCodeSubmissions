func isUgly(n int) bool {
    if n ==1  {
        return true
    }
    if n == 0 {
        return false
    }
    for _, fact := range []int{2, 3, 5} {
        for ;n%fact ==0; {
            n = n/fact
            if n == 1 {
                return true
            }
        }
    }
    return false
}