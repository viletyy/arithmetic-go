/*
 * @Description: Do not edit
 * @Date: 2022-04-15 13:50:42
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-04-15 13:50:43
 * @FilePath: /arithmetic/leetcode/88.go
 */
/**
 * @title 合并两个有序数组
 * @description: 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
 * @param {*} nums1
 * @param {*} resultNums
 * @return {*}
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	var resultNums []int
	var curNums1, curNums2 int
	for {
		if curNums1 == m {
			resultNums = append(resultNums, nums2[curNums2:]...)
			break
		}
		if curNums2 == n {
			resultNums = append(resultNums, nums1[curNums1:]...)
			break
		}
		if nums1[curNums1] < nums2[curNums2] {
			resultNums = append(resultNums, nums1[curNums1])
			curNums1++
		} else {
			resultNums = append(resultNums, nums2[curNums2])
			curNums2++
		}

	}
	copy(nums1, resultNums)
}