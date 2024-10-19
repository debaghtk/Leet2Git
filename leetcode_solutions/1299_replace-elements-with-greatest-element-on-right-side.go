# Replace Elements with Greatest Element on Right Side
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/

func replaceElements(arr []int) []int {
    if len(arr)==1{
        return []int{-1}
    }

    for i:=0;i<len(arr)-1;i++{
        arr[i]=findGreatest(arr[i+1:])
    }
    arr[len(arr)-1]=-1
    return arr
}

func findGreatest(arr []int) int{
    largest:=0
    for i:=0;i<len(arr);i++{
        if arr[i]>largest{
            largest = arr[i]
        }
    }
    return largest
}