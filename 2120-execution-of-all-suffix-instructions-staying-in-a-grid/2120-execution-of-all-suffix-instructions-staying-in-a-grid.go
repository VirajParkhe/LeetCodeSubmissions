func isInBounds(x,y,n int)bool{
    if x >=0 && x<=n-1 && y>=0 && y<=n-1 {
        return true
    }
    return false
}

func getNextVal(x, y int, op string) (int, int){
    if op == "U" {
        return x-1, y
    }
    if op == "R" {
        return x, y+1
    }
    if  op == "D" {
        return x+1, y
    }
    return x, y-1
}

func executeInstructions(n int, startPos []int, s string) []int {
    l := len(s)
    ret := []int{}
    for i:=0;i<l;i++{
        x, y := startPos[0], startPos[1]
        count:=0
        broken := false
        for j:=i;j<l;j++{
            xt, yt := getNextVal(x, y, string(s[j]))
            if !(xt >=0 && xt<=n-1 && yt>=0 && yt<=n-1) {
                broken = true
                ret = append(ret, count)
                break
            }
            x,y = xt, yt
            count++
        }
        if !broken {
            ret = append(ret, count)
        }
    }
    return ret
}