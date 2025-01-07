package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"

	locale "github.com/dejurin/humanizecompact/locales/de"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[language.Tag]hc.Locale{
	language.German: locale.Data,
}

func TestHumanizeDeOptionLong(t *testing.T) {
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
		{"1000", "1 Tausend"},
		{"2000", "2 Tausend"},
		{"3000", "3 Tausend"},
		{"4000", "4 Tausend"},
		{"5000", "5 Tausend"},
		{"11000", "11 Tausend"},
		{"12000", "12 Tausend"},
		{"13000", "13 Tausend"},
		{"14000", "14 Tausend"},
		{"15000", "15 Tausend"},
		{"11100", "11,1 Tausend"},
		{"12100", "12,1 Tausend"},
		{"13100", "13,1 Tausend"},
		{"14100", "14,1 Tausend"},
		{"15100", "15,1 Tausend"},
		{"10000", "10 Tausend"},
		{"12500", "12,5 Tausend"},
		{"15000", "15 Tausend"},
		{"15100", "15,1 Tausend"},
		{"99500", "99,5 Tausend"},
		{"1000000", "1 Million"},
		{"10100000", "10,1 Millionen"},
		{"1100000", "1,1 Millionen"},
		{"1200000", "1,2 Millionen"},
		{"1300000", "1,3 Millionen"},
		{"1400000", "1,4 Millionen"},
		{"1500000", "1,5 Millionen"},
		{"1600000", "1,6 Millionen"},
		{"1700000", "1,7 Millionen"},
		{"1800000", "1,8 Millionen"},
		{"1900000", "1,9 Millionen"},
		{"99900000", "99,9 Millionen"},
		{"99900000000", "99,9 Milliarden"},
		{"99900000000000", "99,9 Billionen"},
		{"2000000", "2 Millionen"},
		{"2300000", "2,3 Millionen"},
		{"1700000000", "1,7 Milliarden"},
		{"1000000000", "1 Milliarde"},
		{"1100000000", "1,1 Milliarden"},
		{"1200000000", "1,2 Milliarden"},
		{"1300000000", "1,3 Milliarden"},
		{"1400000000", "1,4 Milliarden"},
		{"1500000000", "1,5 Milliarden"},
		{"1600000000", "1,6 Milliarden"},
		{"1700000000", "1,7 Milliarden"},
		{"1800000000", "1,8 Milliarden"},
		{"1900000000", "1,9 Milliarden"},
		{"2000000000", "2 Milliarden"},
		{"3000000000", "3 Milliarden"},
		{"4000000000", "4 Milliarden"},
		{"5000000000", "5 Milliarden"},
		{"6000000000", "6 Milliarden"},
		{"7000000000", "7 Milliarden"},
		{"8000000000", "8 Milliarden"},
		{"9000000000", "9 Milliarden"},
		{"10000000000", "10 Milliarden"},
		{"11000000000", "11 Milliarden"},
		{"12000000000", "12 Milliarden"},
		{"13000000000", "13 Milliarden"},
		{"14000000000", "14 Milliarden"},
		{"15000000000", "15 Milliarden"},
		{"16000000000", "16 Milliarden"},
		{"17000000000", "17 Milliarden"},
		{"18000000000", "18 Milliarden"},
		{"19000000000", "19 Milliarden"},
		{"20000000000", "20 Milliarden"},
		{"100000000000", "100 Milliarden"},
		{"200000000000", "200 Milliarden"},
		{"300000000000", "300 Milliarden"},
		{"400000000000", "400 Milliarden"},
		{"500000000000", "500 Milliarden"},
		{"600000000000", "600 Milliarden"},
		{"700000000000", "700 Milliarden"},
		{"800000000000", "800 Milliarden"},
		{"900000000000", "900 Milliarden"},
		{"1000000000000", "1 Billion"},
		{"1100000000000", "1,1 Billionen"},
		{"1200000000000", "1,2 Billionen"},
		{"1300000000000", "1,3 Billionen"},
		{"1400000000000", "1,4 Billionen"},
		{"1500000000000", "1,5 Billionen"},
		{"1600000000000", "1,6 Billionen"},
		{"1700000000000", "1,7 Billionen"},
		{"1800000000000", "1,8 Billionen"},
		{"1900000000000", "1,9 Billionen"},
		{"2000000000000", "2 Billionen"},
		{"3000000000000", "3 Billionen"},
		{"4000000000000", "4 Billionen"},
		{"5000000000000", "5 Billionen"},
		{"6000000000000", "6 Billionen"},
		{"7000000000000", "7 Billionen"},
		{"8000000000000", "8 Billionen"},
		{"9000000000000", "9 Billionen"},
		{"10000000000000", "10 Billionen"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.German)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeDeOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1000"},
		{"2000", "2000"},
		{"3000", "3000"},
		{"4000", "4000"},
		{"10000", "10000"},
		{"20000", "20000"},
		{"30000", "30000"},
		{"40000", "40000"},
		{"50000", "50000"},
		{"100000", "100000"},
		{"200000", "200000"},
		{"300000", "300000"},
		{"400000", "400000"},
		{"500000", "500000"},
		{"600000", "600000"},
		{"700000", "700000"},
		{"800000", "800000"},
		{"900000", "900000"},
		{"1000000", "1 Mio."},
		{"1100000", "1,1 Mio."},
		{"1200000", "1,2 Mio."},
		{"1300000", "1,3 Mio."},
		{"1400000", "1,4 Mio."},
		{"1500000", "1,5 Mio."},
		{"1600000", "1,6 Mio."},
		{"1700000", "1,7 Mio."},
		{"1800000", "1,8 Mio."},
		{"1900000", "1,9 Mio."},
		{"2000000", "2 Mio."},
		{"10000000", "10 Mio."},
		{"100000000", "100 Mio."},
		{"200000000", "200 Mio."},
		{"300000000", "300 Mio."},
		{"400000000", "400 Mio."},
		{"500000000", "500 Mio."},
		{"600000000", "600 Mio."},
		{"700000000", "700 Mio."},
		{"800000000", "800 Mio."},
		{"900000000", "900 Mio."},
		{"1000000000", "1 Mrd."},
		{"1100000000", "1,1 Mrd."},
		{"1200000000", "1,2 Mrd."},
		{"1300000000", "1,3 Mrd."},
		{"1400000000", "1,4 Mrd."},
		{"1500000000", "1,5 Mrd."},
		{"1600000000", "1,6 Mrd."},
		{"1700000000", "1,7 Mrd."},
		{"1800000000", "1,8 Mrd."},
		{"1900000000", "1,9 Mrd."},
		{"1000000000000", "1 Bio."},
		{"1100000000000", "1,1 Bio."},
		{"1200000000000", "1,2 Bio."},
		{"1300000000000", "1,3 Bio."},
		{"1400000000000", "1,4 Bio."},
		{"1500000000000", "1,5 Bio."},
		{"2000000000000", "2 Bio."},
		{"3000000000000", "3 Bio."},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.German)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
