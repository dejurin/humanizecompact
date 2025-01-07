package locale_test

import (
	"testing"

	"github.com/dejurin/humanize-cldr"
	locale "github.com/dejurin/humanize-cldr/locales/cs"
)

func fallback(number string) string {
	return number
}

func TestHumanizeCsOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 tisíc"},
		{"2000", "2 tisíce"},
		{"3000", "3 tisíce"},
		{"4000", "4 tisíce"},
		{"5000", "5 tisíc"},
		{"11000", "11 tisíc"},
		{"12000", "12 tisíc"},
		{"13000", "13 tisíc"},
		{"14000", "14 tisíc"},
		{"15000", "15 tisíc"},
		{"11100", "11,1 tisíce"},
		{"12100", "12,1 tisíce"},
		{"13100", "13,1 tisíce"},
		{"14100", "14,1 tisíce"},
		{"15100", "15,1 tisíce"},
		{"99500", "99,5 tisíce"},
		{"1000000", "1 milion"},
		{"2000000", "2 miliony"},
		{"2300000", "2,3 milionu"},
		{"1100000", "1,1 milionu"},
		{"1200000", "1,2 milionu"},
		{"1300000", "1,3 milionu"},
		{"1400000", "1,4 milionu"},
		{"1500000", "1,5 milionu"},
		{"99900000", "99,9 milionu"},
		{"1000000000", "1 miliarda"},
		{"2000000000", "2 miliardy"},
		{"3000000000", "3 miliardy"},
		{"4000000000", "4 miliardy"},
		{"5000000000", "5 miliard"},
		{"1700000000", "1,7 miliardy"},
		{"99900000000", "99,9 miliardy"},
		{"1000000000000", "1 bilion"},
		{"2000000000000", "2 biliony"},
		{"3000000000000", "3 biliony"},
		{"4000000000000", "4 biliony"},
		{"5000000000000", "5 bilionů"},
		{"1100000000000", "1,1 bilionu"},
		{"1000000000000", "1 bilion"},
		{"1100000000000", "1,1 bilionu"},
		{"1200000000000", "1,2 bilionu"},
		{"1300000000000", "1,3 bilionu"},
		{"1400000000000", "1,4 bilionu"},
		{"1500000000000", "1,5 bilionu"},
		{"1600000000000", "1,6 bilionu"},
		{"1700000000000", "1,7 bilionu"},
		{"1800000000000", "1,8 bilionu"},
		{"1900000000000", "1,9 bilionu"},
		{"2000000000000", "2 biliony"},
		{"3000000000000", "3 biliony"},
		{"4000000000000", "4 biliony"},
		{"5000000000000", "5 bilionů"},
		{"6000000000000", "6 bilionů"},
		{"7000000000000", "7 bilionů"},
		{"8000000000000", "8 bilionů"},
		{"9000000000000", "9 bilionů"},
		{"10000000000000", "10 bilionů"},
		{"11000000000000", "11 bilionů"},
		{"12000000000000", "12 bilionů"},
		{"13000000000000", "13 bilionů"},
		{"14000000000000", "14 bilionů"},
		{"15000000000000", "15 bilionů"},
		{"16000000000000", "16 bilionů"},
		{"17000000000000", "17 bilionů"},
		{"18000000000000", "18 bilionů"},
		{"19000000000000", "19 bilionů"},
		{"20000000000000", "20 bilionů"},
		{"30000000000000", "30 bilionů"},
		{"40000000000000", "40 bilionů"},
		{"50000000000000", "50 bilionů"},
		{"60000000000000", "60 bilionů"},
		{"70000000000000", "70 bilionů"},
		{"80000000000000", "80 bilionů"},
		{"90000000000000", "90 bilionů"},
		{"99900000000000", "99,9 bilionu"},
		{"100000000000000", "100 bilionů"},
		{"200000000000000", "200 bilionů"},
		{"300000000000000", "300 bilionů"},
		{"400000000000000", "400 bilionů"},
		{"500000000000000", "500 bilionů"},
		{"600000000000000", "600 bilionů"},
		{"700000000000000", "700 bilionů"},
		{"800000000000000", "800 bilionů"},
		{"900000000000000", "900 bilionů"},
		{"999000000000000", "999 bilionů"},
		{"999000000000000", "999 bilionů"},
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

func TestHumanizeCsOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 tis."},
		{"2000", "2 tis."},
		{"3000", "3 tis."},
		{"4000", "4 tis."},
		{"5000", "5 tis."},
		{"11000", "11 tis."},
		{"12000", "12 tis."},
		{"13000", "13 tis."},
		{"14000", "14 tis."},
		{"15000", "15 tis."},
		{"1000000", "1 mil."},
		{"1100000", "1,1 mil."},
		{"1200000", "1,2 mil."},
		{"1300000", "1,3 mil."},
		{"1400000", "1,4 mil."},
		{"1500000", "1,5 mil."},
		{"1600000", "1,6 mil."},
		{"1700000", "1,7 mil."},
		{"1800000", "1,8 mil."},
		{"1900000", "1,9 mil."},

		{"10000000", "10 mil."},
		{"1000000000", "1 mld."},
		{"1100000000", "1,1 mld."},
		{"1200000000", "1,2 mld."},
		{"1300000000", "1,3 mld."},
		{"1400000000", "1,4 mld."},
		{"1500000000", "1,5 mld."},
		{"1600000000", "1,6 mld."},
		{"1700000000", "1,7 mld."},
		{"1800000000", "1,8 mld."},
		{"1900000000", "1,9 mld."},
		{"10000000000", "10 mld."},
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
