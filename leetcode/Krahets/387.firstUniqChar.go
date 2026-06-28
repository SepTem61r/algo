package Krahets

const url387 = "https://leetcode.cn/problems/first-unique-character-in-a-string/?envType=study-plan-v2&envId=selected-coding-interview"

func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	var c [26]int

	for i := 0; i < len(s); i++ {
		c[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if c[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
