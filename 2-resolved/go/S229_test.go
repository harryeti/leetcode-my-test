package lcz

//给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：[3,2,3]
//输出：[3]
//
// 示例 2：
//
//
//输入：nums = [1]
//输出：[1]
//
//
// 示例 3：
//
//
//输入：[1,1,1,3,3,2,2,2]
//输出：[1,2]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 104
// -109 <= nums[i] <= 109
//
// Related Topics 数组 哈希表 计数 排序
// 👍 386 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) []int {
	m := map[int]int{}
	for _, num := range nums {
		if _, e := m[num]; e {
			m[num]++
		} else {
			m[num] = 1
		}
	}
	l := len(nums) / 3
	result := []int{}
	for k, v := range m {
		if v > l {
			result = append(result, k)
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
