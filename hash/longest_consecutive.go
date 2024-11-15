package hash

func longestConsecutive(nums []int) int {
    m := make(map[int]bool)
    for _, num := range nums {
        m[num] = true
    }
    var result = 0
    for _, num := range nums {
        // 说明 num 是连续串的第一个元素
        if !m[num-1] {
            current := num + 1
            for m[current] {
                current++
            }
            result = max(result, current-num)
        }
    }
    return result
}

// 利用 hash 实现
// 先初始化一个 map，然后遍历，如果一个value 是初始 value，判断这个串有多长
// 滑动窗口应该也是可以的，但要先排序
func longestConsecutive2(nums []int) int {
    numSet := map[int]bool{}
    for _, num := range nums {
        numSet[num] = true
    }
    longestStreak := 0
    for num := range numSet {
        if !numSet[num-1] {
            currentNum := num
            currentStreak := 1
            for numSet[currentNum+1] {
                currentNum++
                currentStreak++
            }
            if longestStreak < currentStreak {
                longestStreak = currentStreak
            }
        }
    }
    return longestStreak
}
