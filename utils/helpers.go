package utils

import "strconv"

func ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}
