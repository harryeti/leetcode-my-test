package lcz

import (
	"encoding/json"
	"testing"
)

//有 n 个网络节点，标记为 1 到 n。
//
// 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， w
//i 是一个信号从源节点传递到目标节点的时间。
//
// 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
//输出：2
//
//
// 示例 2：
//
//
//输入：times = [[1,2,1]], n = 2, k = 1
//输出：1
//
//
// 示例 3：
//
//
//输入：times = [[1,2,1]], n = 2, k = 2
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= k <= n <= 100
// 1 <= times.length <= 6000
// times[i].length == 3
// 1 <= ui, vi <= n
// ui != vi
// 0 <= wi <= 100
// 所有 (ui, vi) 对都 互不相同（即，不含重复边）
//
// Related Topics 深度优先搜索 广度优先搜索 图 最短路 堆（优先队列）
// 👍 320 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func networkDelayTime(times [][]int, n int, k int) int {
	if len(times) < n-1 {
		return -1
	}
	xMap := map[int]map[int]int{}
	var xTime map[int]int
	for _, time := range times {
		x, e := xMap[time[0]]
		if !e {
			xMap[time[0]] = map[int]int{time[1]: time[2]}
		} else {
			x[time[1]] = time[2]
		}
	}
	xTime, e := xMap[k]
	if !e {
		return -1
	}
	xTime[k] = 0
	for key, v := range xTime {
		fill(&xTime, xMap, key, v)
	}
	if len(xTime) < n {
		return -1
	}
	max := 0
	for key, v := range xTime {
		if key != k && max < v {
			max = v
		}
	}
	return max
}

func fill(xTime *map[int]int, xMap map[int]map[int]int, curK, curV int) {
	m, e := xMap[curK]
	if !e {
		return
	}
	for k, v := range m {
		o, e := (*xTime)[k]
		if !e {
			(*xTime)[k] = curV + v
			fill(xTime, xMap, k, curV+v)
		} else if curV+v < o {
			(*xTime)[k] = curV + v
			fill(xTime, xMap, k, curV + v)
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

type Arg struct {
	Times [][]int
	N     int
	K     int
}

func (args *Arg) getArgs() (times [][]int, n, k int) {
	return args.Times, args.N, args.K
}

func Test_S743(t *testing.T) {
	tests := []struct {
		Name     string
		Args     Arg
		Expected int
	}{
		{
			Name: "1",
			Args: Arg{
				Times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
				N:     4,
				K:     2,
			},
			Expected: 2,
		}, {
			Name: "2",
			Args: Arg{
				Times: [][]int{{1, 2, 1}},
				N:     2,
				K:     1,
			},
			Expected: 1,
		}, {
			Name: "3",
			Args: Arg{
				Times: [][]int{{1, 2, 1}},
				N:     2,
				K:     2,
			},
			Expected: -1,
		}, {
			Name: "4",
			Args: Arg{
				Times: [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}},
				N:     3,
				K:     1,
			},
			Expected: 3,
		}, {
			Name: "5",
			Args: Arg{
				Times: [][]int{{4, 2, 76}, {1, 3, 79}, {3, 1, 81}, {4, 3, 30},
					{2, 1, 47}, {1, 5, 61}, {1, 4, 99}, {3, 4, 68},
					{3, 5, 46}, {4, 1, 6}, {5, 4, 7}, {5, 3, 44},
					{4, 5, 19}, {2, 3, 13}, {3, 2, 18}, {1, 2, 0},
					{5, 1, 25}, {2, 5, 58}, {2, 4, 77}, {5, 2, 74}},
				N:     5,
				K:     3,
			},
			Expected: 59,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := networkDelayTime(test.Args.getArgs())
			arg, _ := json.Marshal(test.Args)
			t.Logf("arg: %v, expected result: %v, run result: %v", string(arg), test.Expected, result)
			if test.Expected != result {
				t.Fatalf("failed")
			}
		})
	}
}
