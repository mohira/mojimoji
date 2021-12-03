package statement

import "testing"

func TestCount(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		s    string
		want int
	}{
		"日本語は1文字でカウント":     {"あいうえおあいうえお", 10},
		"英語は0.5文字はカウント":    {"abcdeabcde", 5},
		"数字と英語は0.5文字カウント":  {"あいうabc123", 6},
		"半角スペースは0.5文字カウント": {"あいう abc 123", 7},
		"半角記号は0.5文字カウント":   {"abc_def_gh()", 6},
		"小数点以下は切り上げ":       {"abcde", 3},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			got := Count(tt.s)
			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}

}
