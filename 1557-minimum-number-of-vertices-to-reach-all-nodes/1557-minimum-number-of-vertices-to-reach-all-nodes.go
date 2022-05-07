type edge struct {
    in int
    out int
}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
    m := map[int]edge{}
    var in, out int
    for i:=0;i<len(edges);i++{
        out = edges[i][0]
        in = edges[i][1]
        if val, ok := m[in]; !ok{
            m[in] = edge{1,0}
        }else{
            val.in +=1
            m[in]= val
        }
        if val, ok := m[out];!ok{
            m[out] = edge{0,1}
        }else{
            val.out +=1
            m[out] = val
        }
    }
    ret := []int{}
    for k, v := range m {
        if v.in == 0 {
            ret = append(ret, k)
        }
    }
    return ret
}