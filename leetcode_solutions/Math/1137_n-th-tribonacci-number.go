# N-th Tribonacci Number
# Difficulty: Easy
# Language: golang
# Topic: Math
# Tags: Math, Dynamic Programming, Memoization
# Link: https://leetcode.com/problems/n-th-tribonacci-number/

func tribonacci(n int) int {
    memo:=map[int]int{}
    if n==0 {
        return 0
    }
    if n==1 {
        return 1
    }
    if n==2 {
        return 1
    }
    return calc(n-1,memo) + calc(n-2,memo) + calc(n-3,memo)
}

func calc(a int, memo map[int]int) int {
    if a==0 {
        return 0
    }
    if a==1 {
        return 1
    }
    if a==2 {
        return 1
    }
    if _,ok:=memo[a];ok{
        return memo[a]
    }
    memo[a]=calc(a-1,memo) + calc(a-2,memo) + calc(a-3,memo)
    return memo[a]
}