# Valid Palindrome
# Difficulty: Easy
# Language: golang
# Topic: String
# Tags: Two Pointers, String
# Link: https://leetcode.com/problems/valid-palindrome/

import "regexp"
import "strings"
func isPalindrome(s string) bool {
    // remove anything which is not english letter
    //s=strings.ReplaceAll(s, " ", "")
    st:=s
    s=""
    regex:= regexp.MustCompile(`[a-zA-Z0-9]*`)
    for _,v := range regex.FindAllString(st, -1){
        s+=strings.ToLower(v)
    }
    // return s=="man,plan,canal"
    for true {
        if len(s)<2{
            return true
        }
        if s[0]==s[len(s)-1]{
            s = s[1:len(s)-1]
            continue
        }
        if s[0]!=s[len(s)-1]{
            break
        }
    }
    return false
}