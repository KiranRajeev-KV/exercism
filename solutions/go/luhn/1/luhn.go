package luhn

import "strings"
import "strconv"

func Valid(id string) bool {
    // trimmed := strings.TrimSpace(id)
     stripped := strings.ReplaceAll(id," ","")
    if len(stripped) <=1 {
        return false
    }
    sum := 0
    flag := false
    
    for i:= len(stripped)-1; i>= 0; i-- {
        n, err := strconv.Atoi(string(stripped[i]))
        if (err != nil) {
            return false
        }
        if(flag) {
            if(n*2 > 9) {
                sum += ((n*2)-9)
            } else {
                sum += (n*2)
            }
            flag = false
        } else {
            sum += n
            flag = true
        }
    }

	if (sum%10 ==0) {
        return true
    } else {
    	return false
    }
}
