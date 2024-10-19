# Find the Index of the First Occurrence in a String
# Difficulty: Easy
# Language: golang
# Topic: String
# Tags: Two Pointers, String, String Matching
# Link: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

func strStr(haystack string, needle string) int {
    if len(haystack) < len(needle){
        return -1
    }
    hashneedle := hash(needle)
    for i,_ := range haystack{
        if i+len(needle)>len(haystack){
            return -1
        }
        hashhaystack := hash(haystack[i:len(needle)+i])
        if hashhaystack!=hashneedle{
            continue
        }
        if needle == haystack[i:len(needle)+i]{
            return i
        }
    }
    return -1

}

func hash(word string) int{
    ans:=0
    for _,v := range word {
        ans += int(v)
    }
    return ans
}