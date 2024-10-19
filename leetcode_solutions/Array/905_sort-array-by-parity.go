# Sort Array By Parity
# Difficulty: Easy
# Language: golang
# Topic: Array
# Tags: Array, Two Pointers, Sorting
# Link: https://leetcode.com/problems/sort-array-by-parity/

func sortArrayByParity(nums []int) []int {
    answer:=make([]int, len(nums))

    for _, val := range nums{
        if val%2 != 0 {
            copy(answer[1:],answer)
            answer[0]=val
            continue
        }
    }
    for _, val := range nums{
        if val%2 == 0 {
            copy(answer[1:],answer)
            answer[0]=val
            
        }
    }
    
    return answer
}