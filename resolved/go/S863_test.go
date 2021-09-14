package lcz

//给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 K 。
//
// 返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。
//
//
//
//
//
//
// 示例 1：
//
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
//输出：[7,4,1]
//解释：
//所求结点为与目标结点（值为 5）距离为 2 的结点，
//值分别为 7，4，以及 1
//
//
//
//注意，输入的 "root" 和 "target" 实际上是树上的结点。
//上面的输入仅仅是对这些对象进行了序列化描述。
//
//
//
//
// 提示：
//
//
// 给定的树是非空的。
// 树上的每个结点都具有唯一的值 0 <= node.val <= 500 。
// 目标结点 target 是树上的结点。
// 0 <= K <= 1000.
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树
// 👍 329 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	result, _ := needCalc(root, target, k, 0)
	return result
}

func needCalc(node, target *TreeNode, k, path int) ([]int, int) {
	if node == nil {
		return []int{}, -1
	}
	if node == target {
		return calc(node, k, 2), path
	}
	childL, pL := needCalc(node.Left, target, k, path+1)
	childR, pR := needCalc(node.Right, target, k, path+1)
	if pL > 0 {
		return merge(childL, calc(node, k+path-pL, 1)), pL
	} else if pR > 0 {
		return merge(childR, calc(node, k+path-pR, 0)), pR
	}
	return []int{}, -1
}

func calc(node *TreeNode, k, lr int) []int {
	if node == nil || k < 0 {
		return []int{}
	}
	if k == 0 {
		return []int{node.Val}
	}
	if lr == 0 {
		return calc(node.Left, k-1, 2)
	} else if lr == 1 {
		return calc(node.Right, k-1, 2)
	}
	return merge(calc(node.Left, k-1, 2), calc(node.Right, k-1, 2))

}

func merge(a, b []int) []int {
	al := len(a)
	arr := make([]int, al+len(b))
	for i, e := range a {
		arr[i] = e
	}
	for i, e := range b {
		arr[al+i] = e
	}
	return arr
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
