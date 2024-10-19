# Remove Duplicates from Sorted Array
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
    answer:=len(nums)
    for i:=1;i<answer;{
        if nums[i]==nums[i-1]{
            answer--
            copy(nums[i:], nums[i+1:])
            continue
        }
        i++
    }
    
    return answer
}