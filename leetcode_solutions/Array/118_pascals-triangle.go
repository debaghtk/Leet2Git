# Pascal's Triangle
# Difficulty: Easy
# Language: golang
# Topic: Array
# Tags: Array, Dynamic Programming
# Link: https://leetcode.com/problems/pascals-triangle/

func generate(numRows int) [][]int {
    result:= [][]int{}
    for i:=1;i<=numRows;i++{
        if i==1 {
            result=append(result, []int{1})
            continue
        }
        if i==2 {
            result=append(result, []int{1,1})
            continue
        }
        arr:=make([]int, i)
        for k,_:=range arr {
            if k==0{
                arr[0]=1
                continue
            }
            if k==i-1{
                arr[i-1]=1
                continue
            }
            arr[k]=result[i-2][k]+result[i-2][k-1]
        }
        result=append(result, arr)
    }
    return result
}