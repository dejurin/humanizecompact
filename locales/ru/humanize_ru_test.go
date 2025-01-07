package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"

	locale "github.com/dejurin/humanizecompact/locales/ru"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[string]hc.Locale{
	"ru": locale.Data,
}

func TestHumanizeRuOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1", "1"}, // fallback
		{"1000", "1 тысяча"},
		{"2000", "2 тысячи"},
		{"3000", "3 тысячи"},
		{"4000", "4 тысячи"},
		{"5000", "5 тысяч"},
		{"10000", "10 тысяч"},
		{"11000", "11 тысяч"},
		{"12000", "12 тысяч"},
		{"13000", "13 тысяч"},
		{"14000", "14 тысяч"},
		{"15000", "15 тысяч"},
		{"11100", "11,1 тысячи"},
		{"12100", "12,1 тысячи"},
		{"13100", "13,1 тысячи"},
		{"14100", "14,1 тысячи"},
		{"15100", "15,1 тысячи"},
		{"12500", "12,5 тысячи"},
		{"15000", "15 тысяч"},
		{"15100", "15,1 тысячи"},
		{"99500", "99,5 тысячи"},
		{"1000000", "1 миллион"},
		{"10100000", "10,1 миллиона"},
		{"1100000", "1,1 миллиона"},
		{"1200000", "1,2 миллиона"},
		{"1300000", "1,3 миллиона"},
		{"1400000", "1,4 миллиона"},
		{"1500000", "1,5 миллиона"},
		{"1600000", "1,6 миллиона"},
		{"1700000", "1,7 миллиона"},
		{"1800000", "1,8 миллиона"},
		{"1900000", "1,9 миллиона"},
		{"99900000", "99,9 миллиона"},
		{"2000000", "2 миллиона"},
		{"2300000", "2,3 миллиона"},
		{"1700000000", "1,7 миллиарда"},
		{"1000000000", "1 миллиард"},
		{"1100000000", "1,1 миллиарда"},
		{"1200000000", "1,2 миллиарда"},
		{"1300000000", "1,3 миллиарда"},
		{"1400000000", "1,4 миллиарда"},
		{"1500000000", "1,5 миллиарда"},
		{"1600000000", "1,6 миллиарда"},
		{"1700000000", "1,7 миллиарда"},
		{"1800000000", "1,8 миллиарда"},
		{"1900000000", "1,9 миллиарда"},
		{"2000000000", "2 миллиарда"},
		{"3000000000", "3 миллиарда"},
		{"4000000000", "4 миллиарда"},
		{"5000000000", "5 миллиардов"},
		{"6000000000", "6 миллиардов"},
		{"7000000000", "7 миллиардов"},
		{"8000000000", "8 миллиардов"},
		{"9000000000", "9 миллиардов"},
		{"10000000000", "10 миллиардов"},
		{"11000000000", "11 миллиардов"},
		{"12000000000", "12 миллиардов"},
		{"13000000000", "13 миллиардов"},
		{"14000000000", "14 миллиардов"},
		{"15000000000", "15 миллиардов"},
		{"16000000000", "16 миллиардов"},
		{"17000000000", "17 миллиардов"},
		{"18000000000", "18 миллиардов"},
		{"19000000000", "19 миллиардов"},
		{"20000000000", "20 миллиардов"},
		{"100000000000", "100 миллиардов"},
		{"200000000000", "200 миллиардов"},
		{"300000000000", "300 миллиардов"},
		{"400000000000", "400 миллиардов"},
		{"500000000000", "500 миллиардов"},
		{"600000000000", "600 миллиардов"},
		{"700000000000", "700 миллиардов"},
		{"800000000000", "800 миллиардов"},
		{"900000000000", "900 миллиардов"},
		{"99900000000", "99,9 миллиарда"},
		{"1000000000000", "1 триллион"},
		{"1100000000000", "1,1 триллиона"},
		{"1200000000000", "1,2 триллиона"},
		{"1300000000000", "1,3 триллиона"},
		{"1400000000000", "1,4 триллиона"},
		{"1500000000000", "1,5 триллиона"},
		{"1600000000000", "1,6 триллиона"},
		{"1700000000000", "1,7 триллиона"},
		{"1800000000000", "1,8 триллиона"},
		{"1900000000000", "1,9 триллиона"},
		{"2000000000000", "2 триллиона"},
		{"3000000000000", "3 триллиона"},
		{"4000000000000", "4 триллиона"},
		{"5000000000000", "5 триллионов"},
		{"6000000000000", "6 триллионов"},
		{"7000000000000", "7 триллионов"},
		{"8000000000000", "8 триллионов"},
		{"9000000000000", "9 триллионов"},
		{"10000000000000", "10 триллионов"},
		{"99900000000000", "99,9 триллиона"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.Russian)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeRuOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 тыс."},
		{"2000", "2 тыс."},
		{"3000", "3 тыс."},
		{"4000", "4 тыс."},
		{"5000", "5 тыс."},
		{"10000", "10 тыс."},
		{"11000", "11 тыс."},
		{"12000", "12 тыс."},
		{"13000", "13 тыс."},
		{"14000", "14 тыс."},
		{"15000", "15 тыс."},
		{"11100", "11,1 тыс."},
		{"12100", "12,1 тыс."},
		{"13100", "13,1 тыс."},
		{"14100", "14,1 тыс."},
		{"15100", "15,1 тыс."},
		{"12500", "12,5 тыс."},
		{"15000", "15 тыс."},
		{"15100", "15,1 тыс."},
		{"99500", "99,5 тыс."},
		{"1000000", "1 млн"},
		{"10100000", "10,1 млн"},
		{"1100000", "1,1 млн"},
		{"1200000", "1,2 млн"},
		{"1300000", "1,3 млн"},
		{"1400000", "1,4 млн"},
		{"1500000", "1,5 млн"},
		{"1600000", "1,6 млн"},
		{"1700000", "1,7 млн"},
		{"1800000", "1,8 млн"},
		{"1900000", "1,9 млн"},
		{"99900000", "99,9 млн"},
		{"2000000", "2 млн"},
		{"2300000", "2,3 млн"},
		{"1700000000", "1,7 млрд"},
		{"1000000000", "1 млрд"},
		{"1100000000", "1,1 млрд"},
		{"1200000000", "1,2 млрд"},
		{"1300000000", "1,3 млрд"},
		{"1400000000", "1,4 млрд"},
		{"1500000000", "1,5 млрд"},
		{"1600000000", "1,6 млрд"},
		{"1700000000", "1,7 млрд"},
		{"1800000000", "1,8 млрд"},
		{"1900000000", "1,9 млрд"},
		{"2000000000", "2 млрд"},
		{"3000000000", "3 млрд"},
		{"4000000000", "4 млрд"},
		{"5000000000", "5 млрд"},
		{"6000000000", "6 млрд"},
		{"7000000000", "7 млрд"},
		{"8000000000", "8 млрд"},
		{"9000000000", "9 млрд"},
		{"10000000000", "10 млрд"},
		{"11000000000", "11 млрд"},
		{"12000000000", "12 млрд"},
		{"13000000000", "13 млрд"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.Russian)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
