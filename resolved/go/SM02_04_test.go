package lcz

import (
	"fmt"
	"testing"
)

//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
// 你应当 保留 两个分区中每个节点的初始相对位置。
//
//
//
// 示例 1：
//
//
//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]
//
//
// 示例 2：
//
//
//输入：head = [2,1], x = 2
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 200] 内
// -100 <= Node.val <= 100
// -200 <= x <= 200
//
// Related Topics 链表 双指针
// 👍 65 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	cur := head
	var next *ListNode
	var less *ListNode
	var first *ListNode
	var more *ListNode
	var mid *ListNode
	for cur != nil {
		next = cur.Next
		if cur.Val < x {
			if less == nil {
				less = cur
				first = cur
			} else {
				less.Next = cur
				less = less.Next
			}
		} else {
			if more == nil {
				more = cur
				mid = cur
			} else {
				more.Next = cur
				more = more.Next
			}
		}
		cur = next
	}
	if less == nil {
		if more != nil {
			more.Next = nil
		}
		return mid
	} else {
		if more != nil {
			more.Next = nil
		}
		less.Next = mid
		return first
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test(t *testing.T) {
	tests := []struct {
		Arr    []int
		Num    int
		Result string
	}{
		{
			Arr:    []int{1, 4, 3, 2, 5, 2},
			Num:    3,
			Result: "[ 1,2,2,4,3,5 ]",
		}, {
			Arr:    []int{2, 1},
			Num:    2,
			Result: "[ 1,2 ]",
		},
	}
	for _, test := range tests {
		t.Run("test_SM02_04", func(t *testing.T) {
			node := buildList(test.Arr)
			t.Logf(buildArrStr(node))
			result := partition(node, test.Num)
			str := buildArrStr(result)
			t.Logf(str)
			if str != test.Result {
				t.Fatalf("expectedResult is incorrect")
			}
		})
	}
}

func buildList(arr []int) *ListNode {
	if len(arr) < 1 {
		return nil
	}
	first := &ListNode{Val: arr[0]}
	node := first
	if len(arr) == 1 {
		return node
	}
	for i := 1; i < len(arr); i++ {
		node.Next = &ListNode{Val: arr[i]}
		node = node.Next
	}
	return first
}

func buildArrStr(node *ListNode) string {
	cp := node
	str := "[ "
	for cp != nil {
		str += fmt.Sprintf("%d,", cp.Val)
		cp = cp.Next
	}
	str = str[0 : len(str)-1]
	str += " ]"
	return str
}

//leetcode submit region end(Prohibit modification and deletion)
