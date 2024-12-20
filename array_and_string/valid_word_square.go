package array_and_string

/*
*
这个Go语言的函数`validWordSquare`用于检查一个给定的字符串数组是否构成一个有效的单词方阵。

首先，让我们理解什么是有效的单词方阵：一个有效的单词方阵是一个由字符串组成的二维数组，其中第 i 行和第 i 列的字符序列（从0开始）相同。例如，下面的单词方阵是有效的，因为对于所有的 i（从0开始），第 i 行和第 i 列的字符序列都相同：

```
"abcd"
"bnrt"
"crmy"
"dttb"
```

现在，让我们详细解析这个函数的工作原理：

1. 函数首先获取输入字符串数组的长度 `n`。如果输入字符串数组的长度为0，函数将返回 `false`，因为空数组不能构成有效的单词方阵。

2. 然后，函数通过两层循环来检查每一行和每一列的字符是否匹配。外层循环遍历数组的每一行，内层循环从当前行开始遍历到数组的最后一行。

3. 在内层循环中，如果当前行的索引 `i` 大于或等于某一列的长度，那么函数将返回 `true`。这是因为在一个有效的单词方阵中，如果某一行的长度小于或等于当前行的索引，那么这一行的剩余部分以及下面的所有行都应该是空的。

4. 然后，函数取出当前行和当前列的字符 `a` 和 `b`。如果这两个字符不相等，函数将返回 `false`，因为在一个有效的单词方阵中，每一行和对应的列的字符应该完全相同。

5. 如果函数遍历完所有的行和列都没有返回 `false`，那么函数将返回 `true`，表示输入的字符串数组构成了一个有效的单词方阵。

需要注意的是，这个函数假设输入的字符串数组中的所有字符串都是小写的，并且没有包含任何非字母的字符。如果输入的字符串数组可能包含大写字母或非字母的字符，那么这个函数可能需要进行一些修改才能正确地判断是否构成了一个有效的单词方阵。
*/
func validWordSquare(words []string) bool {
	n := len(words)
	if n == 0 {
		return false
	}
	for i := 0; i < n; i++ {
		cur := words[i]
		for j := i; j < n; j++ {
			if i >= len(words[j]) {
				return true
			}
			a := cur[j]
			b := words[j][i]
			if a != b {
				return false
			}
		}
	}
	return true
}
