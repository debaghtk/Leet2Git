# Fizz Buzz
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/fizz-buzz/

import "fmt"
func fizzBuzz(n int) []string {
    answer:=[]string{}
    for i:=1;i<=n;i++{
        result:=""
        if i%3==0{
            result = result + "Fizz"
        }
        if i%5==0{
            result = result + "Buzz"
        }
        if result == ""{
            result = fmt.Sprintf("%d",i)
        }
        answer = append(answer, result)
    }
    return answer
}