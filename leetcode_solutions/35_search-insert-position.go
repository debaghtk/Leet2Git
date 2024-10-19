# Search Insert Position
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
    if (len(nums) == 1){
        if (nums[0]<target){
            return 1;
        }
        return 0;
    }
    
    k := len(nums)
    split := int(k/2)
    return searchInsert(nums[:split], target)+searchInsert(nums[split:], target)
}