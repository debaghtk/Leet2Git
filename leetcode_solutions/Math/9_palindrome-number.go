# Palindrome Number
# Difficulty: Easy
# Language: golang
# Topic: Math
# Tags: Math
# Link: https://leetcode.com/problems/palindrome-number/

import "strconv"

func isPalindrome(x int) bool {
    subject := strconv.Itoa(x)
    length := len(strconv.Itoa(x))
    iterations := length/2
    if length%2 == 1 {
        iterations += 1
    }
    for i:=0 ; i< iterations ; i++ {
        if subject[i] == subject[length - (i+1)] {
            continue
        }
        return false
    }
    return true
}