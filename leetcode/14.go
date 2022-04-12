/*
 * @Description: Do not edit
 * @Date: 2022-04-12 18:36:35
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-04-12 18:36:36
 * @FilePath: /arithmetic/leetcode/14.go
 */
package leetcode

/**
 * @title 最长公共前缀
 * @description: 编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。
 * @param {*} strPrefix
 * @return {*}
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	var prefixStr string = strs[0]
	for _, str := range strs {
		var strPrefix []rune
		for j, strRune := range str {
			if len(prefixStr) > j && string(prefixStr[j]) == string(strRune) {
				strPrefix = append(strPrefix, strRune)
			} else {
				break
			}
		}
		if len(strPrefix) > 0 {
			prefixStr = string(strPrefix)
		} else {
			prefixStr = ""
		}
	}

	return prefixStr
}
