package luhn

import (
	"strconv"
	"strings"
)

func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func Valid(s string) bool {
	var rev_str []string
	x := []int{}
	//remove all spaces
	s = strings.ReplaceAll(s, " ", "")
	s = Reverse(s)
	rev_str = strings.Split(s, "")
	if len(s) < 2 {
		return false
	}

	for _, v := range rev_str {
		// check to see if we have a digit [0-9]
		if strings.ContainsAny("0123456789", string(v)) {
			t, _ := strconv.Atoi(string(v))
			x = append(x, t)
		} else {
			return false
		}
	}
	i := 1
	for i <= len(x)-1 {
		x[i] *= 2
		if x[i] > 9 {
			x[i] = x[i] - 9
		}
		i += 2
	}
	//Now add up x
	sum := 0
	for _, v := range x {
		sum += v
	}

	if sum%10 == 0 {
		return true
	} else {
		return false
	}

}
