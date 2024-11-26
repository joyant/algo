package math

/**
1056. 易混淆数

给定一个数字 N，当它满足以下条件的时候返回 true：

原数字旋转 180° 以后可以得到新的数字。

如 0, 1, 6, 8, 9 旋转 180° 以后，得到了新的数字 0, 1, 9, 8, 6 。

2, 3, 4, 5, 7 旋转 180° 后，得到的不是数字。

易混淆数 (confusing number) 在旋转180°以后，可以得到和原来不同的数，且新数字的每一位都是有效的。
*/

// 题干的意思是把每个数字反过来，最后的结果组成的数字是否和原来一样
// 如果反过来就不是个数字，直接返回 false
func confusingNumber(n int) bool {
    var m = map[int]int{
        0: 0,
        1: 1,
        6: 9,
        8: 8,
        9: 6,
    }
    var ori = n
    var newVal = 0
    for n > 0 {
        num := n % 10
        v, ok := m[num]
        if !ok {
            return false
        }
        newVal = newVal*10 + v
        n /= 10
    }
    return newVal != ori
}
