package helpers

import (
	"strconv"
	"strings"
)

func FormatRupiah(amount int64)string{
	s := strconv.FormatInt(amount, 10)

	n := len(s)
	if n <= 3{
		return "Rp" + s
	}
	var result []string

	for n > 3 {
		result = append([]string{s[n-3:]}, result...)
		s = s[:n-3]
		n = len(s)
	}

	if s != ""{
		result = append([]string{s} ,result...)
	}

	return "Rp" + strings.Join(result, ".")
}