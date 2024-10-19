# Sqrt(x)
# Difficulty: Easy
# Language: golang
# Topic: Math
# Tags: Math, Binary Search
# Link: https://leetcode.com/problems/sqrtx/


func mySqrt(x int) int {
    if x == 0 {
        return x
    }
    if x < 4 {
        return 1
    }
    l:= mySqrt(x>>2)<<1
    r:=l+1
    if r*r>x{
        return l
    }
    return r
}