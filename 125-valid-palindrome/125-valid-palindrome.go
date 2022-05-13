import "strings"
func isPalindrome(s string) bool {
    ret := []byte{}
    for i:=0;i<len(s);i++{
        if (s[i] >= 'a' && s[i]<='z') || (s[i] >= '0' && s[i] <='9'){
            ret = append(ret, s[i])
        }else if s[i] >='A' && s[i] <= 'Z'{
            ret = append(ret, strings.ToLower(string(s[i]))[0])
        }
    }
    i, j := 0, len(ret)-1
    for ;i <j;{
        if ret[i] != ret[j] {
            return false
        }
        i++
        j--
    }
    return true
}