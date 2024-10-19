# Third Maximum Number
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/third-maximum-number/

func thirdMax(nums []int) int {
    limit:=int(math.Pow(2,31))
    max:=-1*limit -1
    max2:=-1*limit -1
    max3:=-1*limit -1
    

    
    for _,val := range nums{
        if val > max {
            max3=max2
            max2=max
            max = val
        }
        if val < max && val > max2{
            max3=max2
            max2=val
        }
        if val < max2 && val > max3{
            max3=val
        }
    }
    
    if max3!=-1*limit-1{
        return max3
    }
    return max
}