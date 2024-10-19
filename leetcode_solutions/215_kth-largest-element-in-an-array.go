# Kth Largest Element in an Array
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/kth-largest-element-in-an-array/

func findKthLargest(nums []int, k int) int {
    pivot := len(nums)-1
   // pivot:=nums[len(nums)-1]
    //left:=[]int{}
    //right:=[]int{}
    i:=0
    for true{
        if nums[i]>=nums[pivot]{
           // left=append(left, nums[i])
            i=i+1
            if i>=pivot {
                break
            }
            continue
        }
        pivot=pivot-1
        nums=append(append(nums[:i], nums[i+1:]...), nums[i])
        if i>=pivot {
            break
        }
    }
    if k<pivot+1{
        return findKthLargest(nums[:pivot],k)
    }
    if k==pivot+1{
        return nums[pivot]
    }
    if k>pivot+1{
        return findKthLargest(nums[pivot+1:], k-(pivot+1) )    
    }
    return -1 // should never happen
}