func findTheWinner(n int, k int) int {
    p := make([]int, n)
    for i:=0;i<n;i++{
        p[i] = i+1
    }
    ind := 0
    for;n!=1;{
        ind = (ind + k)%n
        if ind == 0 {
            ind = n-1
        }else {
            ind = ind -1
        }
        if ind == 0 {
            p = p[1:]
        }else{
            if ind != n-1 {
               p = append(p[:ind], p[ind+1:]...) 
            }else {
                p = p[:ind]
            }
        }
        n--
        if n == 1 {
            break
        }
    }
    return p[0]
}