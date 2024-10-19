# House Robber
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/house-robber/

var gmemo map[int]int = map[int]int{}
func rob(nums []int) int {
    gmemo = make(map[int]int)
    if len(nums)==0{
        return 0
    }
    if len(nums)==1{
        return nums[0]
    }
    if len(nums)==2{
        if nums[0]>nums[1]{
            return nums[0]
        }
        return nums[1]
    }
    if len(nums)==3{
        if nums[0]+nums[2]>nums[1]{
            return nums[0]+nums[2]
        }
        return nums[1]
    }
    
    v,ok:=gmemo[len(nums)]
    if ok {
        return v
    }
    gmemo[len(nums)] = max( nums[0]+ srob(nums[2:]) , nums[1]+ srob(nums[3:]) )
    return gmemo[len(nums)]
}

func srob(nums []int) int {
    if len(nums)==0{
        return 0
    }
    if len(nums)==1{
        return nums[0]
    }
    if len(nums)==2{
        if nums[0]>nums[1]{
            return nums[0]
        }
        return nums[1]
    }
    if len(nums)==3{
        if nums[0]+nums[2]>nums[1]{
            return nums[0]+nums[2]
        }
        return nums[1]
    }
    
    v,ok:=gmemo[len(nums)]
    if ok {
        return v
    }
    gmemo[len(nums)] = max( nums[0]+ srob(nums[2:]) , nums[1]+ srob(nums[3:]) )
    return gmemo[len(nums)]
}

func max(a,b int) int{
    if a > b {
        return a
    }
    return b
}