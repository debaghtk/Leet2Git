# Duplicate Zeros
# Difficulty: Easy
# Language: golang
# Topic: Array
# Tags: Array, Two Pointers
# Link: https://leetcode.com/problems/duplicate-zeros/

func duplicateZeros(arr []int)  {
    for i:=0;i<len(arr)-1;i++{
        if arr[i]==0{
            //arr = append(arr,0)
            copy(arr[i+2:], arr[i+1:])
            arr[i+1]=0
            i++
        }
    }
}