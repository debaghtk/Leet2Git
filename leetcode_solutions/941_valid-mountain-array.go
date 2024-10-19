# Valid Mountain Array
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/valid-mountain-array/

func validMountainArray(arr []int) bool {
    if len(arr)<3{
        return false
    }
    tip:=0
    for i:=1; i<len(arr)-1; i++{
        if arr[i] > arr[i-1] && arr[i]>arr[i+1]{
            tip=i
        }
    }
    if tip==0{
        return false
    }
    for i:=1; i<tip; i++{
        if arr[i]<=arr[i-1] {                                          
            return false
        }
    }
    for i:=tip+1; i<len(arr); i++{
        if arr[i]>=arr[i-1] {                                          
            return false
        }
    }
    
    return true
}