package array_and_string

import "math"

/**
1002. 查找共用字符
level: easy
给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答案。


示例 1：
输入：words = ["bella","label","roller"]
输出：["e","l","l"]

示例 2：
输入：words = ["cool","lock","cook"]
输出：["c","o"]
*/

func commonChars(words []string) []string {
	res := [26]int{}
	// 填充最大值
	for i := range res {
		res[i] = math.MaxInt
	}
	// 统计每个单词
	for _, word := range words {
		freq := [26]int{}
		for _, w := range word {
			freq[w-'a']++
		}
		// 把统计结果和 res 结合起来，因为统计的是「每个单词都出现的」，所以出现次数最少的才是答案
		for i, f := range freq {
			if f < res[i] {
				res[i] = f
			}
		}
	}
	// 将 res 中的「字符串」提取出来
	var slice []string
	for char, num := range res {
		// 可能一个字符出现多次
		for i := 0; i < num; i++ {
			slice = append(slice, string(char+'a'))
		}
	}
	return slice
}
