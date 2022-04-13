/*
 * @Description: Do not edit
 * @Date: 2022-04-13 10:45:44
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-04-13 13:59:37
 * @FilePath: /arithmetic/leetcode/26.go
 */

package leetcode

/**
 * @title 删除有序数组中的重复线
 * @description: 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
 * @param {[]int} nums
 * @return {int}
 */
func removeDulicates(nums []int) int {
	i := 0
	j := 1
	for j < len(nums) {
		if nums[j] == nums[i] {
			j++
		} else {
			nums[i+1] = nums[j]
			i++
			j++
		}
	}

	return len(nums)
}
