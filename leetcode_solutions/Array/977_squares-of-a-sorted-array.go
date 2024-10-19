# Squares of a Sorted Array
# Difficulty: Easy
# Language: golang
# Topic: Array
# Tags: Array, Two Pointers, Sorting
# Link: https://leetcode.com/problems/squares-of-a-sorted-array/

func sortedSquares(nums []int) []int {
    nums = splitAndMerge(nums)
    answer:=make([]int, len(nums))
    for i, num := range nums {
        answer[i]=num*num
    }
    return answer
}

func splitAndMerge (nums []int) []int{
    // find 0 element ka index
    for i:=0;i<len(nums);i++{
        if i==0 && nums[i]==0{
            return nums      
        }
        if i==0 && nums[i]>0{
            return nums
        }
        if nums[i]<0{
         nums[i]=nums[i]*(-1)
            continue
        }
        if nums[i]==0||nums[i]>0{
            return merge(nums[:i], nums[i:])
        }
    }
    return reverse(nums)
}

func merge (arr1, arr2 []int)[]int{
    arr1 = reverse(arr1)
    arr3 := make([]int, len(arr1)+len(arr2))
	for i := 0; i < len(arr3); i++ {
		if len(arr2) == 0 {
			copy(arr3[i:], arr1)
			break
		}
		if len(arr1) == 0 {
			copy(arr3[i:], arr2)
			break
		}
		if arr1[0] < arr2[0] {
			arr3[i] = arr1[0]
			arr1 = arr1[1:]
			continue
		}
		arr3[i] = arr2[0]
		arr2 = arr2[1:]
	}
	return arr3
}

func reverse(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}