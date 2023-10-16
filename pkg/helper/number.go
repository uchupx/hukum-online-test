package helper

import "strconv"

// stringToInt64 converts string to int64
func StringToInt64(s string, def int64) int64 {
	if val, ok := strconv.ParseInt(s, 10, 64); ok == nil {
		return val
	}

	return def
}
