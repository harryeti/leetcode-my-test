package lcz

import "sort"

//在一个整数数组中，“峰”是大于或等于相邻整数的元素，相应地，“谷”是小于或等于相邻整数的元素。例如，在数组{5, 8, 4, 2, 3, 4, 6}中，{8
//, 6}是峰， {5, 2}是谷。现在给定一个整数数组，将该数组按峰与谷的交替顺序排序。
//
// 示例:
//
// 输入: [5, 3, 1, 2, 3]
//输出: [5, 1, 3, 2, 3]
//
//
// 提示：
//
//
// nums.length <= 10000
//
// Related Topics 贪心 数组 排序
// 👍 32 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func wiggleSort(nums []int) {
	if len(nums) > 2 {
		sort.Ints(nums)
		start := 1
		for start < len(nums)-1 {
			tmp := nums[start]
			nums[start] = nums[start+1]
			nums[start+1] = tmp
			start += 2
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
