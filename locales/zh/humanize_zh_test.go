package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/zh"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[language.Tag]hc.Locale{
	language.Chinese: locale.Data,
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

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Chinese)
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

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Chinese)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
