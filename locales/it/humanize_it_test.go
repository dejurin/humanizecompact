package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"

	locale "github.com/dejurin/humanizecompact/locales/it"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[string]hc.Locale{
	"it": locale.Data,
}

func TestHumanizeItOptionLong(t *testing.T) {
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
		{"1000", "mille"},
		{"2000", "2 mila"},
		{"3000", "3 mila"},
		{"4000", "4 mila"},
		{"5000", "5 mila"},
		{"11000", "11 mila"},
		{"12000", "12 mila"},
		{"13000", "13 mila"},
		{"14000", "14 mila"},
		{"15000", "15 mila"},
		{"11100", "11,1 mila"},
		{"12100", "12,1 mila"},
		{"13100", "13,1 mila"},
		{"14100", "14,1 mila"},
		{"15100", "15,1 mila"},
		{"10000", "10 mila"},
		{"12500", "12,5 mila"},
		{"15000", "15 mila"},
		{"15100", "15,1 mila"},
		{"99500", "99,5 mila"},
		{"1000000", "1 milione"},
		{"10100000", "10,1 milioni"},
		{"1100000", "1,1 milioni"},
		{"1200000", "1,2 milioni"},
		{"1300000", "1,3 milioni"},
		{"1400000", "1,4 milioni"},
		{"1500000", "1,5 milioni"},
		{"1600000", "1,6 milioni"},
		{"1700000", "1,7 milioni"},
		{"1800000", "1,8 milioni"},
		{"1900000", "1,9 milioni"},
		{"99900000", "99,9 milioni"},
		{"99900000000", "99,9 miliardi"},
		{"99900000000000", "99,9 mila miliardi"},
		{"2000000", "2 milioni"},
		{"2300000", "2,3 milioni"},
		{"1700000000", "1,7 miliardi"},
		{"1000000000", "1 miliardo"},
		{"1100000000", "1,1 miliardi"},
		{"1200000000", "1,2 miliardi"},
		{"1300000000", "1,3 miliardi"},
		{"1400000000", "1,4 miliardi"},
		{"1500000000", "1,5 miliardi"},
		{"1600000000", "1,6 miliardi"},
		{"1700000000", "1,7 miliardi"},
		{"1800000000", "1,8 miliardi"},
		{"1900000000", "1,9 miliardi"},
		{"2000000000", "2 miliardi"},
		{"3000000000", "3 miliardi"},
		{"4000000000", "4 miliardi"},
		{"5000000000", "5 miliardi"},
		{"6000000000", "6 miliardi"},
		{"7000000000", "7 miliardi"},
		{"8000000000", "8 miliardi"},
		{"9000000000", "9 miliardi"},
		{"10000000000", "10 miliardi"},
		{"11000000000", "11 miliardi"},
		{"12000000000", "12 miliardi"},
		{"13000000000", "13 miliardi"},
		{"14000000000", "14 miliardi"},
		{"15000000000", "15 miliardi"},
		{"16000000000", "16 miliardi"},
		{"17000000000", "17 miliardi"},
		{"18000000000", "18 miliardi"},
		{"19000000000", "19 miliardi"},
		{"20000000000", "20 miliardi"},
		{"100000000000", "100 miliardi"},
		{"200000000000", "200 miliardi"},
		{"300000000000", "300 miliardi"},
		{"400000000000", "400 miliardi"},
		{"500000000000", "500 miliardi"},
		{"600000000000", "600 miliardi"},
		{"700000000000", "700 miliardi"},
		{"800000000000", "800 miliardi"},
		{"900000000000", "900 miliardi"},
		{"1000000000000", "mille miliardi"},
		{"1100000000000", "1,1 mila miliardi"},
		{"1200000000000", "1,2 mila miliardi"},
		{"1300000000000", "1,3 mila miliardi"},
		{"1400000000000", "1,4 mila miliardi"},
		{"1500000000000", "1,5 mila miliardi"},
		{"1600000000000", "1,6 mila miliardi"},
		{"1700000000000", "1,7 mila miliardi"},
		{"1800000000000", "1,8 mila miliardi"},
		{"1900000000000", "1,9 mila miliardi"},
		{"2000000000000", "2 mila miliardi"},
		{"3000000000000", "3 mila miliardi"},
		{"4000000000000", "4 mila miliardi"},
		{"5000000000000", "5 mila miliardi"},
		{"6000000000000", "6 mila miliardi"},
		{"7000000000000", "7 mila miliardi"},
		{"8000000000000", "8 mila miliardi"},
		{"9000000000000", "9 mila miliardi"},
		{"10000000000000", "10 mila miliardi"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.Italian)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeItOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1000"},
		{"2000", "2000"},
		{"3000", "3000"},
		{"4000", "4000"},
		{"5000", "5000"},
		{"11000", "11000"},
		{"12000", "12000"},
		{"13000", "13000"},
		{"14000", "14000"},
		{"15000", "15000"},
		{"1000000", "1 Mln"},
		{"10100000", "10,1 Mln"},
		{"1100000", "1,1 Mln"},
		{"1200000", "1,2 Mln"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.Italian)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
