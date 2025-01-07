package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/pt"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[language.Tag]hc.Locale{
	language.Portuguese: locale.Data,
}

func TestHumanizePtOptionLong(t *testing.T) {
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
		{"1000000", "1 milhão"},
		{"10100000", "10,1 milhões"},
		{"1100000", "1,1 milhão"},
		{"1200000", "1,2 milhão"},
		{"1300000", "1,3 milhão"},
		{"1400000", "1,4 milhão"},
		{"1500000", "1,5 milhão"},
		{"1600000", "1,6 milhões"},
		{"1700000", "1,7 milhões"},
		{"1800000", "1,8 milhões"},
		{"1900000", "1,9 milhões"},
		{"99900000", "99,9 milhões"},
		{"99900000000", "99,9 bilhões"},
		{"99900000000000", "99,9 trilhões"},
		{"2000000", "2 milhões"},
		{"2300000", "2,3 milhões"},
		{"1700000000", "1,7 bilhões"},
		{"1000000000", "1 bilhão"},
		{"1100000000", "1,1 bilhão"},
		{"1200000000", "1,2 bilhão"},
		{"1300000000", "1,3 bilhão"},
		{"1400000000", "1,4 bilhão"},
		{"1500000000", "1,5 bilhão"},
		{"1600000000", "1,6 bilhões"},
		{"1700000000", "1,7 bilhões"},
		{"1800000000", "1,8 bilhões"},
		{"1900000000", "1,9 bilhões"},
		{"2000000000", "2 bilhões"},
		{"3000000000", "3 bilhões"},
		{"4000000000", "4 bilhões"},
		{"5000000000", "5 bilhões"},
		{"6000000000", "6 bilhões"},
		{"7000000000", "7 bilhões"},
		{"8000000000", "8 bilhões"},
		{"9000000000", "9 bilhões"},
		{"10000000000", "10 bilhões"},
		{"11000000000", "11 bilhões"},
		{"12000000000", "12 bilhões"},
		{"13000000000", "13 bilhões"},
		{"14000000000", "14 bilhões"},
		{"15000000000", "15 bilhões"},
		{"16000000000", "16 bilhões"},
		{"17000000000", "17 bilhões"},
		{"18000000000", "18 bilhões"},
		{"19000000000", "19 bilhões"},
		{"20000000000", "20 bilhões"},
		{"100000000000", "100 bilhões"},
		{"200000000000", "200 bilhões"},
		{"300000000000", "300 bilhões"},
		{"400000000000", "400 bilhões"},
		{"500000000000", "500 bilhões"},
		{"600000000000", "600 bilhões"},
		{"700000000000", "700 bilhões"},
		{"800000000000", "800 bilhões"},
		{"900000000000", "900 bilhões"},
		{"1000000000000", "1 trilhão"},
		{"1100000000000", "1,1 trilhão"},
		{"1200000000000", "1,2 trilhão"},
		{"1300000000000", "1,3 trilhão"},
		{"1400000000000", "1,4 trilhão"},
		{"1500000000000", "1,5 trilhão"},
		{"1600000000000", "1,6 trilhões"},
		{"1700000000000", "1,7 trilhões"},
		{"1800000000000", "1,8 trilhões"},
		{"1900000000000", "1,9 trilhões"},
		{"2000000000000", "2 trilhões"},
		{"3000000000000", "3 trilhões"},
		{"4000000000000", "4 trilhões"},
		{"5000000000000", "5 trilhões"},
		{"6000000000000", "6 trilhões"},
		{"7000000000000", "7 trilhões"},
		{"8000000000000", "8 trilhões"},
		{"9000000000000", "9 trilhões"},
		{"10000000000000", "10 trilhões"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Portuguese)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizePtOptionShort(t *testing.T) {
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
		{"1000000", "1 mi"},
		{"1100000", "1,1 mi"},
		{"1200000", "1,2 mi"},
		{"1300000", "1,3 mi"},
		{"1400000", "1,4 mi"},
		{"1500000", "1,5 mi"},
		{"1600000", "1,6 mi"},
		{"1700000", "1,7 mi"},
		{"1800000", "1,8 mi"},
		{"1900000", "1,9 mi"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Portuguese)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
