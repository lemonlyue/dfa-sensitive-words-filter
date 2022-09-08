package sensitive_words_filter

import (
	"testing"
)

var tests = []struct{
	level int
	skipDistance int
	text string
	replaceText string
	out string
}{
	{FilterLevelLow, 1, "你好丑", "帅", "你好帅"},
	{FilterLevelMiddle, 1, "你好丑", "帅", "你好帅"},
	{FilterLevelHight, 1, "你好丑", "帅", "你好帅"},
	{FilterLevelLow, 1, "你好傻逼", "*", "你好**"},
	{FilterLevelMiddle, 3, "你好傻  逼", "*", "你好****"},
	{FilterLevelHight, 1, "你好傻 逼", "*", "你好* 逼"},
	{FilterLevelHight, 1, "你好傻逼啊", "*", "你好**啊"},
}

func TestFilter(t *testing.T)  {
	filter := GetInstance()
	filter.Build([]string{"傻逼", "丑", "傻子", "傻"})
	for i, tt := range tests {
		filter.SetLevel(tt.level)
		filter.SetSkipDistance(tt.skipDistance)
		filter.SetReplaceText(tt.replaceText)
		_, text:= filter.Filter(tt.text)
		if text != tt.out {
			t.Errorf("%d . %q => %q, wanted: %q", i, tt.text, text, tt.out)
		}
	}
}
