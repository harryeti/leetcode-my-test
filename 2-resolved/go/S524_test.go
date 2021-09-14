package lcz

import (
	"fmt"
	"testing"
)

//给你一个字符串 s 和一个字符串数组 dictionary 作为字典，找出并返回字典中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
//
// 如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
//输出："apple"
//
//
// 示例 2：
//
//
//输入：s = "abpcplea", dictionary = ["a","b","c"]
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// 1 <= dictionary.length <= 1000
// 1 <= dictionary[i].length <= 1000
// s 和 dictionary[i] 仅由小写英文字母组成
//
// Related Topics 数组 双指针 字符串 排序 👍 200 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findLongestWord(s string, dictionary []string) string {
	max := ""
	for _, d := range dictionary {
		if (len(d) > len(max) && match(s, d)) || (len(d) == len(max) && d < max && match(s, d)) {
			max = d
		}
	}
	return max
}

func match(s, d string) bool {
	si := 0
	di := 0
	for di < len(d) {
		if si >= len(s) {
			return false
		}
		if len(d)-di > len(s)-si {
			return false
		}
		if s[si] == d[di] {
			di++
		}
		si++
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func Test_S524(t *testing.T) {
	result := findLongestWord("abce", []string{"abe", "abc", "monkey", "plea"})
	fmt.Printf("%v", result)
}
