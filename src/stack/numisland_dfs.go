package stack

func NumIslandsDFS(grid [][]byte) int {
    num := 0
    for x := 0; x < len(grid); x++ {
        for y := 0; y < len(grid[x]); y++ {
            if grid[x][y] == '1' {
                num++
                dfs(grid, x, y)
            }
        }
    }
    return num
}

func dfs(grid [][]byte, x int, y int) {
    if x >= len(grid) || x < 0 || y >= len(grid[x]) || y < 0 {
        return
    }
    if grid[x][y] == '0' {
        return
    }
    grid[x][y] = '0'
    dfs(grid, x - 1, y)
    dfs(grid, x + 1, y)
    dfs(grid, x, y + 1)
    dfs(grid, x, y - 1)
}
