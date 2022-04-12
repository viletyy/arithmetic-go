/*
 * @Description: Do not edit
 * @Date: 2022-04-12 18:33:16
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-04-12 18:33:16
 * @FilePath: /arithmetic/leetcode/2.go
 */
package leetcode

/**
 * @title 回文数
 * @description: 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
 * @param {int} x
 * @return {*}
 */
func isPalindrome(x int) bool {
	var source int = x
	var after int
	for x > 0 {
		after = after*10 + x%10
		x = x / 10
	}
	if after == source {
		return true
	}

	return false
}
