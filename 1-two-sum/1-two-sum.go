//import "sort"

type num struct{
    value int
    ind  int
}

type foo []num

func (f foo)Len()int{
    return len(f)
}

func (f foo)Less(i, j int)bool{
    return f[i].value < f[j].value
}

func (f foo)Swap(i, j int){
    f[i], f[j] = f[j], f[i]
    
}

func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i:=0;i<len(nums);i++{
        if v, ok := m[target - nums[i]]; ok {
            return []int{i, v}
        }else{
            m[nums[i]] = i
        }
    }
    return []int{}
}