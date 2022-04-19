package main

import (
	"strconv"
)

//暴力写法
func isPalindrome1(x int) bool {
	if x == 0 {
		return true
	}
	y := strconv.Itoa(x)
	a := []rune(y)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a) == y
}

// https://www.code-recipe.com/post/palindrome-number 原文思路
func isPalindrome2(x int) bool {
	var reversedNum int
	tmp := x
	for tmp > 0 {
		reversedNum = reversedNum*10 + tmp%10
		tmp = tmp / 10
	}
	return x == reversedNum
}
