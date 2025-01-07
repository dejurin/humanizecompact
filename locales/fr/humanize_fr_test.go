package locale_test

import (
	"testing"

	humanize "github.com/dejurin/humanize-cldr"
	locale "github.com/dejurin/humanize-cldr/locales/fr"
)

func fallback(number string) string {
	return number
}

func TestHumanizeFrOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "mille"},
		{"2000", "2 mille"},
		{"3000", "3 mille"},
		{"4000", "4 mille"},
		{"5000", "5 mille"},
		{"11000", "11 mille"},
		{"12000", "12 mille"},
		{"13000", "13 mille"},
		{"14000", "14 mille"},
		{"15000", "15 mille"},
		{"11100", "11,1 mille"},
		{"12100", "12,1 mille"},
		{"13100", "13,1 mille"},
		{"14100", "14,1 mille"},
		{"15100", "15,1 mille"},
		{"10000", "10 mille"},
		{"12500", "12,5 mille"},
		{"15000", "15 mille"},
		{"15100", "15,1 mille"},
		{"99500", "99,5 mille"},
		{"1000000", "1 million"},
		{"10100000", "10,1 millions"},
		{"1100000", "1,1 million"},
		{"1200000", "1,2 million"},
		{"1300000", "1,3 million"},
		{"1400000", "1,4 million"},
		{"1500000", "1,5 million"},
		{"1600000", "1,6 millions"},
		{"1700000", "1,7 millions"},
		{"1800000", "1,8 millions"},
		{"1900000", "1,9 millions"},
		{"99900000", "99,9 millions"},
		{"99900000000", "99,9 milliards"},
		{"99900000000000", "99,9 billions"},
		{"2000000", "2 millions"},
		{"2300000", "2,3 millions"},
		{"1000000000", "1 milliard"},
		{"1100000000", "1,1 milliard"},
		{"1200000000", "1,2 milliard"},
		{"1300000000", "1,3 milliard"},
		{"1400000000", "1,4 milliard"},
		{"1500000000", "1,5 milliard"},
		{"1600000000", "1,6 milliards"},
		{"1700000000", "1,7 milliards"},
		{"1800000000", "1,8 milliards"},
		{"1900000000", "1,9 milliards"},
		{"2000000000", "2 milliards"},
		{"3000000000", "3 milliards"},
		{"4000000000", "4 milliards"},
		{"5000000000", "5 milliards"},
		{"6000000000", "6 milliards"},
		{"7000000000", "7 milliards"},
		{"8000000000", "8 milliards"},
		{"9000000000", "9 milliards"},
		{"10000000000", "10 milliards"},
		{"11000000000", "11 milliards"},
		{"12000000000", "12 milliards"},
		{"13000000000", "13 milliards"},
		{"14000000000", "14 milliards"},
		{"15000000000", "15 milliards"},
		{"16000000000", "16 milliards"},
		{"17000000000", "17 milliards"},
		{"18000000000", "18 milliards"},
		{"19000000000", "19 milliards"},
		{"20000000000", "20 milliards"},
		{"100000000000", "100 milliards"},
		{"200000000000", "200 milliards"},
		{"300000000000", "300 milliards"},
		{"400000000000", "400 milliards"},
		{"500000000000", "500 milliards"},
		{"600000000000", "600 milliards"},
		{"700000000000", "700 milliards"},
		{"800000000000", "800 milliards"},
		{"900000000000", "900 milliards"},
		{"1000000000000", "1 billion"},
		{"1100000000000", "1,1 billion"},
		{"1200000000000", "1,2 billion"},
		{"1300000000000", "1,3 billion"},
		{"1400000000000", "1,4 billion"},
		{"1500000000000", "1,5 billion"},
		{"1600000000000", "1,6 billions"},
		{"1700000000000", "1,7 billions"},
		{"1800000000000", "1,8 billions"},
		{"1900000000000", "1,9 billions"},
		{"2000000000000", "2 billions"},
		{"3000000000000", "3 billions"},
		{"4000000000000", "4 billions"},
		{"5000000000000", "5 billions"},
		{"6000000000000", "6 billions"},
		{"7000000000000", "7 billions"},
		{"8000000000000", "8 billions"},
		{"9000000000000", "9 billions"},
		{"10000000000000", "10 billions"},
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

func TestHumanizeFrOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 k"},
		{"2000", "2 k"},
		{"3000", "3 k"},
		{"4000", "4 k"},
		{"5000", "5 k"},

		{"11000", "11 k"},
		{"12000", "12 k"},
		{"13000", "13 k"},
		{"14000", "14 k"},
		{"15000", "15 k"},

		{"11100", "11,1 k"},
		{"12100", "12,1 k"},
		{"13100", "13,1 k"},
		{"14100", "14,1 k"},
		{"15100", "15,1 k"},

		{"1000000", "1 M"},
		{"1100000", "1,1 M"},
		{"1200000", "1,2 M"},
		{"1300000", "1,3 M"},
		{"1400000", "1,4 M"},
		{"1500000", "1,5 M"},
		{"1600000", "1,6 M"},
		{"1700000", "1,7 M"},
		{"1800000", "1,8 M"},
		{"1900000", "1,9 M"},
		{"99900000", "99,9 M"},

		{"99900000000", "99,9 Md"},
		{"99900000000000", "99,9 Bn"},
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
