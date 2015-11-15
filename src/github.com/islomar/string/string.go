package string

//Function can be exported and used by other packages
func Reverse(s string) string {
	b := []byte(s)
	for i:= 0; i < len(b)/2; i++ {
		j := len(b)-i-1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}