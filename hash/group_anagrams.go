package hash

import (
    "sort"
    "strings"
)

/**
leetcode: 49. 字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]
*/

/**
将每个字符串按字符排序，直接用等号比较，就能知道是不是字母异位词
所以 split、sort、join 操作是少不了的
剩下的就是比较和保存结果就好了
*/
func groupAnagrams(strs []string) [][]string {
    var result [][]string
    var m = make(map[string][]string)
    for _, s := range strs {
        slice := strings.Split(s, "")
        sort.Strings(slice)
        s2 := strings.Join(slice, "")
        m[s2] = append(m[s2], s)
    }
    for _, v := range m {
        result = append(result, v)
    }
    return result
}
