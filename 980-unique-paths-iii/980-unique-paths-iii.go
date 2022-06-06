var valid int
func uniquePathsIII(grid [][]int) int {
    valid = 0
    start,end:=0,0
    r,c := len(grid), len(grid[0])
    for i:=0;i<r;i++{
        for j:=0;j<c;j++{
            if grid[i][j] == 1{
                start, end = i, j
            }
            if grid[i][j] == 0 {
                valid++
            }
        }
    }
    count := 0
    findPath(start, end, 0,r,c,&grid, &count)
    return count
}

func findPath(start, end, trav, r,c int, grid *[][]int, count *int) {
    if (*grid)[start][end] == 2 && trav == valid{
        *count++
        return
    }else if (*grid)[start][end] == 2 {
        return
    }
    
     if (*grid)[start][end] == 0 {
        trav++
        (*grid)[start][end] = -2
     }
   //fmt.Println(start, end, trav, valid,r,c, (*grid))
    if ( (start + 1) < r ) && (((*grid)[start+1][end] == 0 )|| ((*grid)[start+1][end] == 2)){
            findPath(start+1, end,trav,r,c, grid, count)
        }
    if (start - 1 >=0) && ((*grid)[start-1][end] == 0 || (*grid)[start-1][end] == 2){
            findPath(start-1,end,trav,r,c,grid, count)
        }
    if (end +1 < c) && ((*grid)[start][end+1] == 0 || (*grid)[start][end+1] == 2){
            findPath(start, end+1,trav,r,c, grid,count)
        }
    if (end - 1 >= 0) && ((*grid)[start][end-1] == 0|| (*grid)[start][end-1] == 2){
            findPath(start, end-1,trav,r,c, grid, count)
        }
    if (*grid)[start][end] == -2 {
        trav--
        (*grid)[start][end] = 0
    }
        
}