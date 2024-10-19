# Min Cost Climbing Stairs
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/min-cost-climbing-stairs/

func minCostClimbingStairs(cost []int) int {
    memo := map[int]int{}
    length := len(cost)
    if (length ==0){
        return 0
    }
    if (length == 1) {
        return cost[0]
    }
    if (length == 2) {
        if cost[0]>cost[1]{
            return cost[1]
        }
        return cost[0]
    }
    if (length == 3) {
        if(cost[0]+cost[2] > cost[1]){
            return cost[1]
        }
        return cost[0]+cost[2]
    }
    
    return min(cost[0]+minCost(cost[1:],memo),
               cost[0]+minCost(cost[2:],memo),
               cost[1]+minCost(cost[2:],memo),
               cost[1]+minCost(cost[3:],memo))
}

func minCost(cost []int, memo map[int]int) int{
    length := len(cost)
    if (length==0){
        return 0
    }
    if (length == 1) {
        return cost[0]
    }
    if (length == 2) {
        return cost[0]
    }
    if (length == 3) {
        v,ok:=memo[3]
        if ok{
            return v
        }
        if(cost[2] > cost[1]){
            memo[3]=cost[0]+cost[1]
            return cost[0]+cost[1]
        }
        memo[3]=cost[0]+cost[2]
        return cost[0]+cost[2]
    }
    
    
    if v,ok:=memo[length]; ok {
        return v
    }
    
    memo[length] = minOfTwo(cost[0]+minCost(cost[1:], memo),
               cost[0]+minCost(cost[2:], memo))
    return memo[length]
}

func min(a,b,c,d int) int{
    if (a<=b && a<=c && a<=d){
        return a
    }
    if (b<=a && b<=c && b<=d){
        return b
    }
    if (c<=a && c<=b && c<=d){
        return c
    }
    return d
}

func minOfTwo(a,b int) int{
    if a >b {
        return b
    }
    return a
}