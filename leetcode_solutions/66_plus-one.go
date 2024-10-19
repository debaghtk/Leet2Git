# Plus One
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/plus-one/

func plusOne(digits []int) []int {
    if len(digits)==1 && digits[0]==9{
        return []int{1,0}
    }
    if len(digits)==1{
        return []int{digits[0]+1}
    }
    last:=len(digits)-1
    if digits[last]==9{
        return append(plusOne(digits[:last]), 0)
    }
    digits[last]+=1
    return digits
}