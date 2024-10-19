# Height Checker
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/height-checker/

func heightChecker(heights []int) int {
    expected:=make([]int,len(heights))
    copy(expected,heights)
    sort.Ints(expected)
    answer:=0
    for i:=0;i<len(expected);i++{
        if expected[i] != heights[i]{
            answer++
        }
    }
    return answer
}
