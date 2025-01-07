package locale_test

import (
	"testing"

	"github.com/dejurin/humanize-cldr"
	locale "github.com/dejurin/humanize-cldr/locales/hu"
)

func fallback(number string) string {
	return number
}

func TestHumanizeHuOptionLong(t *testing.T) {
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
		{"1000", "1 ezer"},
		{"2000", "2 ezer"},
		{"3000", "3 ezer"},
		{"4000", "4 ezer"},
		{"5000", "5 ezer"},
		{"11000", "11 ezer"},
		{"12000", "12 ezer"},
		{"13000", "13 ezer"},
		{"14000", "14 ezer"},
		{"15000", "15 ezer"},
		{"11100", "11,1 ezer"},
		{"12100", "12,1 ezer"},
		{"13100", "13,1 ezer"},
		{"14100", "14,1 ezer"},
		{"15100", "15,1 ezer"},
		{"10000", "10 ezer"},
		{"12500", "12,5 ezer"},
		{"15000", "15 ezer"},
		{"15100", "15,1 ezer"},
		{"99500", "99,5 ezer"},
		{"1000000", "1 millió"},
		{"10100000", "10,1 millió"},
		{"1100000", "1,1 millió"},
		{"1200000", "1,2 millió"},
		{"1300000", "1,3 millió"},
		{"1400000", "1,4 millió"},
		{"1500000", "1,5 millió"},
		{"1600000", "1,6 millió"},
		{"1700000", "1,7 millió"},
		{"1800000", "1,8 millió"},
		{"1900000", "1,9 millió"},
		{"99900000", "99,9 millió"},
		{"99900000000", "99,9 milliárd"},
		{"99900000000000", "99,9 billió"},
		{"2000000", "2 millió"},
		{"2300000", "2,3 millió"},
		{"1700000000", "1,7 milliárd"},
		{"1000000000", "1 milliárd"},
		{"1100000000", "1,1 milliárd"},
		{"1200000000", "1,2 milliárd"},
		{"1300000000", "1,3 milliárd"},
		{"1400000000", "1,4 milliárd"},
		{"1500000000", "1,5 milliárd"},
		{"1600000000", "1,6 milliárd"},
		{"1700000000", "1,7 milliárd"},
		{"1800000000", "1,8 milliárd"},
		{"1900000000", "1,9 milliárd"},
		{"2000000000", "2 milliárd"},
		{"3000000000", "3 milliárd"},
		{"4000000000", "4 milliárd"},
		{"5000000000", "5 milliárd"},
		{"6000000000", "6 milliárd"},
		{"7000000000", "7 milliárd"},
		{"8000000000", "8 milliárd"},
		{"9000000000", "9 milliárd"},
		{"10000000000", "10 milliárd"},
		{"11000000000", "11 milliárd"},
		{"12000000000", "12 milliárd"},
		{"13000000000", "13 milliárd"},
		{"14000000000", "14 milliárd"},
		{"15000000000", "15 milliárd"},
		{"16000000000", "16 milliárd"},
		{"17000000000", "17 milliárd"},
		{"18000000000", "18 milliárd"},
		{"19000000000", "19 milliárd"},
		{"20000000000", "20 milliárd"},
		{"100000000000", "100 milliárd"},
		{"200000000000", "200 milliárd"},
		{"300000000000", "300 milliárd"},
		{"400000000000", "400 milliárd"},
		{"500000000000", "500 milliárd"},
		{"600000000000", "600 milliárd"},
		{"700000000000", "700 milliárd"},
		{"800000000000", "800 milliárd"},
		{"900000000000", "900 milliárd"},
		{"1000000000000", "1 billió"},
		{"1100000000000", "1,1 billió"},
		{"1200000000000", "1,2 billió"},
		{"1300000000000", "1,3 billió"},
		{"1400000000000", "1,4 billió"},
		{"1500000000000", "1,5 billió"},
		{"1600000000000", "1,6 billió"},
		{"1700000000000", "1,7 billió"},
		{"1800000000000", "1,8 billió"},
		{"1900000000000", "1,9 billió"},
		{"2000000000000", "2 billió"},
		{"3000000000000", "3 billió"},
		{"4000000000000", "4 billió"},
		{"5000000000000", "5 billió"},
		{"6000000000000", "6 billió"},
		{"7000000000000", "7 billió"},
		{"8000000000000", "8 billió"},
		{"9000000000000", "9 billió"},
		{"10000000000000", "10 billió"},
	}

	h := humanize.NewHumanizer(locale.Data, humanize.OptionLong, fallback)

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

func TestHumanizeHuOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 E"},
		{"2000", "2 E"},
		{"3000", "3 E"},
		{"4000", "4 E"},
		{"5000", "5 E"},
		{"11000", "11 E"},
		{"12000", "12 E"},
		{"13000", "13 E"},
		{"14000", "14 E"},
		{"15000", "15 E"},
		{"11100", "11,1 E"},
		{"12100", "12,1 E"},
		{"13100", "13,1 E"},
		{"14100", "14,1 E"},
		{"15100", "15,1 E"},
		{"10000", "10 E"},
		{"12500", "12,5 E"},
		{"15000", "15 E"},
		{"15100", "15,1 E"},
		{"99500", "99,5 E"},
		{"1000000", "1 M"},
		{"10100000", "10,1 M"},
		{"99900000", "99,9 M"},
	}

	h := humanize.NewHumanizer(locale.Data, humanize.OptionShort, fallback)

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
