package lcz

import "testing"

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
//
//
//
// 示例 1：
//
//
//输入：x = 121
//输出：true
//
//
// 示例 2：
//
//
//输入：x = -121
//输出：false
//解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
//
// 示例 3：
//
//
//输入：x = 10
//输出：false
//解释：从右向左读, 为 01 。因此它不是一个回文数。
//
//
// 示例 4：
//
//
//输入：x = -101
//输出：false
//
//
//
//
// 提示：
//
//
// -231 <= x <= 231 - 1
//
//
//
//
// 进阶：你能不将整数转为字符串来解决这个问题吗？
// Related Topics 数学
// 👍 1562 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	o := x
	n := 0
	for o >= 10 {
		n = n*10 + o%10
		o /= 10
	}
	return n*10+o%10 == x
}

//leetcode submit region end(Prohibit modification and deletion)

func Test_S9(t *testing.T) {
	tests := []struct {
		ArgsX  int
		Result bool
	}{
		{
			ArgsX:  121,
			Result: true,
		}, {
			ArgsX:  -121,
			Result: false,
		}, {
			ArgsX:  10,
			Result: false,
		}, {
			ArgsX:  -101,
			Result: false,
		},
	}
	for _, test := range tests {
		t.Run("Test_S9", func(t *testing.T) {
			result := isPalindrome(test.ArgsX)
			t.Logf("x: %d, expected expectedResult: %v, run expectedResult: %v", test.ArgsX, test.Result, result)
			if result != test.Result {
				t.Fatalf("Failed")
			}
		})
	}
}
