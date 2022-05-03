func getOnes(s string)int {
    count := 0
    for _, c := range s {
        if c == '1' {
            count++
        }
    }
    return count
}

func numberOfBeams(bank []string) int {
    sum :=0
    mult := -1
    for i:=0;i<len(bank);i++{
        c := getOnes(bank[i])
        if mult == -1 {
            mult = c
        }else if c != 0{
            sum += mult*c
            mult = c
        }
    }
    return sum
}