package excel

func GetColLetter(col int) string {
	s := int(float64(col / 26))
	f := col % 26
	str := string(65 + f)
	if s != 0 {
		str = string(64+s) + str
	}
	return str
}
