package lcz

import (
	"encoding/json"
	"testing"
)

//在一棵无限的二叉树上，每个节点都有两个子节点，树中的节点 逐行 依次按 “之” 字形进行标记。
//
// 如下图所示，在奇数行（即，第一行、第三行、第五行……）中，按从左到右的顺序进行标记；
//
// 而偶数行（即，第二行、第四行、第六行……）中，按从右到左的顺序进行标记。
//
//
//
// 给你树上某一个节点的标号 label，请你返回从根节点到该标号为 label 节点的路径，该路径是由途经的节点标号所组成的。
//
//
//
// 示例 1：
//
// 输入：label = 14
//输出：[1,3,4,14]
//
//
// 示例 2：
//
// 输入：label = 26
//输出：[1,2,6,10,26]
//
//
//
//
// 提示：
//
//
// 1 <= label <= 10^6
//
// Related Topics 树 数学 二叉树
// 👍 96 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func pathInZigZagTree(label int) []int {
	if label == 1 {
		return []int{1}
	}
	if label < 4 {
		return []int{1, label}
	}
	x := 1
	c := 0
	for label >= x {
		x *= 2
		c++
	}
	result := make([]int, c)
	result[c-1] = label
	result[0] = 1
	if c%2 == 1 {
		label = (label - x/2) / 2
	} else {
		label = (x - label - 1) / 2
	}
	x /= 2
	for i := c - 1; i > 1; i-- {
		if i%2 == 1 {
			result[i-1] = x/2 + label
		} else {
			result[i-1] = x - label - 1
		}
		x /= 2
		label /= 2
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func Test_S1104(t *testing.T) {
	tests := []struct {
		name           string
		label          int
		expectedResult []int
	}{
		{
			name:           "1",
			label:          14,
			expectedResult: []int{1, 3, 4, 14},
		}, {
			name:           "2",
			label:          26,
			expectedResult: []int{1, 2, 6, 10, 26},
		}, {
			name:           "3",
			label:          4,
			expectedResult: []int{1, 3, 4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := pathInZigZagTree(test.label)
			expectStr, _ := json.Marshal(test.expectedResult)
			resultStr, _ := json.Marshal(result)
			t.Logf("label: %d, expected expectedResult: %s, run result: %s", test.label, string(expectStr), string(resultStr))
			if string(expectStr) != string(resultStr) {
				t.Fatalf(" failed")
			}
		})
	}
}
