/*
 * @Description: Do not edit
 * @Date: 2022-04-12 18:41:16
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-04-12 21:13:47
 * @FilePath: /arithmetic/leetcode/21.go
 */
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @title 合并两个有序链表
 * @description: 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 * @param {*ListNode} list1
 * @param {*ListNode} list2
 * @return {*ListNode}
 */
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	var newListNode *ListNode
	var prevAddListNode *ListNode
	for {
		var addListNode *ListNode
		if list1 == nil && list2 != nil {
			addListNode = list2
			list2 = nil
		}

		if list2 == nil && list1 != nil {
			addListNode = list1
			list1 = nil
		}

		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				addListNode = &ListNode{Val: list1.Val}
				list1 = list1.Next
			} else {
				addListNode = &ListNode{Val: list2.Val}
				list2 = list2.Next
			}
		}

		if newListNode == nil {
			newListNode = addListNode
			prevAddListNode = newListNode
		} else {
			midListNode := prevAddListNode
			midListNode.Next = addListNode
			prevAddListNode = addListNode
		}

		if list1 == nil && list2 == nil {
			break
		}

	}
	return newListNode
}
