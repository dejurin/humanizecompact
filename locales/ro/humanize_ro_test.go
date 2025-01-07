package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/ro"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[string]hc.Locale{
	"ro": locale.Data,
}

func TestHumanizeRoOptionLong(t *testing.T) {
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
		{"1000", "1 mie"},
		{"2000", "2 mii"},
		{"3000", "3 mii"},
		{"4000", "4 mii"},
		{"5000", "5 mii"},
		{"11000", "11 mii"},
		{"12000", "12 mii"},
		{"13000", "13 mii"},
		{"14000", "14 mii"},
		{"15000", "15 mii"},
		{"11100", "11,1 mii"},
		{"12100", "12,1 mii"},
		{"13100", "13,1 mii"},
		{"14100", "14,1 mii"},
		{"15100", "15,1 mii"},
		{"10000", "10 mii"},
		{"12500", "12,5 mii"},
		{"15000", "15 mii"},
		{"15100", "15,1 mii"},
		{"99500", "99,5 mii"},
		{"1000000", "1 milion"},
		{"1100000", "1,1 milioane"},
		{"1200000", "1,2 milioane"},
		{"1300000", "1,3 milioane"},
		{"1400000", "1,4 milioane"},
		{"1500000", "1,5 milioane"},
		{"1600000", "1,6 milioane"},
		{"1700000", "1,7 milioane"},
		{"1800000", "1,8 milioane"},
		{"1900000", "1,9 milioane"},
		{"10100000", "10,1 milioane"},
		{"99900000", "99,9 milioane"},
		{"99900000000", "99,9 miliarde"},
		{"99900000000000", "99,9 trilioane"},
		{"2000000", "2 milioane"},
		{"2300000", "2,3 milioane"},
		{"1000000000", "1 miliard"},
		{"1100000000", "1,1 miliarde"},
		{"1200000000", "1,2 miliarde"},
		{"1300000000", "1,3 miliarde"},
		{"1400000000", "1,4 miliarde"},
		{"1500000000", "1,5 miliarde"},
		{"1600000000", "1,6 miliarde"},
		{"1700000000", "1,7 miliarde"},
		{"1800000000", "1,8 miliarde"},
		{"1900000000", "1,9 miliarde"},
		{"2000000000", "2 miliarde"},
		{"3000000000", "3 miliarde"},
		{"4000000000", "4 miliarde"},
		{"5000000000", "5 miliarde"},
		{"6000000000", "6 miliarde"},
		{"7000000000", "7 miliarde"},
		{"8000000000", "8 miliarde"},
		{"9000000000", "9 miliarde"},
		{"10000000000", "10 miliarde"},
		{"11000000000", "11 miliarde"},
		{"12000000000", "12 miliarde"},
		{"13000000000", "13 miliarde"},
		{"14000000000", "14 miliarde"},
		{"15000000000", "15 miliarde"},
		{"16000000000", "16 miliarde"},
		{"17000000000", "17 miliarde"},
		{"18000000000", "18 miliarde"},
		{"19000000000", "19 miliarde"},
		{"20000000000", "20 de miliarde"},
		{"100000000000", "100 de miliarde"},
		{"200000000000", "200 de miliarde"},
		{"300000000000", "300 de miliarde"},
		{"400000000000", "400 de miliarde"},
		{"500000000000", "500 de miliarde"},
		{"600000000000", "600 de miliarde"},
		{"700000000000", "700 de miliarde"},
		{"800000000000", "800 de miliarde"},
		{"900000000000", "900 de miliarde"},
		{"1000000000000", "1 trilion"},
		{"1100000000000", "1,1 trilioane"},
		{"1200000000000", "1,2 trilioane"},
		{"1300000000000", "1,3 trilioane"},
		{"1400000000000", "1,4 trilioane"},
		{"1500000000000", "1,5 trilioane"},
		{"1600000000000", "1,6 trilioane"},
		{"1700000000000", "1,7 trilioane"},
		{"1800000000000", "1,8 trilioane"},
		{"1900000000000", "1,9 trilioane"},
		{"2000000000000", "2 trilioane"},
		{"3000000000000", "3 trilioane"},
		{"4000000000000", "4 trilioane"},
		{"5000000000000", "5 trilioane"},
		{"6000000000000", "6 trilioane"},
		{"7000000000000", "7 trilioane"},
		{"8000000000000", "8 trilioane"},
		{"9000000000000", "9 trilioane"},
		{"10000000000000", "10 trilioane"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, err := h.Humanize(tt.number, language.Romanian)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeRoOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 K"},
		{"10000", "10 K"},
		{"100000", "100 K"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, err := h.Humanize(tt.number, language.Romanian)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
