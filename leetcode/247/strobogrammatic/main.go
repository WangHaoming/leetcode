package strobogrammatic

func strobogrammatic(n, m int) []string {
	var res, resFin []string
	if n == 0 {
		res = []string{""}
		return res
	}
	if n == 1 {
		res = []string{"0", "1", "8"}
		return res
	}
	res = strobogrammatic(n-2, m)
	for _, value := range res {
		if m != n {
			resFin = append(resFin, "0"+value+"0")
		}
		resFin = append(resFin, "1"+value+"1")
		resFin = append(resFin, "6"+value+"9")
		resFin = append(resFin, "8"+value+"8")
		resFin = append(resFin, "9"+value+"6")

	}
	return resFin
}

func Test() {
	n := 5
	res := strobogrammatic(n, n)
	for _, value := range res {
		println(value)
	}
}
