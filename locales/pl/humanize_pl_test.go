package locale_test

import (
	"testing"

	"golang.org/x/text/language"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/pl"
)

func fallback(number string) string {
	return number
}

var locales = map[language.Tag]hc.Locale{
	language.Polish: locale.Data,
}

func TestHumanizePlOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 tysiąc"},
		{"2000", "2 tysiące"},
		{"3000", "3 tysiące"},
		{"4000", "4 tysiące"},
		{"5000", "5 tysięcy"},
		{"11000", "11 tysięcy"},
		{"12000", "12 tysięcy"},
		{"13000", "13 tysięcy"},
		{"14000", "14 tysięcy"},
		{"15000", "15 tysięcy"},
		{"11100", "11,1 tysiąca"},
		{"12100", "12,1 tysiąca"},
		{"13100", "13,1 tysiąca"},
		{"14100", "14,1 tysiąca"},
		{"15100", "15,1 tysiąca"},
		{"10000", "10 tysięcy"},
		{"12500", "12,5 tysiąca"},
		{"15000", "15 tysięcy"},
		{"15100", "15,1 tysiąca"},
		{"99500", "99,5 tysiąca"},
		{"1000000", "1 milion"},
		{"10100000", "10,1 miliona"},
		{"1100000", "1,1 miliona"},
		{"1200000", "1,2 miliona"},
		{"1300000", "1,3 miliona"},
		{"1400000", "1,4 miliona"},
		{"1500000", "1,5 miliona"},
		{"1600000", "1,6 miliona"},
		{"1700000", "1,7 miliona"},
		{"1800000", "1,8 miliona"},
		{"1900000", "1,9 miliona"},
		{"99900000", "99,9 miliona"},
		{"99900000000", "99,9 miliarda"},
		{"99900000000000", "99,9 biliona"},
		{"2000000", "2 miliony"},
		{"2300000", "2,3 miliona"},
		{"1700000000", "1,7 miliarda"},
		{"1000000000", "1 miliard"},
		{"1100000000", "1,1 miliarda"},
		{"1200000000", "1,2 miliarda"},
		{"1300000000", "1,3 miliarda"},
		{"1400000000", "1,4 miliarda"},
		{"1500000000", "1,5 miliarda"},
		{"1600000000", "1,6 miliarda"},
		{"1700000000", "1,7 miliarda"},
		{"1800000000", "1,8 miliarda"},
		{"1900000000", "1,9 miliarda"},
		{"2000000000", "2 miliardy"},
		{"3000000000", "3 miliardy"},
		{"4000000000", "4 miliardy"},
		{"5000000000", "5 miliardów"},
		{"6000000000", "6 miliardów"},
		{"7000000000", "7 miliardów"},
		{"8000000000", "8 miliardów"},
		{"9000000000", "9 miliardów"},
		{"10000000000", "10 miliardów"},
		{"11000000000", "11 miliardów"},
		{"12000000000", "12 miliardów"},
		{"13000000000", "13 miliardów"},
		{"14000000000", "14 miliardów"},
		{"15000000000", "15 miliardów"},
		{"16000000000", "16 miliardów"},
		{"17000000000", "17 miliardów"},
		{"18000000000", "18 miliardów"},
		{"19000000000", "19 miliardów"},
		{"20000000000", "20 miliardów"},
		{"100000000000", "100 miliardów"},
		{"200000000000", "200 miliardów"},
		{"300000000000", "300 miliardów"},
		{"400000000000", "400 miliardów"},
		{"500000000000", "500 miliardów"},
		{"600000000000", "600 miliardów"},
		{"700000000000", "700 miliardów"},
		{"800000000000", "800 miliardów"},
		{"900000000000", "900 miliardów"},
		{"1000000000000", "1 bilion"},
		{"1100000000000", "1,1 biliona"},
		{"1200000000000", "1,2 biliona"},
		{"1300000000000", "1,3 biliona"},
		{"1400000000000", "1,4 biliona"},
		{"1500000000000", "1,5 biliona"},
		{"1600000000000", "1,6 biliona"},
		{"1700000000000", "1,7 biliona"},
		{"1800000000000", "1,8 biliona"},
		{"1900000000000", "1,9 biliona"},
		{"2000000000000", "2 biliony"},
		{"3000000000000", "3 biliony"},
		{"4000000000000", "4 biliony"},
		{"5000000000000", "5 bilionów"},
		{"6000000000000", "6 bilionów"},
		{"7000000000000", "7 bilionów"},
		{"8000000000000", "8 bilionów"},
		{"9000000000000", "9 bilionów"},
		{"10000000000000", "10 bilionów"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Polish)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizePlOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1\u00A0tys."},
		{"2000", "2\u00A0tys."},
		{"3000", "3\u00A0tys."},
		{"4000", "4\u00A0tys."},
		{"5000", "5\u00A0tys."},
		{"11000", "11\u00A0tys."},
		{"12000", "12\u00A0tys."},
		{"13000", "13\u00A0tys."},
		{"14000", "14\u00A0tys."},
		{"15000", "15\u00A0tys."},
		{"11100", "11,1\u00A0tys."},
		{"12100", "12,1\u00A0tys."},
		{"13100", "13,1\u00A0tys."},
		{"14100", "14,1\u00A0tys."},
		{"15100", "15,1\u00A0tys."},
		{"10000", "10\u00A0tys."},
		{"12500", "12,5\u00A0tys."},
		{"15000", "15\u00A0tys."},
		{"15100", "15,1\u00A0tys."},
		{"99500", "99,5\u00A0tys."},
		{"1000000", "1\u00A0mln"},
		{"10100000", "10,1\u00A0mln"},
		{"1100000", "1,1\u00A0mln"},
		{"1200000", "1,2\u00A0mln"},
		{"1300000", "1,3\u00A0mln"},
		{"1400000", "1,4\u00A0mln"},
		{"1500000", "1,5\u00A0mln"},
		{"1600000", "1,6\u00A0mln"},
		{"1700000", "1,7\u00A0mln"},
		{"1800000", "1,8\u00A0mln"},
		{"1900000", "1,9\u00A0mln"},
		{"99900000", "99,9\u00A0mln"},
		{"99900000000", "99,9\u00A0mld"},
		{"2000000", "2\u00A0mln"},
		{"2300000", "2,3\u00A0mln"},
		{"1700000000", "1,7\u00A0mld"},
		{"1000000000", "1\u00A0mld"},
		{"1100000000", "1,1\u00A0mld"},
		{"99900000000000", "99,9\u00A0bln"},
		{"9000000000000", "9\u00A0bln"},
		{"8000000000000", "8\u00A0bln"},
		{"7000000000000", "7\u00A0bln"},
		{"6000000000000", "6\u00A0bln"},
		{"5000000000000", "5\u00A0bln"},
		{"4000000000000", "4\u00A0bln"},
		{"3000000000000", "3\u00A0bln"},
		{"2000000000000", "2\u00A0bln"},
		{"1000000000000", "1\u00A0bln"},
		{"1100000000000", "1,1\u00A0bln"},
		{"1200000000000", "1,2\u00A0bln"},
		{"1300000000000", "1,3\u00A0bln"},
		{"1400000000000", "1,4\u00A0bln"},
		{"1500000000000", "1,5\u00A0bln"},
		{"1600000000000", "1,6\u00A0bln"},
		{"1700000000000", "1,7\u00A0bln"},
		{"1800000000000", "1,8\u00A0bln"},
		{"1900000000000", "1,9\u00A0bln"},
		{"2000000000000", "2\u00A0bln"},
		{"3000000000000", "3\u00A0bln"},
		{"4000000000000", "4\u00A0bln"},
		{"5000000000000", "5\u00A0bln"},
		{"6000000000000", "6\u00A0bln"},
		{"7000000000000", "7\u00A0bln"},
		{"8000000000000", "8\u00A0bln"},
		{"9000000000000", "9\u00A0bln"},
		{"10000000000000", "10\u00A0bln"},
		{"11000000000000", "11\u00A0bln"},
		{"12000000000000", "12\u00A0bln"},
		{"13000000000000", "13\u00A0bln"},
		{"14000000000000", "14\u00A0bln"},
		{"15000000000000", "15\u00A0bln"},
		{"16000000000000", "16\u00A0bln"},
		{"17000000000000", "17\u00A0bln"},
		{"18000000000000", "18\u00A0bln"},
		{"19000000000000", "19\u00A0bln"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, _, err := h.Formatter(tt.number, language.Polish)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
