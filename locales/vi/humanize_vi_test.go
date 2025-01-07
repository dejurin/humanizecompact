package locale_test

import (
	"testing"

	"github.com/dejurin/humanize-cldr"
	locale "github.com/dejurin/humanize-cldr/locales/vi"
)

func fallback(number string) string {
	return number
}

func TestHumanizeViOptionLong(t *testing.T) {
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
		{"1000", "1 nghìn"},
		{"2000", "2 nghìn"},
		{"3000", "3 nghìn"},
		{"4000", "4 nghìn"},
		{"5000", "5 nghìn"},
		{"11000", "11 nghìn"},
		{"12000", "12 nghìn"},
		{"13000", "13 nghìn"},
		{"14000", "14 nghìn"},
		{"15000", "15 nghìn"},
		{"11100", "11,1 nghìn"},
		{"12100", "12,1 nghìn"},
		{"13100", "13,1 nghìn"},
		{"14100", "14,1 nghìn"},
		{"15100", "15,1 nghìn"},
		{"10000", "10 nghìn"},
		{"12500", "12,5 nghìn"},
		{"15000", "15 nghìn"},
		{"15100", "15,1 nghìn"},
		{"99500", "99,5 nghìn"},
		{"1000000", "1 triệu"},
		{"10100000", "10,1 triệu"},
		{"1100000", "1,1 triệu"},
		{"1200000", "1,2 triệu"},
		{"1300000", "1,3 triệu"},
		{"1400000", "1,4 triệu"},
		{"1500000", "1,5 triệu"},
		{"1600000", "1,6 triệu"},
		{"1700000", "1,7 triệu"},
		{"1800000", "1,8 triệu"},
		{"1900000", "1,9 triệu"},
		{"99900000", "99,9 triệu"},
		{"99900000000", "99,9 tỷ"},
		{"99900000000000", "99,9 nghìn tỷ"},
		{"2000000", "2 triệu"},
		{"2300000", "2,3 triệu"},
		{"1700000000", "1,7 tỷ"},
		{"1000000000", "1 tỷ"},
		{"1100000000", "1,1 tỷ"},
		{"1200000000", "1,2 tỷ"},
		{"1300000000", "1,3 tỷ"},
		{"1400000000", "1,4 tỷ"},
		{"1500000000", "1,5 tỷ"},
		{"1600000000", "1,6 tỷ"},
		{"1700000000", "1,7 tỷ"},
		{"1800000000", "1,8 tỷ"},
		{"1900000000", "1,9 tỷ"},
		{"2000000000", "2 tỷ"},
		{"3000000000", "3 tỷ"},
		{"4000000000", "4 tỷ"},
		{"5000000000", "5 tỷ"},
		{"6000000000", "6 tỷ"},
		{"7000000000", "7 tỷ"},
		{"8000000000", "8 tỷ"},
		{"9000000000", "9 tỷ"},
		{"10000000000", "10 tỷ"},
		{"11000000000", "11 tỷ"},
		{"12000000000", "12 tỷ"},
		{"13000000000", "13 tỷ"},
		{"14000000000", "14 tỷ"},
		{"15000000000", "15 tỷ"},
		{"16000000000", "16 tỷ"},
		{"17000000000", "17 tỷ"},
		{"18000000000", "18 tỷ"},
		{"19000000000", "19 tỷ"},
		{"20000000000", "20 tỷ"},
		{"100000000000", "100 tỷ"},
		{"200000000000", "200 tỷ"},
		{"300000000000", "300 tỷ"},
		{"400000000000", "400 tỷ"},
		{"500000000000", "500 tỷ"},
		{"600000000000", "600 tỷ"},
		{"700000000000", "700 tỷ"},
		{"800000000000", "800 tỷ"},
		{"900000000000", "900 tỷ"},
		{"1000000000000", "1 nghìn tỷ"},
		{"1100000000000", "1,1 nghìn tỷ"},
		{"1200000000000", "1,2 nghìn tỷ"},
		{"1300000000000", "1,3 nghìn tỷ"},
		{"1400000000000", "1,4 nghìn tỷ"},
		{"1500000000000", "1,5 nghìn tỷ"},
		{"1600000000000", "1,6 nghìn tỷ"},
		{"1700000000000", "1,7 nghìn tỷ"},
		{"1800000000000", "1,8 nghìn tỷ"},
		{"1900000000000", "1,9 nghìn tỷ"},
		{"2000000000000", "2 nghìn tỷ"},
		{"3000000000000", "3 nghìn tỷ"},
		{"4000000000000", "4 nghìn tỷ"},
		{"5000000000000", "5 nghìn tỷ"},
		{"6000000000000", "6 nghìn tỷ"},
		{"7000000000000", "7 nghìn tỷ"},
		{"8000000000000", "8 nghìn tỷ"},
		{"9000000000000", "9 nghìn tỷ"},
		{"10000000000000", "10 nghìn tỷ"},
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

func TestHumanizeViOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 N"},
		{"10000", "10 N"},
		{"100000", "100 N"},
		{"1000000", "1 Tr"},
		{"1100000", "1,1 Tr"},
		{"1200000", "1,2 Tr"},
		{"1300000", "1,3 Tr"},
		{"1400000", "1,4 Tr"},
		{"1500000", "1,5 Tr"},
		{"10000000", "10 Tr"},
		{"100000000", "100 Tr"},
		{"1000000000", "1 T"},
		{"10000000000", "10 T"},
		{"100000000000", "100 T"},
		{"1000000000000", "1 NT"},
		{"10000000000000", "10 NT"},
		{"100000000000000", "100 NT"},
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
