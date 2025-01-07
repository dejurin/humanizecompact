package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/ko"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[string]hc.Locale{
	"ko": locale.Data,
}

func TestHumanizeKoOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1천"},
		{"2000", "2천"},
		{"10000", "1만"},
		{"20000", "2만"},
		{"100000", "10만"},
		{"1000000", "100만"},
		{"9990000", "999만"},
		{"10000000", "1,000만"},
		{"1000000000000", "1조"},
		{"1000", "1천"},
		{"1100", "1.1천"},
		{"2500", "2.5천"},
		{"10000", "1만"},
		{"11000", "1.1만"},
		{"12500", "1.25만"},
		{"100000", "10만"},
		{"1000000", "100만"},
		{"10000000", "1,000만"},
		{"100000000", "1억"},
		{"230000000", "2.3억"},
		{"1000000000", "10억"},
		{"1230000000", "12.3억"},
		{"10000000000", "100억"},
		{"100000000000", "1,000억"},
		{"1000000000000", "1조"},
		{"1100000000000", "1.1조"},
		{"10000000000000", "10조"},
		{"12300000000000", "12.3조"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, err := h.Humanize(tt.number, language.Korean)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeKoOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1천"},
		{"2000", "2천"},
		{"10000", "1만"},
		{"20000", "2만"},
		{"100000", "10만"},
		{"1000000", "100만"},
		{"9990000", "999만"},
		{"10000000", "1,000만"},
		{"1000000000000", "1조"},
		{"1000", "1천"},
		{"1100", "1.1천"},
		{"2500", "2.5천"},
		{"10000", "1만"},
		{"11000", "1.1만"},
		{"12500", "1.25만"},
		{"100000", "10만"},
		{"1000000", "100만"},
		{"10000000", "1,000만"},
		{"100000000", "1억"},
		{"230000000", "2.3억"},
		{"1000000000", "10억"},
		{"1230000000", "12.3억"},
		{"10000000000", "100억"},
		{"100000000000", "1,000억"},
		{"1000000000000", "1조"},
		{"1100000000000", "1.1조"},
		{"10000000000000", "10조"},
		{"12300000000000", "12.3조"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, err := h.Humanize(tt.number, language.Korean)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
