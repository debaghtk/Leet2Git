# Majority Element
# Difficulty: Easy
# Language: golang
# Topic: Hash Table
# Tags: Array, Hash Table, Divide and Conquer, Sorting, Counting
# Link: https://leetcode.com/problems/majority-element/

func majorityElement(nums []int) int {
    counter := map[int]int{}
    size := len(nums)/2
    for _,v := range nums {
        counter[v]++
        if counter[v] > size{
            return v
        }
    }
    return -1
}