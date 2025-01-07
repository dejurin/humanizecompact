package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/es"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[string]hc.Locale{
	"es": locale.Data,
}

func TestHumanizeEsOptionLong(t *testing.T) {
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
		{"1000", "1 mil"},
		{"2000", "2 mil"},
		{"3000", "3 mil"},
		{"4000", "4 mil"},
		{"5000", "5 mil"},
		{"11000", "11 mil"},
		{"12000", "12 mil"},
		{"13000", "13 mil"},
		{"14000", "14 mil"},
		{"15000", "15 mil"},
		{"11100", "11,1 mil"},
		{"12100", "12,1 mil"},
		{"13100", "13,1 mil"},
		{"14100", "14,1 mil"},
		{"15100", "15,1 mil"},
		{"10000", "10 mil"},
		{"12500", "12,5 mil"},
		{"15000", "15 mil"},
		{"15100", "15,1 mil"},
		{"99500", "99,5 mil"},
		{"1000000", "1 millón"},
		{"1100000", "1,1 millones"},
		{"1200000", "1,2 millones"},
		{"1300000", "1,3 millones"},
		{"1400000", "1,4 millones"},
		{"1500000", "1,5 millones"},
		{"1600000", "1,6 millones"},
		{"1700000", "1,7 millones"},
		{"1800000", "1,8 millones"},
		{"1900000", "1,9 millones"},
		{"10100000", "10,1 millones"},
		{"99900000", "99,9 millones"},
		{"99900000000", "99,9 mil millones"},
		{"99900000000000", "99,9 billones"},
		{"2000000", "2 millones"},
		{"2300000", "2,3 millones"},
		{"1000000000", "1 mil millones"},
		{"1100000000", "1,1 mil millones"},
		{"1200000000", "1,2 mil millones"},
		{"1300000000", "1,3 mil millones"},
		{"1400000000", "1,4 mil millones"},
		{"1500000000", "1,5 mil millones"},
		{"1600000000", "1,6 mil millones"},
		{"1700000000", "1,7 mil millones"},
		{"1800000000", "1,8 mil millones"},
		{"1900000000", "1,9 mil millones"},
		{"2000000000", "2 mil millones"},
		{"3000000000", "3 mil millones"},
		{"4000000000", "4 mil millones"},
		{"5000000000", "5 mil millones"},
		{"6000000000", "6 mil millones"},
		{"7000000000", "7 mil millones"},
		{"8000000000", "8 mil millones"},
		{"9000000000", "9 mil millones"},
		{"10000000000", "10 mil millones"},
		{"11000000000", "11 mil millones"},
		{"12000000000", "12 mil millones"},
		{"13000000000", "13 mil millones"},
		{"14000000000", "14 mil millones"},
		{"15000000000", "15 mil millones"},
		{"16000000000", "16 mil millones"},
		{"17000000000", "17 mil millones"},
		{"18000000000", "18 mil millones"},
		{"19000000000", "19 mil millones"},
		{"20000000000", "20 mil millones"},
		{"100000000000", "100 mil millones"},
		{"200000000000", "200 mil millones"},
		{"300000000000", "300 mil millones"},
		{"400000000000", "400 mil millones"},
		{"500000000000", "500 mil millones"},
		{"600000000000", "600 mil millones"},
		{"700000000000", "700 mil millones"},
		{"800000000000", "800 mil millones"},
		{"900000000000", "900 mil millones"},
		{"1000000000000", "1 billón"},
		{"1100000000000", "1,1 billones"},
		{"1200000000000", "1,2 billones"},
		{"1300000000000", "1,3 billones"},
		{"1400000000000", "1,4 billones"},
		{"1500000000000", "1,5 billones"},
		{"1600000000000", "1,6 billones"},
		{"1700000000000", "1,7 billones"},
		{"1800000000000", "1,8 billones"},
		{"1900000000000", "1,9 billones"},
		{"2000000000000", "2 billones"},
		{"3000000000000", "3 billones"},
		{"4000000000000", "4 billones"},
		{"5000000000000", "5 billones"},
		{"6000000000000", "6 billones"},
		{"7000000000000", "7 billones"},
		{"8000000000000", "8 billones"},
		{"9000000000000", "9 billones"},
		{"10000000000000", "10 billones"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.Spanish)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeEsOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 mil"},
		{"2000", "2 mil"},
		{"3000", "3 mil"},
		{"4000", "4 mil"},
		{"5000", "5 mil"},
		{"11000", "11 mil"},
		{"12000", "12 mil"},
		{"13000", "13 mil"},
		{"14000", "14 mil"},
		{"15000", "15 mil"},
		{"1000000", "1 M"},
		{"1100000", "1,1 M"},
		{"1200000", "1,2 M"},
		{"1300000", "1,3 M"},
		{"1400000", "1,4 M"},
		{"1500000", "1,5 M"},
		{"10000000000", "1 mil M"},
		{"11000000000", "1,1 mil M"},
		{"12000000000", "1,2 mil M"},
		{"13000000000", "1,3 mil M"},
		{"14000000000", "1,4 mil M"},
		{"15000000000", "1,5 mil M"},
		{"1000000000000", "1 B"},
		{"1100000000000", "1,1 B"},
		{"1200000000000", "1,2 B"},
		{"1300000000000", "1,3 B"},
		{"1400000000000", "1,4 B"},
		{"1500000000000", "1,5 B"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.Spanish)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
