package lcz

import (
	"testing"
)

//给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回该列名称对应的列序号。
//
//
//
// 例如，
//
//
//    A -> 1
//    B -> 2
//    C -> 3
//    ...
//    Z -> 26
//    AA -> 27
//    AB -> 28
//    ...
//
//
//
//
// 示例 1:
//
//
//输入: columnTitle = "A"
//输出: 1
//
//
// 示例 2:
//
//
//输入: columnTitle = "AB"
//输出: 28
//
//
// 示例 3:
//
//
//输入: columnTitle = "ZY"
//输出: 701
//
// 示例 4:
//
//
//输入: columnTitle = "FXSHRXW"
//输出: 2147483647
//
//
//
//
// 提示：
//
//
// 1 <= columnTitle.length <= 7
// columnTitle 仅由大写英文组成
// columnTitle 在范围 ["A", "FXSHRXW"] 内
//
// Related Topics 数学 字符串
// 👍 265 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func titleToNumber(columnTitle string) int {
	rowNum := 0
	for _, t := range columnTitle {
		rowNum = rowNum*26 + int(t-'A') + 1
	}
	return rowNum
}

//leetcode submit region end(Prohibit modification and deletion)

func Test_S171(t *testing.T) {
	tests := []struct {
		name        string
		columnTitle string
		expected    int
	}{
		{
			name:        "1",
			columnTitle: "A",
			expected:    1,
		}, {
			name:        "2",
			columnTitle: "AB",
			expected:    28,
		}, {
			name:        "3",
			columnTitle: "ZY",
			expected:    701,
		}, {
			name:        "4",
			columnTitle: "FXSHRXW",
			expected:    2147483647,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := titleToNumber(test.columnTitle)
			t.Logf("arg: %v, expected result: %v, run result: %v", test.columnTitle, test.expected, result)
			if result != test.expected {
				t.Fatalf("failed")
			}
		})
	}
}
