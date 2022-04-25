package main

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// 暴力写法
func romanToInt1(s string) int {
	var result int
	var isSubtraction bool

	a := []rune(s)
	for i := 0; i < len(a); i++ {
		roman := string(a[i])
		if i+1 < len(a) {
			roman1 := string(a[i+1])
			if roman == "I" && (roman1 == "V" || roman1 == "X") {
				isSubtraction = true
			}
			if roman == "X" && (roman1 == "L" || roman1 == "C") {
				isSubtraction = true
			}
			if roman == "C" && (roman1 == "D" || roman1 == "M") {
				isSubtraction = true
			}
			if isSubtraction {
				result = result + (romanMap[roman1] - romanMap[roman])
				i = i + 1
				isSubtraction = false
				continue
			}
		}
		result = result + romanMap[roman]
	}
	return result
}

// https://leetcode.com/problems/roman-to-integer/discuss/6529/My-solution-for-this-question-but-I-don't-know-is-there-any-easier-way/141371
func romanToInt2(s string) int {
	var result int

	a := []rune(s)
	i, j := 0, 1
	for ; j < len(a); i, j = i+1, j+1 {
		if romanMap[string(a[i])] >= romanMap[string(a[j])] {
			result += romanMap[string(a[i])]
		} else {
			result -= romanMap[string(a[i])]
		}
	}
	result += romanMap[string(a[i])]
	return result
}
