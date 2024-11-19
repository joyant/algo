package hash

/**
697. 数组的度
给定一个非空且只包含非负数的整数数组 nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。

你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。
*/

/**
用例子读懂题目：
当然，我会举更多的例子来帮助你理解这个问题。

1. **例子一**：如果数组是 [1, 2, 2, 3, 1, 4, 2]，那么这个数组的度就是 3，因为 2 出现了三次，这是所有元素中出现次数最多的。满足条件的子数组有 [2, 2, 3, 1, 4, 2] 和 [1, 2, 2, 3, 1, 4, 2]，他们的度都是 3，但是我们需要找到最短的子数组，所以答案是 [2, 2, 3, 1, 4, 2]，长度为 6。

2. **例子二**：如果数组是 [1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 4]，那么这个数组的度就是 4，因为 4 出现了四次，这是所有元素中出现次数最多的。满足条件的子数组只有一个，就是 [4, 4, 4, 4]，所以答案是这个子数组，长度为 4。

3. **例子三**：如果数组是 [1, 2, 3, 4, 5]，那么这个数组的度就是 1，因为所有的元素都只出现了一次。在这种情况下，任何一个由单一元素组成的子数组都满足条件，所以答案是任何一个长度为 1 的子数组。

以上的例子都是在假设数组中的元素是整数的情况下给出的，但是这个问题的解决方法对于任何类型的元素都是适用的，只要这些元素可以被比较。希望这些例子能帮助你理解这个问题。
*/

// 利用 hash 解决
func findShortestSubArray(nums []int) int {
    count, first, last := map[int]int{}, map[int]int{}, map[int]int{}
    maxCount, minLen := 0, 0
    for index, num := range nums {
        if _, ok := first[num]; !ok {
            first[num] = index
        }
        last[num] = index
        count[num] += 1
        // 判断
        distance := last[num] - first[num] + 1
        if count[num] > maxCount {
            maxCount = count[num]
            minLen = distance
        } else if count[num] == maxCount {
            minLen = min(minLen, distance)
        }
    }
    return minLen
}
