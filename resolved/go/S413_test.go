package lcz

import (
	"encoding/json"
	"testing"
)

//如果一个数列 至少有三个元素 ，并且任意两个相邻元素之差相同，则称该数列为等差数列。
//
//
// 例如，[1,3,5,7,9]、[7,7,7,7] 和 [3,-1,-5,-9] 都是等差数列。
//
//
//
//
// 给你一个整数数组 nums ，返回数组 nums 中所有为等差数组的 子数组 个数。
//
// 子数组 是数组中的一个连续序列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,4]
//输出：3
//解释：nums 中有三个子等差数组：[1, 2, 3]、[2, 3, 4] 和 [1,2,3,4] 自身。
//
//
// 示例 2：
//
//
//输入：nums = [1]
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5000
// -1000 <= nums[i] <= 1000
//
//
//
// Related Topics 数组 动态规划
// 👍 314 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	count := 0
	diff := 3000
	index := 0
	total := 0
	for index < len(nums)-1 {
		if nums[index+1]-nums[index] != diff {
			diff = nums[index+1] - nums[index]
			total += (count + 1) * count / 2
			count = 0
		} else {
			count++
		}

		index++
	}
	total += (count + 1) * count / 2
	return total
}

//leetcode submit region end(Prohibit modification and deletion)
type Arg struct {
	Nums []int
}

func (args *Arg) getArgs() []int {
	return args.Nums
}

func Test_S413(t *testing.T) {
	tests := []struct {
		Name     string
		Args     Arg
		Expected int
	}{
		{
			Name: "1",
			Args: Arg{
				Nums: []int{1, 3, 5, 7, 9},
			},
			Expected: 6,
		}, {
			Name: "2",
			Args: Arg{
				Nums: []int{1, 2, 3, 4},
			},
			Expected: 3,
		}, {
			Name: "3",
			Args: Arg{
				Nums: []int{1},
			},
			Expected: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := numberOfArithmeticSlices(test.Args.getArgs())
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
