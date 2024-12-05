package dfs

/**
200. 岛屿数量
level: medium
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。
*/

/**
这是一个典型的深度优先搜索（DFS）或广度优先搜索（BFS）问题。我们可以通过遍历整个网格，每遇到一个陆地（'1'）就进行一次深度优先搜索或广度优先搜索，
将相邻的陆地都标记为已访问，然后岛屿数量加一。这样，我们就可以找出所有的岛屿。

以下是具体的步骤：

初始化岛屿数量为0。
遍历整个二维网格，对于每一个单元格，如果它是陆地（'1'），则进行深度优先搜索或广度优先搜索，并将岛屿数量加一。
在深度优先搜索或广度优先搜索中，我们首先将当前单元格标记为已访问（例如将其值改为'0'），然后分别向上、下、左、右四个方向进行搜索，
如果相邻的单元格是陆地，则继续进行深度优先搜索或广度优先搜索。
最后返回岛屿数量。
*/

// 深度优先遍历
// 习惯的，我们将 len(grip)对应的索引成为 i，而将 grid[0]对应的索引成为 j
// i 的方向是从上到下，j 的方向是从左到右
func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    res := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            // 找到了一个1，可能是一片1，把与它相邻的1都重置成0
            if grid[i][j] == '1' {
                numIslandsDFS(grid, i, j)
                res++
            }
        }
    }
    return res
}

func numIslandsDFS(grid [][]byte, i, j int) {
    // 判断
    // 1. i和 j 的范围，不能大于多少，不能小于多少
    // 2. 值是不是0
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
        return
    }
    grid[i][j] = '0'
    // 遍历grid[i][j]的上下左右, 将相邻的1都重置成0
    numIslandsDFS(grid, i, j+1)
    numIslandsDFS(grid, i, j-1)
    numIslandsDFS(grid, i+1, j)
    numIslandsDFS(grid, i-1, j)
}
