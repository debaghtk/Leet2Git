# Contains Duplicate
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
    hmap := map[int]int{}
    for _,v := range nums {
        if _,ok:=hmap[v];ok{
            return true
        }
        hmap[v]=1
    }
    return false
}