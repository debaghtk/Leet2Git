# Power of Three
# Difficulty: Easy
# Language: golang
# Topic: Math
# Tags: Math, Recursion
# Link: https://leetcode.com/problems/power-of-three/

import "math"
func isPowerOfThree(n int) bool {
    if n<= 0{
        return false
    }
    log3 := math.Log(3.0)
    logn := math.Log(float64(n))
    lognbase3 := math.Round(logn/log3)
    if math.Pow(3,lognbase3)==float64(n){
        return true
    }
    return false
}