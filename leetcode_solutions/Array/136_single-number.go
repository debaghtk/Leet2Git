# Single Number
# Difficulty: Easy
# Language: golang
# Topic: Array
# Tags: Array, Bit Manipulation
# Link: https://leetcode.com/problems/single-number/

func singleNumber(nums []int) int {
    result := nums[0]
    if len(nums)==1{
        return result
    }
    for _,v := range nums[1:] {
        result = result ^ v
    }
    return result
}