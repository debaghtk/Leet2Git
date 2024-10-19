# Sort an Array
# Difficulty: Medium
# Language: golang
# Topic: Heap (Priority Queue)
# Tags: Array, Divide and Conquer, Sorting, Heap (Priority Queue), Merge Sort, Bucket Sort, Radix Sort, Counting Sort
# Link: https://leetcode.com/problems/sort-an-array/

func sortArray(nums []int) []int {
    if len(nums)==0 || len(nums)==1{
        return nums
    }
    if len(nums)==2{
        if nums[0]>nums[1]{
            return []int{nums[1],nums[0]}
        }
        return nums
    }
    
    half := int(len(nums)/2)
    left:= nums[:half]
    right:= nums[half:]
    sortedLeft:= sortArray(left)
    sortedRight:=sortArray(right)
    
    return merge(sortedLeft, sortedRight)
}

func merge(left, right []int) []int {
    result := []int{}
    for i,v := range left {
        if len(right)==0{
            result = append(result, left[i:]...)
            return result
        }
        if v <= right[0]{
            result = append(result, v)
            continue
        }
        for (v > right[0]) {
            result = append(result, right[0])
            if len(right)==1{
                right = []int{}
                break
            }
            right = right[1:]
        }
        result = append(result, v)
    }
    result = append(result, right...)
    return result
}