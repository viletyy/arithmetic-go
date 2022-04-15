/*
 * @Description: Do not edit
 * @Date: 2022-04-15 14:15:42
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-04-15 14:16:05
 * @FilePath: /arithmetic/leetcode/350.go
 */
package leetcode

/**
 * @title 两个数组的交集2
 * @description: 给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
 * @param {*} result
 * @param {*} nums1
 * @return {*}
 */
func intersect(nums1 []int, nums2 []int) []int {
	hashMap := map[int]int{}
	var result []int
	if len(nums1) > len(nums2) {
		for i := 0; i < len(nums1); i++ {
			hashMap[nums1[i]] += 1
		}
		for j := 0; j < len(nums2); j++ {
			if v, has := hashMap[nums2[j]]; has && v > 0 {
				result = append(result, nums2[j])
				hashMap[nums2[j]] -= 1
			}
		}
	} else {
		for i := 0; i < len(nums2); i++ {
			hashMap[nums2[i]] += 1
		}
		for j := 0; j < len(nums1); j++ {
			if v, has := hashMap[nums1[j]]; has && v > 0 {
				result = append(result, nums1[j])
				hashMap[nums1[j]] -= 1
			}
		}
	}

	return result
}
