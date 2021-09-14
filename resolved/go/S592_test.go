package lcz

import (
	"fmt"
	"testing"
)

//给定一个表示分数加减运算表达式的字符串，你需要返回一个字符串形式的计算结果。 这个结果应该是不可约分的分数，即最简分数。 如果最终结果是一个整数，例如 2，
//你需要将它转换成分数形式，其分母为 1。所以在上述例子中, 2 应该被转换为 2/1。
//
// 示例 1:
//
//
//输入:"-1/2+1/2"
//输出: "0/1"
//
//
// 示例 2:
//
//
//输入:"-1/2+1/2+1/3"
//输出: "1/3"
//
//
// 示例 3:
//
//
//输入:"1/3-1/2"
//输出: "-1/6"
//
//
// 示例 4:
//
//
//输入:"5/3+1/3"
//输出: "2/1"
//
//
// 说明:
//
//
// 输入和输出字符串只包含 '0' 到 '9' 的数字，以及 '/', '+' 和 '-'。
// 输入和输出分数格式均为 ±分子/分母。如果输入的第一个分数或者输出的分数是正数，则 '+' 会被省略掉。
// 输入只包含合法的最简分数，每个分数的分子与分母的范围是 [1,10]。 如果分母是1，意味着这个分数实际上是一个整数。
// 输入的分数个数范围是 [1,10]。
// 最终结果的分子与分母保证是 32 位整数范围内的有效整数。
//
// Related Topics 数学 字符串 模拟 👍 53 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func fractionAddition(expression string) string {
	rch := 0
	rma := 1
	la := '+'
	cch := 0
	cma := -1
	for _, e := range expression {
		switch e {
		case '+':
			fallthrough
		case '-':
			if cma > 0 {
				if la == '+' {
					rch, rma = calc(rch, rma, cch, cma)
				} else {
					rch, rma = calc(rch, rma, -cch, cma)
				}
				cch = 0
				cma = -1
			}
			la = e
		case '/':
			cma = 0
		case '0':
			if cma < 0 {
				cch *= 10
			} else {
				cma *= 10
			}
		case '1':
			if cma < 0 {
				cch = cch*10 + 1
			} else {
				cma = cma*10 + 1
			}
		case '2':
			if cma < 0 {
				cch = cch*10 + 2
			} else {
				cma = cma*10 + 2
			}
		case '3':
			if cma < 0 {
				cch = cch*10 + 3
			} else {
				cma = cma*10 + 3
			}
		case '4':
			if cma < 0 {
				cch = cch*10 + 4
			} else {
				cma = cma*10 + 4
			}
		case '5':
			if cma < 0 {
				cch = cch*10 + 5
			} else {
				cma = cma*10 + 5
			}
		case '6':
			if cma < 0 {
				cch = cch*10 + 6
			} else {
				cma = cma*10 + 6
			}
		case '7':
			if cma < 0 {
				cch = cch*10 + 7
			} else {
				cma = cma*10 + 7
			}
		case '8':
			if cma < 0 {
				cch = cch*10 + 8
			} else {
				cma = cma*10 + 8
			}
		case '9':
			if cma < 0 {
				cch = cch*10 + 9
			} else {
				cma = cma*10 + 9
			}
		}
	}
	if la == '+' {
		rch, rma = calc(rch, rma, cch, cma)
	} else {
		rch, rma = calc(rch, rma, -cch, cma)
	}
	return fmt.Sprintf("%d/%d", rch, rma)
}

func calc(rch, rma, cch, cma int) (nch, nma int) {
	nma = rma * cma
	nch = rch*cma + cch*rma
	if nch == 0 {
		return 0, 1
	}
	i := 2
	for ((nch > 0 && i <= nch) || (nch < 0 && i <= -nch)) && i <= nma {
		if nch%i == 0 && nma%i == 0 {
			nch /= i
			nma /= i
			i = 1
		}
		i++
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestS592(t *testing.T) {
	println(fractionAddition("7/3+5/2-3/10"))
}
