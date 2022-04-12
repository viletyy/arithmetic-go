/*
 * @Description: Do not edit
 * @Date: 2022-03-25 18:16:46
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-03-25 18:35:16
 * @FilePath: /arithmetic/jiuzhang/binary_search.go
 */
package jiuzhang

/**
 * @description:
 * @param {[]int} sortedArray
 * @return {int}
 */
func BinarySearch(sortedArray []int, target int) int {
	if len(sortedArray) == 0 {
		return -1
	}

	var start, end, mid int

	end = len(sortedArray) - 1

	for start+1 < end {
		mid = start + (end-start)/2

		if sortedArray[mid] == target {
			end = mid
		} else {
			if sortedArray[mid] < target {
				start = mid
			} else {
				end = mid
			}
		}
	}

	if sortedArray[start] == target {
		return start
	}

	if sortedArray[end] == target {
		return end
	}

	return -1
}
