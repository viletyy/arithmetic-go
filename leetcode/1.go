/*
 * @Description: Do not edit
 * @Date: 2022-04-12 18:29:41
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-04-12 18:32:35
 * @FilePath: /arithmetic/leetcode/1.go
 */
package leetcode

/**
 * @title 两数之和
 * @description: 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
 * @param {[]int} nums
 * @param {int} target
 * @return {*}
 */
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
