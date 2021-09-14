package lcz

import (
	"encoding/json"
	"sort"
	"testing"
)

//给你一个整数数组 arr 和一个整数 k 。
//
// 设 m 为数组的中位数，只要满足下述两个前提之一，就可以判定 arr[i] 的值比 arr[j] 的值更强：
//
//
// |arr[i] - m| > |arr[j] - m|
// |arr[i] - m| == |arr[j] - m|，且 arr[i] > arr[j]
//
//
// 请返回由数组中最强的 k 个值组成的列表。答案可以以 任意顺序 返回。
//
// 中位数 是一个有序整数列表中处于中间位置的值。形式上，如果列表的长度为 n ，那么中位数就是该有序列表（下标从 0 开始）中位于 ((n - 1) / 2
//) 的元素。
//
//
// 例如 arr = [6, -3, 7, 2, 11]，n = 5：数组排序后得到 arr = [-3, 2, 6, 7, 11] ，数组的中间位置为 m
//= ((5 - 1) / 2) = 2 ，中位数 arr[m] 的值为 6 。
// 例如 arr = [-7, 22, 17, 3]，n = 4：数组排序后得到 arr = [-7, 3, 17, 22] ，数组的中间位置为 m = ((
//4 - 1) / 2) = 1 ，中位数 arr[m] 的值为 3 。
//
//
//
//
// 示例 1：
//
// 输入：arr = [1,2,3,4,5], k = 2
//输出：[5,1]
//解释：中位数为 3，按从强到弱顺序排序后，数组变为 [5,1,4,2,3]。最强的两个元素是 [5, 1]。[1, 5] 也是正确答案。
//注意，尽管 |5 - 3| == |1 - 3| ，但是 5 比 1 更强，因为 5 > 1 。
//
//
// 示例 2：
//
// 输入：arr = [1,1,3,5,5], k = 2
//输出：[5,5]
//解释：中位数为 3, 按从强到弱顺序排序后，数组变为 [5,5,1,1,3]。最强的两个元素是 [5, 5]。
//
//
// 示例 3：
//
// 输入：arr = [6,7,11,7,6,8], k = 5
//输出：[11,8,6,6,7]
//解释：中位数为 7, 按从强到弱顺序排序后，数组变为 [11,8,6,6,7,7]。
//[11,8,6,6,7] 的任何排列都是正确答案。
//
// 示例 4：
//
// 输入：arr = [6,-3,7,2,11], k = 3
//输出：[-3,11,2]
//
//
// 示例 5：
//
// 输入：arr = [-7,22,17,3], k = 2
//输出：[22,17]
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 10^5
// -10^5 <= arr[i] <= 10^5
// 1 <= k <= arr.length
//
// Related Topics 数组 双指针 排序
// 👍 20 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func getStrongest(arr []int, k int) []int {
	if len(arr) == 1 || len(arr) == k {
		return arr
	}
	sort.Ints(arr)
	s := 0
	e := len(arr) - 1
	m := arr[e/2]
	result := make([]int, k)
	for i := 0; i < k; i++ {
		if m-arr[s] > arr[e]-m {
			result[i] = arr[s]
			s++
		} else {
			result[i] = arr[e]
			e--
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func Test_S1471(t *testing.T) {
	tests := []struct {
		ArgsArr        []int
		ArgsK          int
		ExpectedResult []int
	}{
		{
			ArgsArr:        []int{1, 2, 3, 4, 5},
			ArgsK:          2,
			ExpectedResult: []int{1, 5},
		}, {
			ArgsArr:        []int{1, 1, 3, 5, 5},
			ArgsK:          2,
			ExpectedResult: []int{5, 5},
		}, {
			ArgsArr:        []int{6, 7, 11, 7, 6, 8},
			ArgsK:          5,
			ExpectedResult: []int{11, 8, 6, 6, 7},
		}, {
			ArgsArr:        []int{6, -3, 7, 2, 11},
			ArgsK:          3,
			ExpectedResult: []int{-3, 11, 2},
		}, {
			ArgsArr:        []int{-7, 22, 17, 3},
			ArgsK:          2,
			ExpectedResult: []int{22, 17},
		},
	}
	for _, test := range tests {
		t.Run("test_S1471", func(t *testing.T) {
			result := getStrongest(test.ArgsArr, test.ArgsK)
			sort.Ints(result)
			expect := test.ExpectedResult
			sort.Ints(expect)
			resultStr, _ := json.Marshal(result)
			expectStr, _ := json.Marshal(expect)
			t.Logf("arr: %v, k: %d, expected expectedResult: %s, run expectedResult: %s", test.ArgsArr, test.ArgsK, expectStr, resultStr)
			if string(resultStr) != string(expectStr) {
				t.Fatalf("failed")
			}
		})
	}
}
