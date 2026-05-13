# Intuition
Use a sliding window to keep a substring without repeating characters.

# Approach
Move the right pointer through the string and store the last seen index of each character in a map.  
If the current character was already seen inside the current window, move the left pointer after its previous position.  
Update the maximum window length on each step.

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(k)$$  
where `k` is the number of unique characters in the string.

# Code
```golang []
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
```
