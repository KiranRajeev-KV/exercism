package anagram

import (
    "sort"
    "strings"
    )


func sortString(s string) string {
    // converting to []rune for easier sorting.
    r := []rune(strings.ToLower(s))
    
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

    return string(r)
}

func Detect(subject string, candidates []string) []string {
    sortedSubject := sortString(subject)
    lowerSubject := strings.ToLower(subject)
    var res []string

    for _,v := range(candidates) {
        if strings.ToLower(v) == lowerSubject {
			continue
		}
        if (sortString(v) == sortedSubject) {
            res = append(res,v)
        }
    }

	return res
}
