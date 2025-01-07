package locale_test

import (
	"testing"

	"github.com/dejurin/humanize-cldr"
	locale "github.com/dejurin/humanize-cldr/locales/tr"
)

func fallback(number string) string {
	return number
}

func TestHumanizeTrOptionLong(t *testing.T) {
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
		{"1000", "1 bin"},
		{"2000", "2 bin"},
		{"3000", "3 bin"},
		{"4000", "4 bin"},
		{"5000", "5 bin"},
		{"11000", "11 bin"},
		{"12000", "12 bin"},
		{"13000", "13 bin"},
		{"14000", "14 bin"},
		{"15000", "15 bin"},
		{"11100", "11,1 bin"},
		{"12100", "12,1 bin"},
		{"13100", "13,1 bin"},
		{"14100", "14,1 bin"},
		{"15100", "15,1 bin"},
		{"10000", "10 bin"},
		{"12500", "12,5 bin"},
		{"15000", "15 bin"},
		{"15100", "15,1 bin"},
		{"99500", "99,5 bin"},
		{"1000000", "1 milyon"},
		{"10100000", "10,1 milyon"},
		{"1100000", "1,1 milyon"},
		{"1200000", "1,2 milyon"},
		{"1300000", "1,3 milyon"},
		{"1400000", "1,4 milyon"},
		{"1500000", "1,5 milyon"},
		{"1600000", "1,6 milyon"},
		{"1700000", "1,7 milyon"},
		{"1800000", "1,8 milyon"},
		{"1900000", "1,9 milyon"},
		{"99900000", "99,9 milyon"},
		{"99900000000", "99,9 milyar"},
		{"99900000000000", "99,9 trilyon"},
		{"2000000", "2 milyon"},
		{"2300000", "2,3 milyon"},
		{"1700000000", "1,7 milyar"},
		{"1000000000", "1 milyar"},
		{"1100000000", "1,1 milyar"},
		{"1200000000", "1,2 milyar"},
		{"1300000000", "1,3 milyar"},
		{"1400000000", "1,4 milyar"},
		{"1500000000", "1,5 milyar"},
		{"1600000000", "1,6 milyar"},
		{"1700000000", "1,7 milyar"},
		{"1800000000", "1,8 milyar"},
		{"1900000000", "1,9 milyar"},
		{"2000000000", "2 milyar"},
		{"3000000000", "3 milyar"},
		{"4000000000", "4 milyar"},
		{"5000000000", "5 milyar"},
		{"6000000000", "6 milyar"},
		{"7000000000", "7 milyar"},
		{"8000000000", "8 milyar"},
		{"9000000000", "9 milyar"},
		{"10000000000", "10 milyar"},
		{"11000000000", "11 milyar"},
		{"12000000000", "12 milyar"},
		{"13000000000", "13 milyar"},
		{"14000000000", "14 milyar"},
		{"15000000000", "15 milyar"},
		{"16000000000", "16 milyar"},
		{"17000000000", "17 milyar"},
		{"18000000000", "18 milyar"},
		{"19000000000", "19 milyar"},
		{"20000000000", "20 milyar"},
		{"100000000000", "100 milyar"},
		{"200000000000", "200 milyar"},
		{"300000000000", "300 milyar"},
		{"400000000000", "400 milyar"},
		{"500000000000", "500 milyar"},
		{"600000000000", "600 milyar"},
		{"700000000000", "700 milyar"},
		{"800000000000", "800 milyar"},
		{"900000000000", "900 milyar"},
		{"1000000000000", "1 trilyon"},
		{"1100000000000", "1,1 trilyon"},
		{"1200000000000", "1,2 trilyon"},
		{"1300000000000", "1,3 trilyon"},
		{"1400000000000", "1,4 trilyon"},
		{"1500000000000", "1,5 trilyon"},
		{"1600000000000", "1,6 trilyon"},
		{"1700000000000", "1,7 trilyon"},
		{"1800000000000", "1,8 trilyon"},
		{"1900000000000", "1,9 trilyon"},
		{"2000000000000", "2 trilyon"},
		{"3000000000000", "3 trilyon"},
		{"4000000000000", "4 trilyon"},
		{"5000000000000", "5 trilyon"},
		{"6000000000000", "6 trilyon"},
		{"7000000000000", "7 trilyon"},
		{"8000000000000", "8 trilyon"},
		{"9000000000000", "9 trilyon"},
		{"10000000000000", "10 trilyon"},
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

func TestHumanizeTrOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 B"},
		{"10000", "10 B"},
		{"100000", "100 B"},
		{"1000000", "1 Mn"},
		{"1100000", "1,1 Mn"},
		{"1200000", "1,2 Mn"},
		{"1300000", "1,3 Mn"},
		{"1400000", "1,4 Mn"},
		{"1500000", "1,5 Mn"},
		{"10000000", "10 Mn"},
		{"100000000", "100 Mn"},
		{"1000000000", "1 Mr"},
		{"1100000000", "1,1 Mr"},
		{"1200000000", "1,2 Mr"},
		{"1300000000", "1,3 Mr"},
		{"1400000000", "1,4 Mr"},
		{"1500000000", "1,5 Mr"},
		{"10000000000", "10 Mr"},
		{"100000000000", "100 Mr"},
		{"1000000000000", "1 Tn"},
		{"1100000000000", "1,1 Tn"},
		{"1200000000000", "1,2 Tn"},
		{"1300000000000", "1,3 Tn"},
		{"1400000000000", "1,4 Tn"},
		{"1500000000000", "1,5 Tn"},
		{"10000000000000", "10 Tn"},
		{"100000000000000", "100 Tn"},
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
