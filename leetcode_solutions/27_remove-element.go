# Remove Element
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/remove-element/

func removeElement(nums []int, val int) int {
    answer := len(nums)
    for i:=0; i< answer;{
        if nums[i]== val{
            answer--
            copy(nums[i:], nums[i+1:])
            continue
        }
        i++
    }
    return answer 
}