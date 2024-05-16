package lib

import "strconv"

func FormatCurrency(num int) (decoratedNum string) {
	strNum := strconv.Itoa(num)

	revStrNum := ReverseString(strNum)
	var revDecoratedNum string
	for i := 0; i < len(strNum); i++ {
		revDecoratedNum += revStrNum[i : i+1]
		if (i+1)%3 == 0 && i+1 < len(strNum) {
			revDecoratedNum += ","
		}
	}

	for i := len(revDecoratedNum); i > 0; i-- {
		decoratedNum += revDecoratedNum[i-1 : i]
	}

	return
}

func ReverseString(s string) (rs string) {
	for i := len(s); i > 0; i-- {
		rs += s[i-1 : i]
	}
	return
}
