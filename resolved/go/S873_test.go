package lcz

import (
	"encoding/json"
	"testing"
)

//如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：
//
//
// n >= 3
// 对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2}
//
//
// 给定一个严格递增的正整数数组形成序列 arr ，找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回 0 。
//
// （回想一下，子序列是从原序列 arr 中派生出来的，它从 arr 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如， [3, 5, 8]
//是 [3, 4, 5, 6, 7, 8] 的一个子序列）
//
//
//
//
//
//
// 示例 1：
//
//
//输入: arr = [1,2,3,4,5,6,7,8]
//输出: 5
//解释: 最长的斐波那契式子序列为 [1,2,3,5,8] 。
//
//
// 示例 2：
//
//
//输入: arr = [1,3,7,11,12,14,18]
//输出: 3
//解释: 最长的斐波那契式子序列有 [1,11,12]、[3,11,14] 以及 [7,11,18] 。
//
//
//
//
// 提示：
//
//
// 3 <= arr.length <= 1000
//
// 1 <= arr[i] < arr[i + 1] <= 10^9
//
//
// Related Topics 数组 哈希表 动态规划
// 👍 181 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func lenLongestFibSubseq(arr []int) int {
	if len(arr) < 3 {
		return 0
	}
	result := make([]*map[int]int, len(arr))
	max := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j] < arr[k] {
					break
				}
				if arr[i]+arr[j] == arr[k] {
					o := 0
					if result[i] == nil {
						o = 1
					} else if r, e := (*result[i])[j]; e {
						o = r + 1
					} else {
						o = 1
					}
					if result[j] == nil {
						result[j] = &map[int]int{k: o}
					} else {
						(*result[j])[k] = o
					}
					if o > max {
						max = o
					}
					break
				}
			}
		}
	}
	if max == 0 {
		return 0
	}
	return max + 2
}

//leetcode submit region end(Prohibit modification and deletion)
type Arg struct {
	Arr []int
}

func (args *Arg) getArgs() []int {
	return args.Arr
}

func Test_S873(t *testing.T) {
	tests := []struct {
		Name     string
		Args     Arg
		Expected int
	}{
		{
			Name: "1",
			Args: Arg{
				Arr: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			Expected: 5,
		}, {
			Name: "2",
			Args: Arg{
				Arr: []int{1, 3, 7, 11, 12, 14, 18},
			},
			Expected: 3,
		}, {
			Name: "3",
			Args: Arg{
				Arr: []int{2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50},
			},
			Expected: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := lenLongestFibSubseq(test.Args.getArgs())
			arg, _ := json.Marshal(test.Args)
			resultStr, _ := json.Marshal(result)
			expectedStr, _ := json.Marshal(test.Expected)
			t.Logf("arg: %v, expected result: %v, run result: %v", string(arg), string(expectedStr), string(resultStr))
			if string(expectedStr) != string(resultStr) {
				t.Fatalf("failed")
			}
		})
	}
}
