package locale_test

import (
	"testing"

	"github.com/dejurin/humanize-cldr"
	locale "github.com/dejurin/humanize-cldr/locales/zh"
)

func fallback(number string) string {
	return number
}

func TestHumanizeZhOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1000"},
		{"1100", "1100"},
		{"10000", "1万"},
		{"100000", "10万"},
		{"1000000", "100万"},
		{"10000000", "1,000万"},
		{"100000000", "1亿"},
		{"1000000000", "10亿"},
		{"10000000000", "100亿"},
		{"100000000000", "1,000亿"},
		{"1000000000000", "1万亿"},
		{"10000000000000", "10万亿"},
		{"100000000000000", "100万亿"},
	}

	h := humanize.NewHumanizer(locale.Data, humanize.OptionLong, fallback)

	for _, tt := range tests {
		res, err := h.Humanize(tt.number)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeZhOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1000"},
		{"1100", "1100"},
		{"10000", "1万"},
		{"100000", "10万"},
		{"1000000", "100万"},
		{"10000000", "1,000万"},
		{"100000000", "1亿"},
		{"1000000000", "10亿"},
		{"10000000000", "100亿"},
		{"100000000000", "1,000亿"},
		{"1000000000000", "1万亿"},
		{"10000000000000", "10万亿"},
		{"100000000000000", "100万亿"},
	}

	h := humanize.NewHumanizer(locale.Data, humanize.OptionShort, fallback)

	for _, tt := range tests {
		res, err := h.Humanize(tt.number)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
