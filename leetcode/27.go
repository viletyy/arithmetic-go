/*
 * @Description: Do not edit
 * @Date: 2022-04-13 14:09:28
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-04-13 14:17:33
 * @FilePath: /arithmetic/leetcode/27.go
 */
package leetcode

/**
 * @title 移除元素
 * @description: 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
 * @param {*} nums
 * @return {*}
 */
func removeElement(nums []int, val int) int {
	var i, j int = 0, 0

	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
			j++
		} else {
			j++
		}
	}

	nums = nums[:i+1]
	return len(nums)
}
