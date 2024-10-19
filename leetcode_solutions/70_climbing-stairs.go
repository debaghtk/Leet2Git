# Climbing Stairs
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/climbing-stairs/

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2{
        return 2
    }
    memory:= map[int]*int{}
    return climb(n-1, memory) + climb(n-2, memory)
}

func climb(n int, memory map[int]*int) int {
    if n == 1 {
        return 1
    }
    if n == 2{
        return 2
    }
    v,ok := memory[n]
    if ok{
        return *v
    }

    a := climb(n-1, memory) + climb(n-2, memory)
    memory[n]=&a
    return *memory[n]
}