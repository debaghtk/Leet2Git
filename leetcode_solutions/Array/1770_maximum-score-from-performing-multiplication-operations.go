# Maximum Score from Performing Multiplication Operations
# Difficulty: Hard
# Language: golang
# Topic: Array
# Tags: Array, Dynamic Programming
# Link: https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/

func maximumScore(nums []int, multipliers []int) int {
    memo:=map[int]map[int]int{}
    n:=len(nums)
    m:=len(multipliers)
    if m==1{
        return max(multipliers[0]*nums[n-1],multipliers[0]*nums[0]) 
    }
    return max(multipliers[0]*nums[n-1]+maxScore(nums[:n-1],multipliers[1:], memo, 0),
               multipliers[0]*nums[0]+maxScore(nums[1:],multipliers[1:], memo, 1))
}

func max(a,b int) int{
    if a > b {
        return a
    }
    return b
}

func maxScore(nums []int, multipliers []int, memo map[int]map[int]int, left int) int {
    n:=len(nums)
    m:=len(multipliers)
    if m==1{
        return max(multipliers[0]*nums[n-1],multipliers[0]*nums[0]) 
    }
    
    if v,ok:= memo[m][left]; ok{
        return v
    }
    leftMap:=map[int]int{}
    leftMap[left] = max(multipliers[0]*nums[n-1]+maxScore(nums[:n-1],multipliers[1:], memo, left),
               multipliers[0]*nums[0]+maxScore(nums[1:],multipliers[1:], memo, left+1))
    memo[m] = leftMap
    return memo[m][left]
}