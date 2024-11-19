package sort

func countingSortNaive(nums []int) {
    // 1. 统计数组最大元素 m
    m := 0
    for _, num := range nums {
        if num > m {
            m = num
        }
    }
    // 2. 统计各数字的出现次数
    // counter[num] 代表 num 的出现次数
    counter := make([]int, m+1)
    for _, num := range nums {
        counter[num]++
    }
    var idx = 0
    for index, num := range counter {
        if num > 0 {
            nums[idx] = index
            idx++
        }
    }
}
