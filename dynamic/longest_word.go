package dynamic

import "sort"

func longestWord(words []string) string {
	// 按照长度和字典序对单词进行排序
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) > len(words[j])
	})

	wordSet := make(map[string]bool)
	for _, word := range words {
		if canForm(word, wordSet) {
			return word
		}
	}
	return ""
}

func canForm(word string, wordSet map[string]bool) bool {
	if len(wordSet) == 0 {
		wordSet[word] = true
		return false
	}

	dp := make([]bool, len(word)+1)
	dp[0] = true

	for i := 1; i <= len(word); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[word[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	// 在哈希集合中添加当前单词
	wordSet[word] = true
	return dp[len(word)]
}
