package isbnverifier

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	clean := strings.ReplaceAll(isbn, "-", "")

	if len(clean) != 10 {
		return false
	}

	sum := 0

	for i := 0; i < 10; i++ {
		var n int

		if clean[i] == 'X' {
			// X is only valid as the check digit
			if i != 9 {
				return false
			}
			n = 10
		} else {
			v, err := strconv.Atoi(string(clean[i]))
			if err != nil {
				return false
			}
			n = v
		}

		sum += n * (10 - i)
	}

	return sum%11 == 0
}