func countBattleships(board [][]byte) int {
    count := 0
    for i:=0;i<len(board);i++{
        for j:=0;j<len(board[0]);j++{
            if board[i][j] == 'X' {
                up := false
                left := false
                if i!=0 && board[i-1][j] == 'X' {
                    up = true
                }
                if j!=0  && board[i][j-1] == 'X' {
                    left = true
                }
                if !up && !left{
                    count++
                }
            }
        }
    }
    return count
}