package lcz

import (
	"encoding/json"
	"testing"
)

//给你两个整数 n 和 k ，请你构造一个答案列表 answer ，该列表应当包含从 1 到 n 的 n 个不同正整数，并同时满足下述条件：
//
//
// 假设该列表是 answer = [a1, a2, a3, ... , an] ，那么列表 [|a1 - a2|, |a2 - a3|, |a3 - a4|
//, ... , |an-1 - an|] 中应该有且仅有 k 个不同整数。
//
//
// 返回列表 answer 。如果存在多种答案，只需返回其中 任意一种 。
//
//
//
// 示例 1：
//
//
//输入：n = 3, k = 1
//输出：[1, 2, 3]
//解释：[1, 2, 3] 包含 3 个范围在 1-3 的不同整数，并且 [1, 1] 中有且仅有 1 个不同整数：1
//
//
// 示例 2：
//
//
//输入：n = 3, k = 2
//输出：[1, 3, 2]
//解释：[1, 3, 2] 包含 3 个范围在 1-3 的不同整数，并且 [2, 1] 中有且仅有 2 个不同整数：1 和 2
//
//
//
//
// 提示：
//
//
// 1 <= k < n <= 104
//
//
//
// Related Topics 数组 数学
// 👍 91 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func constructArray(n int, k int) []int {
	var high bool
	result := make([]int, n)
	num := 1
	index := 0
	for i := k; i >= 0; i-- {
		result[index] = num
		if high {
			num -= i
			high = false
		} else {
			num += i
			high = true
		}
		index++
	}
	num = 1 + k + 1
	for index < n {
		result[index] = num
		num++
		index++
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func Test_S667(t *testing.T) {
	tests := []struct {
		N      int
		K      int
		Result []int
	}{
		{
			N:      3,
			K:      1,
			Result: []int{1, 2, 3},
		},
		{
			N:      3,
			K:      2,
			Result: []int{1, 3, 2},
		},
	}
	for _, test := range tests {
		t.Run("Test_S667", func(t *testing.T) {
			j, _ := json.Marshal(test.Result)
			t.Logf(string(j))
			arr := constructArray(test.N, test.K)
			j, _ = json.Marshal(arr)
			t.Logf(string(j))
			if len(arr) != len(test.Result) {
				t.Fatalf("failed")
			}
			for i := range arr {
				if arr[i] != test.Result[i] {
					t.Fatalf("failed")
				}
			}
		})
	}
}
