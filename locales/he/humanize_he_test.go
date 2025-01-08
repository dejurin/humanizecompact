package locale_test

import (
	"testing"

	"golang.org/x/text/language"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/he"
)

func fallback(number string) string {
	return number
}

var locales = map[language.Tag]hc.Locale{
	language.Hebrew: locale.Data,
}

func TestHumanizeHeOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "\u200f1 אלף"},
		{"2000", "\u200f2 אלף"},
		{"3000", "\u200f3 אלף"},
		{"4000", "\u200f4 אלף"},
		{"5000", "\u200f5 אלף"},
		{"10000", "\u200f10 אלף"},
		{"21000", "\u200f21 אלף"},
		{"31000", "\u200f31 אלף"},
		{"41000", "\u200f41 אלף"},
		{"51000", "\u200f51 אלף"},
		{"1000000", "\u200f1 מיליון"},
		{"2000000", "\u200f2 מיליון"},
		{"3000000", "\u200f3 מיליון"},
		{"4000000", "\u200f4 מיליון"},
		{"5000000", "\u200f5 מיליון"},
		{"1000000000", "\u200f1 מיליארד"},
		{"2000000000", "\u200f2 מיליארד"},
		{"3000000000", "\u200f3 מיליארד"},
		{"4000000000", "\u200f4 מיליארד"},
		{"5000000000", "\u200f5 מיליארד"},
		{"1000000000000", "\u200f1 טריליון"},
		{"2000000000000", "\u200f2 טריליון"},
		{"3000000000000", "\u200f3 טריליון"},
		{"4000000000000", "\u200f4 טריליון"},
		{"5000000000000", "\u200f5 טריליון"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Hebrew)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeHeOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1K\u200f"},
		{"10000", "10K\u200f"},
		{"100000", "100K\u200f"},
		{"1000000", "1M\u200f"},
		{"1100000", "1.1M\u200f"},
		{"1200000", "1.2M\u200f"},
		{"1300000", "1.3M\u200f"},
		{"1400000", "1.4M\u200f"},
		{"1500000", "1.5M\u200f"},
		{"10000000", "10M\u200f"},
		{"100000000", "100M\u200f"},
		{"1000000000", "1B\u200f"},
		{"10000000000", "10B\u200f"},
		{"100000000000", "100B\u200f"},
		{"1000000000000", "1T\u200f"},
		{"10000000000000", "10T\u200f"},
		{"100000000000000", "100T\u200f"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Hebrew)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
