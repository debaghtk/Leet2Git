# Evaluate Division
# Difficulty: Medium
# Language: golang
# Topic: Graph
# Tags: Array, String, Depth-First Search, Breadth-First Search, Union Find, Graph, Shortest Path
# Link: https://leetcode.com/problems/evaluate-division/

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    type division struct {
        output string
        divisor float64
    }
    graph := map[string][]division{}
    exists := map[string]bool{}
    for i,v := range equations{
        numerator := v[0]
        denominator := v[1]
        graph[numerator] = append(graph[numerator], division{output: denominator, divisor: values[i]})
        graph[denominator] = append(graph[denominator], division{output: numerator, divisor: float64(1)/values[i]})
        exists[numerator]=true
        exists[denominator]=true
    }

    var findDivisor func(num, den string, ans float64, visited map[string]bool) float64
    findDivisor = func(num, den string, ans float64, visited map[string]bool) float64{
        if num == den {
            return ans
        }
        visited[num]=true
        for _,v:= range graph[num]{
            if visited[v.output] {
                continue
            }
            
            if v.output == den{
                return v.divisor*ans
            }
           next := findDivisor(v.output, den, ans*v.divisor, visited)
            if  next != float64(-1){
                return next
            }
        }
        return float64(-1)
    }

    answer :=[]float64{}
    for _, v:= range queries{
        num := v[0]
        den := v[1]
        if !exists[num] || !exists[den]{
            answer = append(answer, float64(-1))
            continue
        }
        // if num == den {
        //     answer = append(answer, float64(1))
        //     continue
        // }

        answer = append(answer, findDivisor(num, den, float64(1), map[string]bool{}))

    }
    return answer
}