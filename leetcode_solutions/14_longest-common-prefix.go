# Longest Common Prefix
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
    length:=len(strs)
    if length == 0{
        return ""
    }
    if length == 1{
        return strs[0]
    }
    if length == 2{
        var i int
        for i=0;true;i++{
            if i>=len(strs[0]) || i>=len(strs[1]) {
                break
            }
            if strs[0][i] != strs[1][i] {
                break
            }
        }
        return strs[0][:i]
    }
    half:=int(length/2)
    return longestCommonPrefix([]string{longestCommonPrefix(strs[:half]), longestCommonPrefix(strs[half:])})
}