/*
 * @Description: Do not edit
 * @Date: 2022-04-12 19:28:45
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-04-12 19:28:45
 * @FilePath: /arithmetic/leetcode/main.go
 */
package leetcode

func Run() {

	// 合并两个有序链表
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	mergeTwoLists(list1, list2)
}
