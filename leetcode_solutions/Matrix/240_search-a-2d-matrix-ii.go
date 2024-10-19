# Search a 2D Matrix II
# Difficulty: Medium
# Language: golang
# Topic: Matrix
# Tags: Array, Binary Search, Divide and Conquer, Matrix
# Link: https://leetcode.com/problems/search-a-2d-matrix-ii/

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix)==1 && len(matrix[0])==1{
        return target == matrix[0][0]
    }
    if len(matrix)==1{
        return searchArr(matrix[0], target)
    }
    result := false
    for i,_ := range matrix {
        if target == matrix[i][0] {
            return true
        }
        if target > matrix[i][0] {
            result = result || searchArr(matrix[i],target)
        }
        if target < matrix[i][0] {
            break
        }
    }
    return result
}

func searchArr(arr []int, target int) bool{
    if len(arr)==1 {
        return target == arr[0]
    }
    k := int(len(arr)/2)
    if target > arr[k]{
        return searchArr(arr[k:],target)
    }
    if target < arr[k]{
        return searchArr(arr[:k],target)
    }
    return true
}