/*
 * @Description: Do not edit
 * @Date: 2022-04-12 18:35:01
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-04-12 18:35:01
 * @FilePath: /arithmetic/leetcode/13.go
 */
package leetcode

/**
 * @title 罗马数字转整数
 * @description: 给定一个罗马数字，将其转换成整数。
 * @param {*} item
 * @return {*}
 */
func romanToInt(s string) int {
	var romanHashMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var value int
	var preItem string

	for _, item := range s {
		itemValue := romanHashMap[string(item)]
		preItemValue := romanHashMap[preItem]
		if preItem != "" && itemValue > preItemValue {
			value = value - 2*preItemValue + itemValue
		} else {
			value = value + itemValue
		}
		preItem = string(item)
	}

	return value
}
