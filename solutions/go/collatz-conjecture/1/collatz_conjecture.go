package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
    if(n<=0) {
        return -1, errors.New("n must be a positive integer")
    }
    cnt := 0
    for {
        if(n == 1) {
            return cnt,nil
        } else if(n%2==0) {
            n = n/2
            cnt++
        } else {
            n = n*3 + 1
            cnt++
        }
    }
}
