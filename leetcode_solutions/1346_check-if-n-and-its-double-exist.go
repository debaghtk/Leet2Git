# Check If N and Its Double Exist
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/check-if-n-and-its-double-exist/

func checkIfExist(arr []int) bool {
    for i:=0; i<len(arr);i++{
        for j:=0;j<len(arr);j++{
           if i==j{
                continue
            }
            if arr[i]==2*arr[j] || arr[j]==arr[i]*2{
                return true
            }   
        }     
    }   
    return false
}