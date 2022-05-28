import "math"
func pathInZigZagTree(label int) []int {
    if label == 1 {
        return []int{1}
    }
    level := int(math.Log2(float64(label)))
    start := int(math.Pow(float64(2), float64(level)))
    end :=  int(math.Pow(float64(2), float64(level+1))) -1
    path := []int{}
    isDesc:=true
    if level%2 == 0 {
        isDesc = false
    }
    if isDesc {
        start,end = end, start
    }
    for;;{
        if isDesc && start-1== end {
            break
        }else if !isDesc && start+1 == end {
            break
        }
        mid := float64((start+end)/2)
        fmt.Println(start, end, mid, isDesc, label)
        if float64(label) > mid {
            if isDesc {
                path = append(path, 0)
                end =  int(mid)
            }else{
                start = int(mid)
                path = append(path, 1)
            }
        }else {
            if isDesc {
                start = int(mid)
                path = append(path, 1)
            }else {
                end =  int(mid)
                path= append(path, 0)
            }
        }
        
    }
    if start == label {
        
            path = append(path, 0)
        
        
    }else{
       
            path = append(path, 1)
        
        
    }
    fmt.Println(path)
    ret := []int{1}
    start = 1
    isDesc = false
    level = 1
    for _, val := range path {
        if isDesc {
            isDesc = false
        }else{
            isDesc = true
        }
        left := 2*start
        right :=  2*start+1
        total := int(math.Pow(float64(2), float64(level))) + int(math.Pow(float64(2), float64(level+1))) -1
        fmt.Println(left, right, total, level, isDesc, val)
        level++
        if isDesc{
            if val == 0 {
                if total - left <= label {
                    ret = append(ret, total - left)
                }
                
                start= total-left
            }else {
                 if total - right <= label {
                ret = append(ret, total - right)
                 }
                start = total - right
            }
        }else {
            if val ==0 {
                 if total - right <= label {
                ret = append(ret, total-right)
                 }
                start = total-right
            }else{
                 if total - left <= label {
                ret = append(ret, total-left)
                 }
                start = total-left
            }
        }
    }
    
    fmt.Println(ret)
    return ret
    
}