package golang

func romanToInt(s string) int {
	r := make(map[string]int)
	r["I"] = 1
	r["V"] = 5
	r["X"] = 10
	r["L"] = 50
	r["C"] = 100
	r["D"] = 500
	r["M"] = 1000
	s1 := []rune(s)
	ans := r[string(s1[len(s1)-1])]
	for i := len(s1) - 1; i > 0; i-- {
		if r[string(s1[i])] <= r[string(s1[i-1])] {
			ans += r[string(s1[i-1])]
		} else {
			ans -= r[string(s1[i-1])]
		}
	}
	return ans
}
