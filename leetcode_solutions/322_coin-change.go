# Coin Change
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/coin-change/

func coinChange(coins []int, amount int) int {
    results := make([]int, amount+1) // Initialize results array with amount+1 size
    for i := 1; i <= amount; i++ {
        results = coinImpl(coins, i, results)
    }
    return results[amount]
}

func coinImpl(coins []int, amount int, results []int) []int {
    answer := -1 // Initialize with -1 to indicate no solution

    if amount == 0 {
        return append(results, 0) // Base case: 0 amount requires 0 coins
    }

    for _, coin := range coins {
        if amount-coin < 0 {
            continue // Skip if coin is larger than the current amount
        }

        if results[amount-coin] != -1 { // Check if there is a valid solution for the subproblem
            current := results[amount-coin] + 1
            if answer == -1 || current < answer {
                answer = current // Update answer with the minimum number of coins
            }
        }
    }

    return append(results, answer)
}
