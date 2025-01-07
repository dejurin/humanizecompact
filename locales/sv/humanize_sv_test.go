package locale_test

import (
	"testing"

	"github.com/dejurin/humanize-cldr"
	locale "github.com/dejurin/humanize-cldr/locales/sv"
)

func fallback(number string) string {
	return number
}

func TestHumanizeSvOptionLong(t *testing.T) {
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
		{"1000", "1 tusen"},
		{"2000", "2 tusen"},
		{"3000", "3 tusen"},
		{"4000", "4 tusen"},
		{"5000", "5 tusen"},
		{"11000", "11 tusen"},
		{"12000", "12 tusen"},
		{"13000", "13 tusen"},
		{"14000", "14 tusen"},
		{"15000", "15 tusen"},
		{"11100", "11,1 tusen"},
		{"12100", "12,1 tusen"},
		{"13100", "13,1 tusen"},
		{"14100", "14,1 tusen"},
		{"15100", "15,1 tusen"},
		{"10000", "10 tusen"},
		{"12500", "12,5 tusen"},
		{"15000", "15 tusen"},
		{"15100", "15,1 tusen"},
		{"99500", "99,5 tusen"},
		{"1000000", "1 miljon"},
		{"10100000", "10,1 miljoner"},
		{"1100000", "1,1 miljoner"},
		{"1200000", "1,2 miljoner"},
		{"1300000", "1,3 miljoner"},
		{"1400000", "1,4 miljoner"},
		{"1500000", "1,5 miljoner"},
		{"1600000", "1,6 miljoner"},
		{"1700000", "1,7 miljoner"},
		{"1800000", "1,8 miljoner"},
		{"1900000", "1,9 miljoner"},
		{"99900000", "99,9 miljoner"},
		{"99900000000", "99,9 miljarder"},
		{"99900000000000", "99,9 biljoner"},
		{"2000000", "2 miljoner"},
		{"2300000", "2,3 miljoner"},
		{"1700000000", "1,7 miljarder"},
		{"1000000000", "1 miljard"},
		{"1100000000", "1,1 miljarder"},
		{"1200000000", "1,2 miljarder"},
		{"1300000000", "1,3 miljarder"},
		{"1400000000", "1,4 miljarder"},
		{"1500000000", "1,5 miljarder"},
		{"1600000000", "1,6 miljarder"},
		{"1700000000", "1,7 miljarder"},
		{"1800000000", "1,8 miljarder"},
		{"1900000000", "1,9 miljarder"},
		{"2000000000", "2 miljarder"},
		{"3000000000", "3 miljarder"},
		{"4000000000", "4 miljarder"},
		{"5000000000", "5 miljarder"},
		{"6000000000", "6 miljarder"},
		{"7000000000", "7 miljarder"},
		{"8000000000", "8 miljarder"},
		{"9000000000", "9 miljarder"},
		{"10000000000", "10 miljarder"},
		{"11000000000", "11 miljarder"},
		{"12000000000", "12 miljarder"},
		{"13000000000", "13 miljarder"},
		{"14000000000", "14 miljarder"},
		{"15000000000", "15 miljarder"},
		{"16000000000", "16 miljarder"},
		{"17000000000", "17 miljarder"},
		{"18000000000", "18 miljarder"},
		{"19000000000", "19 miljarder"},
		{"20000000000", "20 miljarder"},
		{"100000000000", "100 miljarder"},
		{"200000000000", "200 miljarder"},
		{"300000000000", "300 miljarder"},
		{"400000000000", "400 miljarder"},
		{"500000000000", "500 miljarder"},
		{"600000000000", "600 miljarder"},
		{"700000000000", "700 miljarder"},
		{"800000000000", "800 miljarder"},
		{"900000000000", "900 miljarder"},
		{"1000000000000", "1 biljon"},
		{"1100000000000", "1,1 biljoner"},
		{"1200000000000", "1,2 biljoner"},
		{"1300000000000", "1,3 biljoner"},
		{"1400000000000", "1,4 biljoner"},
		{"1500000000000", "1,5 biljoner"},
		{"1600000000000", "1,6 biljoner"},
		{"1700000000000", "1,7 biljoner"},
		{"1800000000000", "1,8 biljoner"},
		{"1900000000000", "1,9 biljoner"},
		{"2000000000000", "2 biljoner"},
		{"3000000000000", "3 biljoner"},
		{"4000000000000", "4 biljoner"},
		{"5000000000000", "5 biljoner"},
		{"6000000000000", "6 biljoner"},
		{"7000000000000", "7 biljoner"},
		{"8000000000000", "8 biljoner"},
		{"9000000000000", "9 biljoner"},
		{"10000000000000", "10 biljoner"},
	}

	h := humanize.New(locale.Data, humanize.Long, fallback)

	for _, tt := range tests {
		res, err := h.Humanize(tt.number)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeSvOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 tn"},
		{"2000", "2 tn"},
		{"3000", "3 tn"},
		{"4000", "4 tn"},
		{"5000", "5 tn"},
		{"11000", "11 tn"},
		{"12000", "12 tn"},
		{"13000", "13 tn"},
		{"14000", "14 tn"},
		{"15000", "15 tn"},
		{"1000000", "1 mn"},
		{"1100000", "1,1 mn"},
		{"1200000", "1,2 mn"},
		{"1300000", "1,3 mn"},
		{"1400000", "1,4 mn"},
		{"1500000", "1,5 mn"},
		{"1600000", "1,6 mn"},
		{"1700000", "1,7 mn"},
		{"1800000", "1,8 mn"},
		{"1900000", "1,9 mn"},
	}

	h := humanize.New(locale.Data, humanize.Short, fallback)

	for _, tt := range tests {
		res, err := h.Humanize(tt.number)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
