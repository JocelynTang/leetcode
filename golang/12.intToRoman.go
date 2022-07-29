package golang

func intToRoman(num int) string {
	var s string
	/*
		r := make(map[int]string)
		r[1000] = "M"
		r[900] = "CM"
		r[500] = "D"
		r[400] = "CD"
		r[100] = "C"
		r[90] = "XC"
		r[50] = "L"
		r[40] = "XL"
		r[10] = "X"
		r[9] = "IX"
		r[5] = "V"
		r[4] = "IV"
		r[1] = "I"
	*/
	for num != 0 {
		if num >= 1000 {
			s += "M"
			num -= 1000
		}
		if num < 1000 && num >= 900 {
			s += "CM"
			num -= 900
		}
		if num < 900 && num >= 500 {
			s += "D"
			num -= 500
		}
		if num < 500 && num >= 400 {
			s += "CD"
			num -= 400
		}
		if num < 400 && num >= 100 {
			s += "C"
			num -= 100
		}
		if num < 100 && num >= 90 {
			s += "XC"
			num -= 90
		}
		if num < 90 && num >= 50 {
			s += "L"
			num -= 50
		}
		if num < 50 && num >= 40 {
			s += "XL"
			num -= 40
		}
		if num < 40 && num >= 10 {
			s += "X"
			num -= 10
		}
		if num < 10 && num >= 9 {
			s += "IX"
			num -= 9
		}
		if num < 9 && num >= 5 {
			s += "V"
			num -= 5
		}
		if num < 5 && num >= 4 {
			s += "IV"
			num -= 4
		}
		if num < 4 && num >= 1 {
			s += "I"
			num -= 1
		}
	}

	return s
}
