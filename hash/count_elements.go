package hash

/**
1426. 数元素
level: easy
给你一个整数数组 arr， 对于元素 x ，只有当 x + 1 也在数组 arr 里时，才能记为 1 个数。

如果数组 arr 里有重复的数，每个重复的数单独计算。



示例 1：

输入：arr = [1,2,3]
输出：2
解释：1 和 2 被计算次数因为 2 和 3 在数组 arr 里。
示例 2：

输入：arr = [1,1,3,3,5,5,7,7]
输出：0
解释：所有的数都不算, 因为数组里没有 2、4、6、8。

*/

func countElements(arr []int) int {
    m := map[int]int{}
    for _, num := range arr {
        m[num]++
    }
    res := 0
    for _, num := range arr {
        // 看下一个数是否也在 m 中
        _, ok := m[num+1]
        if ok {
            res++
        }
    }
    return res
}
