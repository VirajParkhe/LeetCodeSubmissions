func isValid(board [][]byte, row, col int)bool{
    val := board[row][col]
    for i:=0;i<9;i++{
        if i !=row{
            if val == board[i][col]{
                return false
            }
        }
    }
    for i:=0;i<9;i++{
        if i !=col{
            if val == board[row][i]{
                return false
            }
        }
    }
    t1 := row/3
    t2 := col/3
    var startr,endr,startc,endc int
    startr = t1*3
    endr = (t1+1)*3 -1
    startc = t2*3
    endc = (t2+1)*3 -1
    for i:=startr;i<=endr;i++{
        for j:=startc;j<=endc;j++{
            if i!=row && j!=col{
                if val == board[i][j]{
                    return false
                }
            }
        }
    }
    return true
}
func isValidSudoku(board [][]byte) bool {
    for i:=0;i<len(board);i++{
        for j:=0;j<len(board[i]);j++{
            if board[i][j] != 46 {
                ok:= isValid(board,i,j)
                if !ok {
                    return ok
                }
            }
        }
    }
    return true
}