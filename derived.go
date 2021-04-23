package main

// mask certain raw string
func mask(raw string, from int, mask rune) string {
	rs := []rune(raw)
	for i := from; i < len(rs); i++ {
		rs[i] = mask
	}
	return string(rs)
}
