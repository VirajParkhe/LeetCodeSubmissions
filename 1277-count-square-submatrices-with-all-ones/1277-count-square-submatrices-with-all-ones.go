func countSquares(matrix [][]int) int {
    l := len(matrix)
    if l > len(matrix[0]) {
        l = len(matrix[0])
    }
    for i:=1;i<l;i++{
        for k:=i;k<len(matrix);k++{
            for m:=i;m<len(matrix[0]);m++{
                if matrix[k-1][m-1] >=i && matrix[k-1][m] >=i && matrix[k][m-1] >= i && matrix[k][m] != 0{
                    matrix[k][m] = i+1
                }
            }
        }
    }
    count := 0
    for i:=0;i<len(matrix);i++{
        for j:=0;j<len(matrix[0]);j++{
            if matrix[i][j] > 0 {
                count += matrix[i][j]   
            } 
        }
    }
    return count
}