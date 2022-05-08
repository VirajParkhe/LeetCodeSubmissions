func countTriplets(arr []int) int {
    l := len(arr)
    count := 0
    for i:=0;i<=l-2;i++{
        xor1:= arr[i]
        for j:=i+1;j<=l-1;j++{
            xor1^=arr[j]
            xor2:=arr[j]
            for k:=j;k<=l-1;k++{
                xor2^=arr[k]
                if xor2 == xor1 {
                    count++      
                }
            }
        }
    }
    return count
}