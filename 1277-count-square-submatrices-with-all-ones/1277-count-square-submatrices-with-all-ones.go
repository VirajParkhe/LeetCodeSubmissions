func countSquares(matrix [][]int) int {
    l := len(matrix)
    if l > len(matrix[0]) {
        l = len(matrix[0])
    }
    r := len(matrix)
    c := len(matrix[0])
    for i:=1;i<l;i++{
        for k:=i;k<r;k++{
            for m:=i;m<c;m++{
                if matrix[k-1][m-1] >=i && matrix[k-1][m] >=i && matrix[k][m-1] >= i && matrix[k][m] != 0{
                    matrix[k][m] = i+1
                }
            }
        }
    }
    count := 0
    for i:=0;i<r;i++{
        for j:=0;j<c;j++{
            if matrix[i][j] > 0 {
                count += matrix[i][j]   
            } 
        }
    }
    return count
}