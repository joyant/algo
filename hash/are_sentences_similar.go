package hash

/**
734. 句子相似性

如果 sentence1 和 sentence2 相似则返回 true ，如果不相似则返回 false 。

两个句子是相似的，如果:

它们具有 相同的长度 (即相同的字数)
sentence1[i] 和 sentence2[i] 是相似的
请注意，一个词总是与它自己相似，也请注意，相似关系是不可传递的。例如，如果单词 a 和 b 是相似的，单词 b 和 c 也是相似的，那么 a 和 c  不一定相似 。



示例 1:

输入: sentence1 = ["great","acting","skills"], sentence2 = ["fine","drama","talent"], similarPairs = [["great","fine"],["drama","acting"],["skills","talent"]]
输出: true
解释: 这两个句子长度相同，每个单词都相似。
示例 2:

输入: sentence1 = ["great"], sentence2 = ["great"], similarPairs = []
输出: true
解释: 一个单词和它本身相似。
示例 3:

输入: sentence1 = ["great"], sentence2 = ["doubleplus","good"], similarPairs = [["great","doubleplus"]]
输出: false
解释: 因为它们长度不同，所以返回false。
*/

// 题很简单，关键是读懂题干，需要仔细看一下示例
func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
    if len(sentence1) != len(sentence2) {
        return false
    }
    m := map[string]bool{}
    for _, p := range similarPairs {
        m[p[0]+"#"+p[1]] = true
    }
    for i := range sentence1 {
        a, b := sentence1[i], sentence2[i]
        x, y := a+"#"+b, b+"#"+a
        _, ok1 := m[x]
        _, ok2 := m[y]
        if len(a) != len(b) && !ok1 && !ok2 {
            return false
        }
    }
    return true
}
