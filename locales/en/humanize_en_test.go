package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/en"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[string]hc.Locale{
	"en": locale.Data,
}

func TestHumanizeEnOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"-1000", "-1000"},                 // fallback
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
		{"1000", "1 thousand"},
		{"2000", "2 thousand"},
		{"3000", "3 thousand"},
		{"4000", "4 thousand"},
		{"5000", "5 thousand"},
		{"11000", "11 thousand"},
		{"12000", "12 thousand"},
		{"13000", "13 thousand"},
		{"14000", "14 thousand"},
		{"15000", "15 thousand"},
		{"11100", "11.1 thousand"},
		{"12100", "12.1 thousand"},
		{"13100", "13.1 thousand"},
		{"14100", "14.1 thousand"},
		{"15100", "15.1 thousand"},
		{"10000", "10 thousand"},
		{"12500", "12.5 thousand"},
		{"15000", "15 thousand"},
		{"15100", "15.1 thousand"},
		{"99500", "99.5 thousand"},
		{"1000000", "1 million"},
		{"10100000", "10.1 million"},
		{"1100000", "1.1 million"},
		{"1200000", "1.2 million"},
		{"1300000", "1.3 million"},
		{"1400000", "1.4 million"},
		{"1500000", "1.5 million"},
		{"1600000", "1.6 million"},
		{"1700000", "1.7 million"},
		{"1800000", "1.8 million"},
		{"1900000", "1.9 million"},
		{"99900000", "99.9 million"},
		{"99900000000", "99.9 billion"},
		{"99900000000000", "99.9 trillion"},
		{"2000000", "2 million"},
		{"2300000", "2.3 million"},
		{"1700000000", "1.7 billion"},
		{"1000000000", "1 billion"},
		{"1100000000", "1.1 billion"},
		{"1200000000", "1.2 billion"},
		{"1300000000", "1.3 billion"},
		{"1400000000", "1.4 billion"},
		{"1500000000", "1.5 billion"},
		{"1600000000", "1.6 billion"},
		{"1700000000", "1.7 billion"},
		{"1800000000", "1.8 billion"},
		{"1900000000", "1.9 billion"},
		{"2000000000", "2 billion"},
		{"3000000000", "3 billion"},
		{"4000000000", "4 billion"},
		{"5000000000", "5 billion"},
		{"6000000000", "6 billion"},
		{"7000000000", "7 billion"},
		{"8000000000", "8 billion"},
		{"9000000000", "9 billion"},
		{"10000000000", "10 billion"},
		{"11000000000", "11 billion"},
		{"12000000000", "12 billion"},
		{"13000000000", "13 billion"},
		{"14000000000", "14 billion"},
		{"15000000000", "15 billion"},
		{"16000000000", "16 billion"},
		{"17000000000", "17 billion"},
		{"18000000000", "18 billion"},
		{"19000000000", "19 billion"},
		{"20000000000", "20 billion"},
		{"100000000000", "100 billion"},
		{"200000000000", "200 billion"},
		{"300000000000", "300 billion"},
		{"400000000000", "400 billion"},
		{"500000000000", "500 billion"},
		{"600000000000", "600 billion"},
		{"700000000000", "700 billion"},
		{"800000000000", "800 billion"},
		{"900000000000", "900 billion"},
		{"1000000000000", "1 trillion"},
		{"1100000000000", "1.1 trillion"},
		{"1200000000000", "1.2 trillion"},
		{"1300000000000", "1.3 trillion"},
		{"1400000000000", "1.4 trillion"},
		{"1500000000000", "1.5 trillion"},
		{"1600000000000", "1.6 trillion"},
		{"1700000000000", "1.7 trillion"},
		{"1800000000000", "1.8 trillion"},
		{"1900000000000", "1.9 trillion"},
		{"2000000000000", "2 trillion"},
		{"3000000000000", "3 trillion"},
		{"4000000000000", "4 trillion"},
		{"5000000000000", "5 trillion"},
		{"6000000000000", "6 trillion"},
		{"7000000000000", "7 trillion"},
		{"8000000000000", "8 trillion"},
		{"9000000000000", "9 trillion"},
		{"10000000000000", "10 trillion"},
		{"99999999999999999", "99999999999999999"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, err := h.Humanize(tt.number, language.English)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeEnOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
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
		res, err := h.Humanize(tt.number, language.English)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
