# Surrounded Regions
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/surrounded-regions/

func solve(grid [][]byte)  {
    rows := len(grid)
    if rows <=0 {
        return
    }
    columns := len(grid[0])

    var dfs func(i,j int)
    dfs = func(i,j int) {

        if i<0 || j<0 || i>=rows || j >= columns || grid[i][j]== 'X' || grid[i][j]== 'F'{
            return
        }

        grid[i][j]='F'

        dfs(i+1,j)
        dfs(i,j+1)
        dfs(i,j-1)
        dfs(i-1,j)
    }

    for i:=0;i<rows;i++{
        if grid[i][0] == 'O'{
            dfs(i,0)
        }
        if grid[i][columns-1] == 'O'{
            dfs(i,columns-1)
        }
    }

    for j:=0;j<columns;j++{
        if grid[0][j] == 'O'{
            dfs(0,j)
        }
        if grid[rows-1][j] == 'O'{
            dfs(rows-1,j)
        }
    }

    for i:=0;i<rows;i++{
        for j:=0;j<columns;j++{
            if grid[i][j] == 'F'{
                grid[i][j]='O'
                continue
            }
            if grid[i][j] == 'O'{
                grid[i][j]='X'
            }
        }
    }
}