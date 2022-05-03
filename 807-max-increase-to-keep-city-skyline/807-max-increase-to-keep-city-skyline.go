func maxIncreaseKeepingSkyline(grid [][]int) int {
    m := map[int]int{}
    m2 := map[int]int{}
    rows := len(grid)
    cols := len(grid[0])
    for i:=0;i<rows;i++{
        max := grid[i][0]
        for j:=0;j<cols;j++{
            if grid[i][j] > max {
                max = grid[i][j]
            } 
        }
        m[i] = max
    }
    
     for j:=0;j<cols;j++{
        max := grid[0][j]
        for i:=0;i<rows;i++{
            if grid[i][j] > max {
                max = grid[i][j]
            } 
        }
        m2[j] = max
    }
    count := 0
    for i:=0;i<rows;i++{
        for j:=0;j<cols;j++{
            min := int(math.Min(float64(m[i]), float64(m2[j])))
            if grid[i][j] < min {
                count+= min - grid[i][j]  
            }
        }
    }
    return count
}