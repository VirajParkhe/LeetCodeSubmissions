import "strings"
import "strconv"
func getSubDomains(s string)[]string{
    ret := []string{}
    temp := ""
    splitted := strings.Split(s, ".")
    for i:=len(splitted)-1; i>=0;i--{
        if temp == ""{
            temp = splitted[i]
        }else{
         temp = splitted[i] + "." + temp   
        } 
        ret = append(ret, temp)
    }
    return ret
}

func subdomainVisits(cpdomains []string) []string {
    counts := map[string]int{}
    ret := []string{}
    for _, cpDomain := range cpdomains {
        splitted := strings.Split(cpDomain, " ")
        count, domain := splitted[0], splitted[1] 
        for _, subDomain := range getSubDomains(domain) {
            countInt, _ := strconv.Atoi(count)
            if _ , ok := counts[subDomain]; ok {
                counts[subDomain] += countInt
            }else{
                counts[subDomain] = countInt
            }
        }
    }
    for k, v := range counts {
        ret = append(ret, strconv.Itoa(v) + " " + k)
    }
    return ret
}