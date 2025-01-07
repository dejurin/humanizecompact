package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"

	locale "github.com/dejurin/humanizecompact/locales/id"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[language.Tag]hc.Locale{
	language.Indonesian: locale.Data,
}

func TestHumanizeIdOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1", "1"},                         // fallback
		{"9999", "9999"},                   // fallback
		{"100100", "100100"},               // fallback
		{"100001", "100001"},               // fallback
		{"1230000", "1230000"},             // fallback
		{"1234000", "1234000"},             // fallback
		{"1234200", "1234200"},             // fallback
		{"1234500", "1234500"},             // fallback
		{"2340000", "2340000"},             // fallback
		{"2345678", "2345678"},             // fallback
		{"1000000.1", "1000000.1"},         // fallback
		{"1234000000", "1234000000"},       // fallback
		{"1234000001", "1234000001"},       // fallback
		{"1234500000", "1234500000"},       // fallback
		{"1910000000000", "1910000000000"}, // fallback
		{"100100000", "100100000"},         // fallback
		{"1000", "1 ribu"},
		{"2000", "2 ribu"},
		{"3000", "3 ribu"},
		{"4000", "4 ribu"},
		{"5000", "5 ribu"},
		{"11000", "11 ribu"},
		{"12000", "12 ribu"},
		{"13000", "13 ribu"},
		{"14000", "14 ribu"},
		{"15000", "15 ribu"},
		{"11100", "11,1 ribu"},
		{"12100", "12,1 ribu"},
		{"13100", "13,1 ribu"},
		{"14100", "14,1 ribu"},
		{"15100", "15,1 ribu"},
		{"10000", "10 ribu"},
		{"12500", "12,5 ribu"},
		{"15000", "15 ribu"},
		{"15100", "15,1 ribu"},
		{"99500", "99,5 ribu"},
		{"1000000", "1 juta"},
		{"10100000", "10,1 juta"},
		{"1100000", "1,1 juta"},
		{"1200000", "1,2 juta"},
		{"1300000", "1,3 juta"},
		{"1400000", "1,4 juta"},
		{"1500000", "1,5 juta"},
		{"1600000", "1,6 juta"},
		{"1700000", "1,7 juta"},
		{"1800000", "1,8 juta"},
		{"1900000", "1,9 juta"},
		{"99900000", "99,9 juta"},
		{"99900000000", "99,9 miliar"},
		{"99900000000000", "99,9 triliun"},
		{"2000000", "2 juta"},
		{"2300000", "2,3 juta"},
		{"1700000000", "1,7 miliar"},
		{"1000000000", "1 miliar"},
		{"1100000000", "1,1 miliar"},
		{"1200000000", "1,2 miliar"},
		{"1300000000", "1,3 miliar"},
		{"1400000000", "1,4 miliar"},
		{"1500000000", "1,5 miliar"},
		{"1600000000", "1,6 miliar"},
		{"1700000000", "1,7 miliar"},
		{"1800000000", "1,8 miliar"},
		{"1900000000", "1,9 miliar"},
		{"2000000000", "2 miliar"},
		{"3000000000", "3 miliar"},
		{"4000000000", "4 miliar"},
		{"5000000000", "5 miliar"},
		{"6000000000", "6 miliar"},
		{"7000000000", "7 miliar"},
		{"8000000000", "8 miliar"},
		{"9000000000", "9 miliar"},
		{"10000000000", "10 miliar"},
		{"11000000000", "11 miliar"},
		{"12000000000", "12 miliar"},
		{"13000000000", "13 miliar"},
		{"14000000000", "14 miliar"},
		{"15000000000", "15 miliar"},
		{"16000000000", "16 miliar"},
		{"17000000000", "17 miliar"},
		{"18000000000", "18 miliar"},
		{"19000000000", "19 miliar"},
		{"20000000000", "20 miliar"},
		{"100000000000", "100 miliar"},
		{"200000000000", "200 miliar"},
		{"300000000000", "300 miliar"},
		{"400000000000", "400 miliar"},
		{"500000000000", "500 miliar"},
		{"600000000000", "600 miliar"},
		{"700000000000", "700 miliar"},
		{"800000000000", "800 miliar"},
		{"900000000000", "900 miliar"},
		{"1000000000000", "1 triliun"},
		{"1100000000000", "1,1 triliun"},
		{"1200000000000", "1,2 triliun"},
		{"1300000000000", "1,3 triliun"},
		{"1400000000000", "1,4 triliun"},
		{"1500000000000", "1,5 triliun"},
		{"1600000000000", "1,6 triliun"},
		{"1700000000000", "1,7 triliun"},
		{"1800000000000", "1,8 triliun"},
		{"1900000000000", "1,9 triliun"},
		{"2000000000000", "2 triliun"},
		{"3000000000000", "3 triliun"},
		{"4000000000000", "4 triliun"},
		{"5000000000000", "5 triliun"},
		{"6000000000000", "6 triliun"},
		{"7000000000000", "7 triliun"},
		{"8000000000000", "8 triliun"},
		{"9000000000000", "9 triliun"},
		{"10000000000000", "10 triliun"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.Indonesian)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeIdOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 rb"},
		{"2000", "2 rb"},
		{"3000", "3 rb"},
		{"4000", "4 rb"},
		{"5000", "5 rb"},
		{"11000", "11 rb"},
		{"12000", "12 rb"},
		{"13000", "13 rb"},
		{"14000", "14 rb"},
		{"15000", "15 rb"},
		{"11100", "11,1 rb"},
		{"12100", "12,1 rb"},
		{"13100", "13,1 rb"},
		{"14100", "14,1 rb"},
		{"15100", "15,1 rb"},
		{"10000", "10 rb"},
		{"12500", "12,5 rb"},
		{"15000", "15 rb"},
		{"15100", "15,1 rb"},
		{"99500", "99,5 rb"},
		{"1000000", "1 jt"},
		{"10100000", "10,1 jt"},
		{"99900000", "99,9 jt"},
		{"100000000", "100 jt"},
		{"101000000", "101 jt"},
		{"999000000", "999 jt"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.Indonesian)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
