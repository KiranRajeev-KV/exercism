package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("length of given strings are not equal")
    }
    distance := 0
    for i := range(len(a)) {
        if a[i]!=b[i] {
            distance++
        }
    }

    return distance,nil
}
