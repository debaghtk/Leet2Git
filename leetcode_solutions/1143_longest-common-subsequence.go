# Longest Common Subsequence
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/longest-common-subsequence/

func longestCommonSubsequence(text1 string, text2 string) int {
    if text1 == text2{
        return len(text1)
    }
    x := len(text1)
    y := len(text2)
    if x == 0 || y==0 {
        return 0
    }
    if x == 1 {
        if exists(text1, text2) {
            return 1
        }
    }
    if y == 1 {
        if exists(text2, text1) {
            return 1
        }
    }
    memo := make([][]int,x)
    for i:=0;i<x;i++{
        memo[i]=make([]int,y)
    }
    if text1[0]==text2[0]{
        return 1+lcs(text1[1:], text2[1:], memo)
    }
    return max(lcs(text1, text2[1:], memo), lcs(text1[1:], text2, memo))
}

func lcs(text1, text2 string, memo [][]int) int{
        if text1 == text2{
        return len(text1)
    }
    x := len(text1)
    y := len(text2)
    if x == 0 || y==0 {
        return 0
    }
    if x == 1 {
        if exists(text1, text2) {
            return 1
        }
    }
    if y == 1 {
        if exists(text2, text1) {
            return 1
        }
    }
    
    if v := memo[x-1][y-1];v!=0{
        return v
    }
    
    if text1[0]==text2[0]{
        memo[x-1][y-1]= 1+lcs(text1[1:], text2[1:], memo)
        
    } else {
        memo[x-1][y-1]=max(lcs(text1, text2[1:], memo), lcs(text1[1:], text2, memo))
    }
    return memo[x-1][y-1]
}

func exists (char , text string) bool {
    for _,v:= range text{
        if string(v)==char{
            return true
        }
    }
    return false
}

func max (b,c int) int {
    if b >= c {
        return b
    }
    return c
}