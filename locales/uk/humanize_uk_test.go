package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/uk"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[language.Tag]hc.Locale{
	language.Ukrainian: locale.Data,
}

func TestHumanizeUkOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 тисяча"},
		{"2000", "2 тисячі"},
		{"3000", "3 тисячі"},
		{"4000", "4 тисячі"},
		{"5000", "5 тисяч"},
		{"99900", "99,9 тисячі"},
		{"1000000", "1 мільйон"},
		{"2000000", "2 мільйони"},
		{"3000000", "3 мільйони"},
		{"4000000", "4 мільйони"},
		{"5000000", "5 мільйонів"},
		{"1000000000", "1 мільярд"},
		{"2000000000", "2 мільярди"},
		{"3000000000", "3 мільярди"},
		{"4000000000", "4 мільярди"},
		{"5000000000", "5 мільярдів"},
		{"1000000000000", "1 трильйон"},
		{"2000000000000", "2 трильйони"},
		{"3000000000000", "3 трильйони"},
		{"4000000000000", "4 трильйони"},
		{"5000000000000", "5 трильйонів"},
		{"1100", "1,1 тисячі"},
		{"2100", "2,1 тисячі"},
		{"3100", "3,1 тисячі"},
		{"4100", "4,1 тисячі"},
		{"5100", "5,1 тисячі"},
		{"1100000", "1,1 мільйона"},
		{"2100000", "2,1 мільйона"},
		{"3100000", "3,1 мільйона"},
		{"4100000", "4,1 мільйона"},
		{"5100000", "5,1 мільйона"},
		{"1100000000", "1,1 мільярда"},
		{"2100000000", "2,1 мільярда"},
		{"3100000000", "3,1 мільярда"},
		{"4100000000", "4,1 мільярда"},
		{"5100000000", "5,1 мільярда"},
		{"99900000000", "99,9 мільярда"},
		{"1100000000000", "1,1 трильйона"},
		{"2100000000000", "2,1 трильйона"},
		{"3100000000000", "3,1 трильйона"},
		{"4100000000000", "4,1 трильйона"},
		{"5100000000000", "5,1 трильйона"},
		{"99900000000000", "99,9 трильйона"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Ukrainian)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeUkOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 тис."},
		{"2000", "2 тис."},
		{"3000", "3 тис."},
		{"4000", "4 тис."},
		{"5000", "5 тис."},
		{"1100", "1,1 тис."},
		{"2100", "2,1 тис."},
		{"5100", "5,1 тис."},
		{"1000000", "1 млн"},
		{"2000000", "2 млн"},
		{"5000000", "5 млн"},
		{"1100000", "1,1 млн"},
		{"2100000", "2,1 млн"},
		{"5100000", "5,1 млн"},
		{"1000000000", "1 млрд"},
		{"2000000000", "2 млрд"},
		{"3000000000", "3 млрд"},
		{"4000000000", "4 млрд"},
		{"5000000000", "5 млрд"},
		{"99900000000", "99,9 млрд"},
		{"1100000000", "1,1 млрд"},
		{"2100000000", "2,1 млрд"},
		{"5100000000", "5,1 млрд"},
		{"1100000000000", "1,1 трлн"},
		{"2100000000000", "2,1 трлн"},
		{"5100000000000", "5,1 трлн"},
		{"1000000000000", "1 трлн"},
		{"2000000000000", "2 трлн"},
		{"3000000000000", "3 трлн"},
		{"4000000000000", "4 трлн"},
		{"5000000000000", "5 трлн"},
		{"99900000000000", "99,9 трлн"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Ukrainian)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
