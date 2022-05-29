import "strings"
func hash(a string)string {
    arr := make([]rune, 26)
    for i:=0;i<len(a);i++{
        arr[int(a[i])-int('a')] =1
    }
    ret := []string{}
    for i:=0;i<26;i++{
        if arr[i] == 1 {
            ret = append(ret, "1")
        }else{
            ret = append(ret, "0")
        }
    }
    return strings.Join(ret, "")
}

func compare(a, b string)bool{
    for i:=0;i<len(a);i++{
        if a[i] == b[i] && a[i] != '0'{
            return false
        }
    }
    return true
}

func maxProduct(words []string) int {
    mH := map[string]string{}
    mL := map[string]int{}
    for i:=0;i<len(words);i++{
        mH[words[i]] = hash(words[i])
        mL[words[i]] = len(words[i])
    }
    max := 0
    for i:=0;i<len(words)-1;i++{
        for j:=i+1;j<len(words);j++{
            if (mL[words[i]]*mL[words[j]] > max) && (compare(mH[words[i]], mH[words[j]])) {
                max = mL[words[i]]*mL[words[j]]
            }
        }
    }
    return max
}