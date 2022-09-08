package sensitive_words_filter

import (
	"sync"
)

const (
	FilterLevelLow = 1// filter level low
	FilterLevelMiddle = 2// filter level middle
	FilterLevelHight = 3// filter level hight
)

type SensitiveWordsFilter struct {
	level int
	skipDistance int
	replaceText rune
	root *Node
}

var (
	instance *SensitiveWordsFilter
	once sync.Once
)

// Get Instance
func GetInstance() *SensitiveWordsFilter {
	once.Do(func() {
		instance = &SensitiveWordsFilter{
			root: &Node{
				End:false,
			},
		}
	})
	return instance
}

// Build
func (filter *SensitiveWordsFilter) Build(words []string) {
	for _, word := range words {
		filter.root.AddWord(word)
	}
}

// Filter Text
func (filter *SensitiveWordsFilter) Filter(text string) (sensitiveWords []string, replaceText string) {
	textChars := []rune(text)
	textCharsCopy := make([]rune, len(textChars))
	copy(textCharsCopy, textChars)

	// get max distance
	maxDistance := filter.getMaxDistance(text)

	length := len(textChars)
	for i := 0; i < length; i++ {
		temp := filter.root.FindChild(textChars[i])
		if temp == nil {
			continue
		}
		j := i + 1
		// current distance
		distance := 0
		replaceLength := j
		for ; j < length; j++ {
			distance++
			if temp.FindChild(textChars[j]) == nil || distance > maxDistance {
				continue
			}
			if temp.End {
				sensitiveWords = append(sensitiveWords, string(textChars[i:j]))
				replaceRune(textCharsCopy, filter.replaceText, i, j)
			}
			temp = temp.FindChild(textChars[j])
			replaceLength = j
		}

		if j == length && temp != nil && temp.End {
			sensitiveWords = append(sensitiveWords, string(textChars[i:length]))
			end := replaceLength - 1
			if replaceLength > length - 1 {
				end = length - 1
			}
			replaceRune(textCharsCopy, filter.replaceText, i, end)
		}
	}

	return sensitiveWords, string(textCharsCopy)
}

// Replace Rune
func replaceRune(chars []rune, replaceChar rune, begin int, end int)  {
	for i := begin; i <= end; i++ {
		chars[i] = replaceChar
	}
}


// Get Filter Max Distance
func (filter *SensitiveWordsFilter) getMaxDistance(text string) int {
	var maxDistance int = 1
	switch filter.level {
	case FilterLevelLow:
		maxDistance = len(text) + 1
	case FilterLevelMiddle:
		maxDistance = filter.skipDistance + 1
	case FilterLevelHight:
		maxDistance = 1
	}
	return maxDistance
}

// Set Filter Level
func (filter *SensitiveWordsFilter) SetLevel(level int) {
	filter.level = level
}

// Set Skip Distance
func (filter *SensitiveWordsFilter) SetSkipDistance(skipDistance int) {
	filter.skipDistance = skipDistance
}

// Set Replace Text
func (filter *SensitiveWordsFilter) SetReplaceText(replaceText string) {
	replaceTextRune := []rune(replaceText)
	filter.replaceText = replaceTextRune[0]
}