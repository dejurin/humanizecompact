package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/ja"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[language.Tag]hc.Locale{
	language.Japanese: locale.Data,
}

func TestHumanizeJaOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1000"},
		{"2000", "2000"},
		{"10000", "1万"},
		{"20000", "2万"},
		{"12345", "12345"},
		{"100000", "10万"},
		{"1000000", "100万"},
		{"9990000", "999万"},
		{"10000000", "1,000万"},
		{"100000000", "1億"},
		{"1000000000000", "1兆"},
		{"10000000000000000", "1京"},
		{"1234.5678", "1234.5678"},
		{"100000.45", "100000.45"},
		{"123456789.99", "123456789.99"},
		{"1000000000000000000", "100京"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Japanese)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeJaOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1000"},
		{"2000", "2000"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Japanese)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
