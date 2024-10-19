# Number of 1 Bits
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/number-of-1-bits/

func hammingWeight(num uint32) int {
    result:=0
    for num>0 {
        num = num & (num-1)
        result++
    }
    return result
}