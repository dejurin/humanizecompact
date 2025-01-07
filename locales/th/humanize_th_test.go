package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/th"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[language.Tag]hc.Locale{
	language.Thai: locale.Data,
}

func TestHumanizeThOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 พัน"},
		{"1100", "1.1 พัน"},
		{"10000", "1 หมื่น"},
		{"100000", "1 แสน"},
		{"1000000", "1 ล้าน"},
		{"10000000", "10 ล้าน"},
		{"100000000", "100 ล้าน"},
		{"1000000000", "1 พันล้าน"},
		{"10000000000", "1 หมื่นล้าน"},
		{"100000000000", "1 แสนล้าน"},
		{"1000000000000", "1 ล้านล้าน"},
		{"10000000000000", "10 ล้านล้าน"},
		{"100000000000000", "100 ล้านล้าน"},
		{"1000", "1 พัน"},
		{"1500", "1.5 พัน"},
		{"1600", "1.6 พัน"},
		{"2000", "2 พัน"},
		{"2500", "2.5 พัน"},
		{"9999", "9999"}, // fallback
		{"10000", "1 หมื่น"},
		{"23000", "2.3 หมื่น"},
		{"99999", "99999"}, // fallback
		{"100000", "1 แสน"},
		{"125000", "12.5 หมื่น"},
		{"1000000", "1 ล้าน"},
		{"1500000", "1.5 ล้าน"},
		{"9990000", "99.9 แสน"},
		{"10000000", "10 ล้าน"},
		{"12300000", "12.3 ล้าน"},
		{"100000000", "100 ล้าน"},
		{"999000000", "999 ล้าน"},
		{"1000000000", "1 พันล้าน"},
		{"1100000000", "1.1 พันล้าน"},
		{"9990000000", "9990000000"},
		{"10000000000", "1 หมื่นล้าน"},
		{"25000000000", "2.5 หมื่นล้าน"},
		{"100000000000", "1 แสนล้าน"},
		{"990000000000", "9.9 แสนล้าน"},
		{"1000000000000", "1 ล้านล้าน"},
		{"1600000000000", "1.6 ล้านล้าน"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Thai)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeThOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1K"},
		{"1500", "1.5K"},
		{"10000", "10K"},
		{"9999", "9999"}, // fallback
		{"1500000", "1.5M"},
		{"9990000", "9990000"}, // fallback
		{"1000000000", "1B"},
		{"1600000000000", "1.6T"},
		{"1000", "1K"},
		{"10000", "10K"},
		{"100000", "100K"},
		{"1000000", "1M"},
		{"1100000", "1.1M"},
		{"1200000", "1.2M"},
		{"1300000", "1.3M"},
		{"1400000", "1.4M"},
		{"1500000", "1.5M"},
		{"10000000", "10M"},
		{"100000000", "100M"},
		{"1000000000", "1B"},
		{"10000000000", "10B"},
		{"100000000000", "100B"},
		{"1000000000000", "1T"},
		{"10000000000000", "10T"},
		{"100000000000000", "100T"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Thai)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
