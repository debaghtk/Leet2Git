# Remove Duplicates from Sorted Array II
# Difficulty: Medium
# Language: golang
# Topic: Array
# Tags: Array, Two Pointers
# Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

func removeDuplicates(nums []int) int {
        counter := map[int]int{}
    idx:=0
    for i,v := range nums {
        if counter[v]<2{
            counter[v]++
            nums[idx]=nums[i]
            idx++
        }
    }
    return idx
}