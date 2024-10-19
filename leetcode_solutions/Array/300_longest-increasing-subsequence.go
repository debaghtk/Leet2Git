# Longest Increasing Subsequence
# Difficulty: Medium
# Language: golang
# Topic: Array
# Tags: Array, Binary Search, Dynamic Programming
# Link: https://leetcode.com/problems/longest-increasing-subsequence/

import "slices"
//import "math"
func lengthOfLIS(nums []int) int {
    results := make([]int, len(nums))
    for i, num := range nums {
        if i == 0{
            results[0]=1
            continue
        }
        for j:=i-1;j>=0;j--{
            if num > nums[j] {
                results[i] = max(results[i] , results[j] + 1)
                //break
            }
        }
        if results[i]==0 {
            //results[i]=results[i-1]
            results[i]=1
        }
        
    }
    return slices.Max(results)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}