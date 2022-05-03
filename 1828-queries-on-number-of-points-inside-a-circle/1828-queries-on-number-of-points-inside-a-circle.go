func countPoints(points [][]int, queries [][]int) []int {
    ret := []int{}
    for _, q := range queries {
        count := 0
        for _, p := range points {
            if ((p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])) <= q[2]*q[2] {
                count++
            }
        }
        ret = append(ret, count)
    }
    return ret
}