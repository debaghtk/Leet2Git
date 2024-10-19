# Two Sum
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
   // answer:=[]int{}
    if len(nums)==2{
        return []int{0,1}
    }
    hashmap:=mapping(nums)
    for i,v:=range nums{
        complement:=target-v
        pos,ok:= hashmap[complement]
        if ok && pos[0]!=i {
            return []int{i,pos[0]}
        }
        if len(pos)==2{
            if pos[1]!=i{
                return []int{i,pos[1]}
            }
        }
    }
    return []int{}
}

func mapping(nums []int) map[int][]int{
    result:=map[int][]int{}
    for i,v:=range nums{
        result[v]=append(result[v],i)
    }
    return result
}