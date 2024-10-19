# Number of Islands
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
    rows := len(grid)
    if rows==0{
        return 0
    }
    columns := len(grid[0])
    var answer int

    var dfs func(i,j int)
    dfs = func(i,j int) {
        rightOkay := i+1<rows
        leftOkay := i-1>=0
        upOkay:= j-1>=0
        downOkay := j+1<columns
        grid[i][j]='0'

        if rightOkay && grid[i+1][j]=='1'{
            dfs(i+1,j)
        }
        if leftOkay && grid[i-1][j]=='1'{
            dfs(i-1,j)
        }
        if upOkay && grid[i][j-1]=='1' {
            dfs(i,j-1)
        }
        if downOkay && grid[i][j+1]=='1'{
            dfs(i,j+1)
        }


        return
    }

    for i:=0;i<rows;i++{
        for j:=0;j<columns;j++{
            if grid[i][j] == '1' {
                dfs(i,j)
                answer++
            }
        }
    }

    return answer

}