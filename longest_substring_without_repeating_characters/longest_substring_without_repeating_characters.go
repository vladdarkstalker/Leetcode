package main

func lengthOfLongestSubstring(s string) int {
	var res int
	var buffer = make(map[byte]int)
	var l int

	for r := 0; r < len(s); r++ {
		index, exist := buffer[s[r]]
		if index >= l && exist {
			l = buffer[s[r]] + 1
		}
		buffer[s[r]] = r
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}
