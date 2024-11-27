package dynamic

/**
72. 编辑距离

给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
*/

/**
好的，让我们使用 "horse" 和 "ros" 作为例子。这个例子同时覆盖了 `word1[i-1] == word2[j-1]` 和 `word1[i-1] != word2[j-1]` 的情况。

我们的 DP 表格初始化如下：

```
    "" r o s
""  0  1 2 3
h   1
o   2
r   3
s   4
e   5
```

其中，"" 表示空字符串。

下面我们开始填表：

- 当 `i=1, j=1` 时，`word1[0] != word2[0]`（h != r），所以 `dp[1][1] = min(dp[0][0], dp[0][1], dp[1][0]) + 1 = 1`。

- 当 `i=1, j=2` 时，`word1[0] != word2[1]`（h != o），所以 `dp[1][2] = min(dp[0][1], dp[0][2], dp[1][1]) + 1 = 2`。

- 当 `i=1, j=3` 时，`word1[0] != word2[2]`（h != s），所以 `dp[1][3] = min(dp[0][2], dp[0][3], dp[1][2]) + 1 = 3`。

- 当 `i=2, j=1` 时，`word1[1] != word2[0]`（o != r），所以 `dp[2][1] = min(dp[1][0], dp[1][1], dp[2][0]) + 1 = 2`。

- 当 `i=2, j=2` 时，`word1[1] == word2[1]`（o == o），所以 `dp[2][2] = dp[1][1] = 1`。

- 当 `i=2, j=3` 时，`word1[1] != word2[2]`（o != s），所以 `dp[2][3] = min(dp[1][2], dp[1][3], dp[2][2]) + 1 = 2`。

以此类推，最后我们得到的 DP 表格如下：

```
    "" r o s
""  0  1 2 3
h   1  1 2 3
o   2  2 1 2
r   3  2 2 2
s   4  3 3 2
e   5  4 4 3
```

所以，将 "horse" 转换为 "ros" 需要的最少步骤数是 3。
*/
func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
    }
    // fill first element
    for i := 0; i < m+1; i++ {
        dp[i][0] = i
    }
    for j := 0; j < n+1; j++ {
        dp[0][j] = j
    }
    // calc
    for i := 1; i < m+1; i++ {
        for j := 1; j < n+1; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
            }
        }
    }
    return dp[m][n]
}
