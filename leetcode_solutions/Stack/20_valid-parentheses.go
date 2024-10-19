# Valid Parentheses
# Difficulty: Easy
# Language: golang
# Topic: Stack
# Tags: String, Stack
# Link: https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
    symbolMap := map[string]string{
        ")":"(",
        "}":"{",
        "]":"[",
    }

    stack := []string{}

    for _, v := range s{
        if openingSymbol, ok := symbolMap[string(v)]; ok{
            // check stack
            if len(stack) ==0 {
                return false
            }
            if stack[len(stack)-1]!= openingSymbol{
                return false
            }
            stack = stack[:len(stack)-1]
            continue
        }
        stack = append(stack, string(v))
    }
    if len(stack) !=0 {
        return false
    }
    return true
}