package lcz

import "sort"

//给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
//
// 示例 1:
//
//
//输入: [2,2,3,4]
//输出: 3
//解释:
//有效的组合是:
//2,3,4 (使用第一个 2)
//2,3,4 (使用第二个 2)
//2,2,3
//
//
// 注意:
//
//
// 数组长度不超过1000。
// 数组里整数的范围为 [0, 1000]。
//
// Related Topics 贪心 数组 双指针 二分查找 排序
// 👍 230 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func triangleNumber(nums []int) int {
	count := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			k := j + 1
			for k < len(nums) && nums[i]+nums[j] > nums[k] {
				k++
				count++
			}
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
