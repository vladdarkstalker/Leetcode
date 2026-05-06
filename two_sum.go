# Intuition
The simplest idea is to check every pair of numbers and find the pair whose sum equals the target.

# Approach
Use two nested loops.  
For each number, compare it with every number after it.  
If their sum equals `target`, return their indexes.


# Complexity
- Time complexity: $$O(n^2)$$
- Space complexity: $$O(1)$$

# Code
```golang []
package main

func twoSum(nums []int, target int) []int {
    result := make([]int, 2)
    len_nums := len(nums)
    for i := 0; i < len_nums; i++ {
        for j := i + 1; j < len_nums; j++ {
            if nums[i] + nums[j] == target {
                result[0] = i
                result[1] = j
                break
            }
        }
    }
    return result
}

```
