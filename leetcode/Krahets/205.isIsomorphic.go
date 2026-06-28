package Krahets

const url205 = "https://leetcode.cn/problems/isomorphic-strings/?envType=study-plan-v2&envId=selected-coding-interview"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var sTt [256]byte
	var tTs [256]byte

	for i := 0; i < len(s); i++ {
		S := s[i]
		T := t[i]
		if sTt[S] != 0 && sTt[S] != T {
			return false
		}
		if tTs[T] != 0 && tTs[T] != S {
			return false
		}

		sTt[S] = T
		tTs[T] = S
	}
	return true

}
