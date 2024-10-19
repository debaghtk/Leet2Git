# Is Subsequence
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/is-subsequence/

func isSubsequence(s string, t string) bool {
    if len(s) == 0{
        return true
    }
    if len(s) > len(t) {
        return false
    }

    if len(s) == len(t){
        if s==t{
            return true
        }
        return false
    }

    if s[0]==t[0]{
        return isSubsequence(s[1:], t[1:])
    }
    return isSubsequence(s, t[1:])
}