package lcz

import (
	"encoding/json"
	"testing"
)

//给你一个字符串 s 表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
//
//
// 'A'：Absent，缺勤
// 'L'：Late，迟到
// 'P'：Present，到场
//
//
// 如果学生能够 同时 满足下面两个条件，则可以获得出勤奖励：
//
//
// 按 总出勤 计，学生缺勤（'A'）严格 少于两天。
// 学生 不会 存在 连续 3 天或 3 天以上的迟到（'L'）记录。
//
//
// 如果学生可以获得出勤奖励，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：s = "PPALLP"
//输出：true
//解释：学生缺勤次数少于 2 次，且不存在 3 天或以上的连续迟到记录。
//
//
// 示例 2：
//
//
//输入：s = "PPALLL"
//输出：false
//解释：学生最后三天连续迟到，所以不满足出勤奖励的条件。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s[i] 为 'A'、'L' 或 'P'
//
// Related Topics 字符串 👍 99 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func checkRecord(s string) bool {
	p := false
	l := 0
	ll := false
	for _, i := range s {
		if i == 'A' {
			if p {
				return false
			}
			p = true
			ll = false
			l = 0
		} else if i == 'L' {
			if ll {
				if l > 1 {
					return false
				}
				l++
			} else {
				l++
				ll = true
			}
		} else {
			ll = false
			l = 0
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
type Arg struct {
	S string
}

func (args *Arg) getArgs() string {
	return args.S
}

func Test_S551(t *testing.T) {
	tests := []struct {
		Name     string
		Args     Arg
		Expected bool
	}{
		{
			Name: "1",
			Args: Arg{
				S: "PPALLP",
			},
			Expected: true,
		}, {
			Name: "2",
			Args: Arg{
				S: "PPALLL",
			},
			Expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := checkRecord(test.Args.getArgs())
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