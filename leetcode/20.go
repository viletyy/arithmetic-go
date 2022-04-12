/*
 * @Description: Do not edit
 * @Date: 2022-04-12 18:38:14
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-04-12 18:38:14
 * @FilePath: /arithmetic/leetcode/20.go
 */
package leetcode

/**
 * @title 有效的括号
 * @description: 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 * @param {*} stackArray
 * @return {*}
 */
func isValid(s string) bool {
	var stackArray []string
	var validHashMap = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	var result, flag bool

	for _, item := range s {
		validItemValue := validHashMap[string(item)]
		if len(stackArray) == 0 && validItemValue == "" {
			flag = true
			continue
		} else {
			if validItemValue != "" {
				stackArray = append(stackArray, string(item))
				continue
			}

			if validItemValue == "" {
				if string(item) == validHashMap[stackArray[len(stackArray)-1]] {
					stackArray = stackArray[:len(stackArray)-1]
				} else {
					break
				}
			}
		}
	}

	if len(stackArray) == 0 && !flag {
		result = true
	}

	return result

}
