/*
 * @Description: Do not edit
 * @Date: 2022-04-15 14:25:32
 * @LastEditors: viletyy
 * @Author: viletyy
 * @LastEditTime: 2022-04-15 14:26:44
 * @FilePath: /arithmetic/leetcode/121.go
 */
package leetcode

/**
 * @title 买卖股票的最佳时机
 * @description: 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
 *							 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
 *							 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
 * @param {*} prices
 * @return {*}
 */
func maxProfit(prices []int) int {
	var minPrice, maxProfit int
	minPrice = prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
