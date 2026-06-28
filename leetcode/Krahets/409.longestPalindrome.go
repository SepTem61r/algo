package Krahets

const url409 = "https://leetcode.cn/problems/longest-palindrome/?envType=study-plan-v2&envId=selected-coding-interview"

func longestPalindrome(s string) int {

	var c [128]int

	n := 0
	on := false
	for i := 0; i < len(s); i++ {
		c[s[i]-'A']++
	}

	for i := 0; i < len(c); i++ {
		n += (c[i] / 2) * 2
		if c[i]%2 != 0 {
			on = true
		}

	}
	if on == true {
		n++
	}
	return n
}
