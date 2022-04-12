/*
 * @Description: Do not edit
 * @Date: 2022-03-25 17:56:17
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-03-25 18:36:16
 * @FilePath: /arithmetic/main.go
 */
package main

import (
	"fmt"

	"github.com/viletyy/arithmetic-go/jiuzhang"
)

func main() {

	fmt.Printf("jiuzhang.BinarySearch Result: %d\n", jiuzhang.BinarySearch([]int{1, 3, 5, 5, 7, 14, 25}, 5))
	fmt.Println("this arithmetic repo by golang.")
}
