# Word Break
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool , len(s)+1)
    indexes := []int{0}

    for i:=1; i<len(s)+1;i++ {
        substrings := calcSubstrings(s,i,indexes)
        for _,substring := range substrings {
            if wordExists(substring, wordDict){
                dp[i] = true
                indexes = append(indexes, i)
                break
            }
        }
    }
    return dp[len(s)]
}

func calcSubstrings(s string, i int, indexes []int) []string{
    answer:= []string{}
    for _, index := range indexes {
        answer = append(answer, s[index:i])
    }
    return answer
}

func wordExists(s string, wordDict []string) bool {
    for _,word:= range wordDict{
        if word == s{
            return true
        }
    }
    return false
}