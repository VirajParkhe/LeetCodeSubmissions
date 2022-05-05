func minOperations(n int) int {
    sum := 0
    mid := 2*(n/2)+1
    if n%2 ==0 {
        mid =n
    }
    for i:=0;i<n/2;i++{
        sum += mid - (2*i +1)
    }
    return sum
}