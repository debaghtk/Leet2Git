# Move Zeroes
# Difficulty: Easy
# Language: golang
# Topic: Array
# Tags: Array, Two Pointers
# Link: https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int)  {
    count:=0
    for _,val:=range nums{
        if val==0{
            count++
        }
    }
    for i:=0; count>0;{
        if nums[i]==0{
            count--
            if i<len(nums)-1{
                copy(nums[i:], nums[i+1:])
                nums[len(nums)-1]=0
                if nums[i]==0{
                    continue
                }
            }
        }
        i++
    }
}