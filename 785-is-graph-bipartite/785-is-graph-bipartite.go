func traverse(set []int, start, cSet int, graph [][]int) bool {
    if set[start] != -1 {
        return set[start] == cSet
    }    
    set[start] = cSet
    v := 0
    if cSet == 0 {
        v = 1
    }
    for _, node := range graph[start] {
        if !traverse(set, node, v, graph){
            return false
        }
    }
    return true
}

func isBipartite(graph [][]int) bool {
    set := make([]int, len(graph))
    for i:=0;i<len(graph);i++{
        set[i] = -1
    }
    for i:=0;i<len(set);i++{
        if set[i] == -1 {
            if !traverse(set, i, 0, graph){
                return false
            }
        }

    }
    return true
}