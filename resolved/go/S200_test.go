package lcz

import (
	"encoding/json"
	"testing"
)

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。 
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。 
//
// 此外，你可以假设该网格的四条边均被水包围。 
//
// 
//
// 示例 1： 
//
// 
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
// 
//
// 示例 2： 
//
// 
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
// 
//
// 
//
// 提示： 
//
// 
// m == grid.length 
// n == grid[i].length 
// 1 <= m, n <= 300 
// grid[i][j] 的值为 '0' 或 '1' 
// 
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 
// 👍 1246 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	mapped := make([][]bool, len(grid))
	for i := 0; i < len(mapped); i++ {
		mapped[i] = make([]bool, len(grid[0]))
	}
	count := 0
	for i, bytes := range grid {
		for j, b := range bytes {
			if !mapped[i][j] && b == '1' {
				count++
				findRange(&mapped, &grid, i, j)
			}
		}
	}
	return count
}

func findRange(mapped *[][]bool, grid *[][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*mapped)[i][j] || (*grid)[i][j] != '1' {
		return
	}
	(*mapped)[i][j] = true
	findRange(mapped, grid, i-1, j)
	findRange(mapped, grid, i+1, j)
	findRange(mapped, grid, i, j-1)
	findRange(mapped, grid, i, j+1)
}

//leetcode submit region end(Prohibit modification and deletion)

func Test_S200(t *testing.T) {
	tests := []struct {
		Grid   [][]byte
		Result int
	}{
		{
			Grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			Result: 1,
		},
		{
			Grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			Result: 3,
		},
	}
	for _, test := range tests {
		t.Run("S200_test", func(t *testing.T) {
			result := numIslands(test.Grid)
			x := make([][]string, len(test.Grid))
			for i, bytes := range test.Grid {
				t := make([]string, len(bytes))
				for j, b := range bytes {
					t[j] = string(b)
				}
				x[i] = t
			}
			j, _ := json.Marshal(x)
			t.Logf("grid: %s, expected expectedResult: %d, run expectedResult: %d", string(j), test.Result, result)
			if result != test.Result {
				t.Fatalf("failed")
			}
		})
	}
}
