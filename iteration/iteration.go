package iteration

func Repeat(s string, x int) string {
	var out string
	for i := 0; i < x; i++ {
		out += s
	}
	return out
}
