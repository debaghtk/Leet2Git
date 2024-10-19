# Delete and Earn
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/delete-and-earn/

import "sort"

func deleteAndEarn(nums []int) int {
    earning:=0
    sort.Ints(nums)
    sums := map[int]int{}
    distinct:=[]int{}
    length:=len(nums)
    memo:=map[int]int{}
    if length==1{
        return nums[0]
    }
    if length==2{
        if modIsNotOne(nums[0],nums[1]){
            return nums[0]+nums[1]
        }
        if nums[0]>nums[1]{
            return nums[0]
        }
        return nums[1]
    }
    // calculate sums and distincts
    for _,v:=range nums{
        sums[v]+=v
        if len(distinct)==0 || distinct[len(distinct)-1]!=v{
            distinct = append(distinct, v)
        }
    }
    
    if len(distinct)==1{
        return sums[distinct[0]]
    }
    
    if modIsNotOne(distinct[0],distinct[1]){
        earning = sums[distinct[0]]+compute(distinct[1:], sums, memo)
    } else {
        earning = max(sums[distinct[0]]+compute(distinct[2:], sums, memo), compute(distinct[1:], sums, memo))
    }
    return earning
}

func modIsNotOne(a,b int) bool{
    if a-b==1 || b-a==1{
        return false
    }
    return true
}

func max(a,b int) int{
    if a>b {
        return a
    }
    return b
}

func compute(distinct []int, sums,memo map[int]int) int{
    length:=len(distinct)
    if length==1{
        return sums[distinct[0]]
    }
    if length==2{
        if modIsNotOne(distinct[0],distinct[1]){
            return sums[distinct[0]]+compute(distinct[1:], sums, memo)
        } else {
            return max(sums[distinct[0]], sums[distinct[1]])
                      
        }
    }
    
    if v,ok:= memo[length]; ok{
        return v
    }
                       
    if modIsNotOne(distinct[0],distinct[1]){
    memo[length] = sums[distinct[0]]+compute(distinct[1:], sums, memo)
    } else {
        memo[length] = max(sums[distinct[0]]+compute(distinct[2:], sums, memo), compute(distinct[1:], sums, memo))
    }
    return memo[length]
}