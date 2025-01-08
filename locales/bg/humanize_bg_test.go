package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/bg"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[language.Tag]hc.Locale{
	language.Bulgarian: locale.Data,
}

func TestHumanizeBgOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1", "1"}, // fallback
		{"1000", "1 хил."},
		{"2000", "2 хиляди"},
		{"3000", "3 хиляди"},
		{"4000", "4 хиляди"},
		{"5000", "5 хиляди"},
		{"11000", "11 хиляди"},
		{"12000", "12 хиляди"},
		{"13000", "13 хиляди"},
		{"14000", "14 хиляди"},
		{"15000", "15 хиляди"},
		{"16000", "16 хиляди"},
		{"1000000", "1 милион"},
		{"1100000", "1,1 милион"},
		{"1200000", "1,2 милион"},
		{"1300000", "1,3 милион"},
		{"1400000", "1,4 милион"},
		{"1500000", "1,5 милиона"},
		{"1600000", "1,6 милиона"},
		{"1700000", "1,7 милиона"},
		{"1800000", "1,8 милиона"},
		{"1900000", "1,9 милиона"},
		{"10000000", "10 милиона"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Bulgarian)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeBgOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1\u00a0хил."},
		{"2000", "2\u00a0хил."},
		{"3000", "3\u00a0хил."},
		{"4000", "4\u00a0хил."},
		{"5000", "5\u00a0хил."},
		{"11000", "11\u00a0хил."},
		{"12000", "12\u00a0хил."},
		{"13000", "13\u00a0хил."},
		{"14000", "14\u00a0хил."},
		{"15000", "15\u00a0хил."},
		{"11100", "11,1\u00a0хил."},
		{"12100", "12,1\u00a0хил."},
		{"13100", "13,1\u00a0хил."},
		{"14100", "14,1\u00a0хил."},
		{"15100", "15,1\u00a0хил."},
		{"12500", "12,5\u00a0хил."},
		{"15000", "15\u00a0хил."},
		{"15100", "15,1\u00a0хил."},
		{"99500", "99,5\u00a0хил."},
		{"1000000", "1\u00a0млн."},
		{"1100000", "1,1\u00a0млн."},
		{"1200000", "1,2\u00a0млн."},
		{"1300000", "1,3\u00a0млн."},
		{"1400000", "1,4\u00a0млн."},
		{"1500000", "1,5\u00a0млн."},
		{"1600000", "1,6\u00a0млн."},
		{"1700000", "1,7\u00a0млн."},
		{"1800000", "1,8\u00a0млн."},
		{"1900000", "1,9\u00a0млн."},
		{"10000000", "10\u00a0млн."},
		{"11000000", "11\u00a0млн."},
		{"12000000", "12\u00a0млн."},
		{"13000000", "13\u00a0млн."},
		{"14000000", "14\u00a0млн."},
		{"15000000", "15\u00a0млн."},
		{"16000000", "16\u00a0млн."},
		{"17000000", "17\u00a0млн."},
		{"18000000", "18\u00a0млн."},
		{"19000000", "19\u00a0млн."},
		{"20000000", "20\u00a0млн."},
		{"21000000", "21\u00a0млн."},
		{"22000000", "22\u00a0млн."},
		{"23000000", "23\u00a0млн."},
		{"24000000", "24\u00a0млн."},
		{"25000000", "25\u00a0млн."},
		{"26000000", "26\u00a0млн."},
		{"27000000", "27\u00a0млн."},
		{"28000000", "28\u00a0млн."},
		{"1000000000", "1\u00a0млрд."},
		{"1100000000", "1,1\u00a0млрд."},
		{"1200000000", "1,2\u00a0млрд."},
		{"1300000000", "1,3\u00a0млрд."},
		{"1400000000", "1,4\u00a0млрд."},
		{"1500000000", "1,5\u00a0млрд."},
		{"1600000000", "1,6\u00a0млрд."},
		{"1000000000000", "1\u00a0трлн."},
		{"1100000000000", "1,1\u00a0трлн."},
		{"1200000000000", "1,2\u00a0трлн."},
		{"1300000000000", "1,3\u00a0трлн."},
		{"1400000000000", "1,4\u00a0трлн."},
		{"1500000000000", "1,5\u00a0трлн."},
		{"1600000000000", "1,6\u00a0трлн."},
		{"1700000000000", "1,7\u00a0трлн."},
		{"1800000000000", "1,8\u00a0трлн."},
		{"1900000000000", "1,9\u00a0трлн."},
		{"2000000000000", "2\u00a0трлн."},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Bulgarian)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
