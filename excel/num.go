package excel

import "strconv"

func StringToFloat(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}
