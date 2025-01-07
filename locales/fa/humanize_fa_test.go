package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/fa"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[language.Tag]hc.Locale{
	language.Persian: locale.Data,
}

func TestHumanizeFaOptionLong(t *testing.T) {
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
		{"1000", "۱ هزار"},
		{"2000", "۲ هزار"},
		{"3000", "۳ هزار"},
		{"4000", "۴ هزار"},
		{"5000", "۵ هزار"},
		{"11000", "۱۱ هزار"},
		{"12000", "۱۲ هزار"},
		{"13000", "۱۳ هزار"},
		{"14000", "۱۴ هزار"},
		{"15000", "۱۵ هزار"},
		{"11100", "۱۱٫۱ هزار"},
		{"12100", "۱۲٫۱ هزار"},
		{"13100", "۱۳٫۱ هزار"},
		{"14100", "۱۴٫۱ هزار"},
		{"15100", "۱۵٫۱ هزار"},
		{"10000", "۱۰ هزار"},
		{"12500", "۱۲٫۵ هزار"},
		{"15000", "۱۵ هزار"},
		{"15100", "۱۵٫۱ هزار"},
		{"99500", "۹۹٫۵ هزار"},
		{"1000000", "۱ میلیون"},
		{"10100000", "۱۰٫۱ میلیون"},
		{"1100000", "۱٫۱ میلیون"},
		{"1200000", "۱٫۲ میلیون"},
		{"1300000", "۱٫۳ میلیون"},
		{"1400000", "۱٫۴ میلیون"},
		{"1500000", "۱٫۵ میلیون"},
		{"1600000", "۱٫۶ میلیون"},
		{"1700000", "۱٫۷ میلیون"},
		{"1800000", "۱٫۸ میلیون"},
		{"1900000", "۱٫۹ میلیون"},
		{"99900000", "۹۹٫۹ میلیون"},
		{"99900000000", "۹۹٫۹ میلیارد"},
		{"99900000000000", "۹۹٫۹ هزارمیلیارد"},
		{"2000000", "۲ میلیون"},
		{"2300000", "۲٫۳ میلیون"},
		{"1700000000", "۱٫۷ میلیارد"},
		{"1000000000", "۱ میلیارد"},
		{"1100000000", "۱٫۱ میلیارد"},
		{"1200000000", "۱٫۲ میلیارد"},
		{"1300000000", "۱٫۳ میلیارد"},
		{"1400000000", "۱٫۴ میلیارد"},
		{"1500000000", "۱٫۵ میلیارد"},
		{"1600000000", "۱٫۶ میلیارد"},
		{"1700000000", "۱٫۷ میلیارد"},
		{"1800000000", "۱٫۸ میلیارد"},
		{"1900000000", "۱٫۹ میلیارد"},
		{"2000000000", "۲ میلیارد"},
		{"3000000000", "۳ میلیارد"},
		{"4000000000", "۴ میلیارد"},
		{"5000000000", "۵ میلیارد"},
		{"6000000000", "۶ میلیارد"},
		{"7000000000", "۷ میلیارد"},
		{"8000000000", "۸ میلیارد"},
		{"9000000000", "۹ میلیارد"},
		{"10000000000", "۱۰ میلیارد"},
		{"11000000000", "۱۱ میلیارد"},
		{"12000000000", "۱۲ میلیارد"},
		{"13000000000", "۱۳ میلیارد"},
		{"14000000000", "۱۴ میلیارد"},
		{"15000000000", "۱۵ میلیارد"},
		{"16000000000", "۱۶ میلیارد"},
		{"17000000000", "۱۷ میلیارد"},
		{"18000000000", "۱۸ میلیارد"},
		{"19000000000", "۱۹ میلیارد"},
		{"20000000000", "۲۰ میلیارد"},
		{"100000000000", "۱۰۰ میلیارد"},
		{"200000000000", "۲۰۰ میلیارد"},
		{"300000000000", "۳۰۰ میلیارد"},
		{"400000000000", "۴۰۰ میلیارد"},
		{"500000000000", "۵۰۰ میلیارد"},
		{"600000000000", "۶۰۰ میلیارد"},
		{"700000000000", "۷۰۰ میلیارد"},
		{"800000000000", "۸۰۰ میلیارد"},
		{"900000000000", "۹۰۰ میلیارد"},
		{"1000000000000", "۱ هزارمیلیارد"},
		{"1100000000000", "۱٫۱ هزارمیلیارد"},
		{"1200000000000", "۱٫۲ هزارمیلیارد"},
		{"1300000000000", "۱٫۳ هزارمیلیارد"},
		{"1400000000000", "۱٫۴ هزارمیلیارد"},
		{"1500000000000", "۱٫۵ هزارمیلیارد"},
		{"1600000000000", "۱٫۶ هزارمیلیارد"},
		{"1700000000000", "۱٫۷ هزارمیلیارد"},
		{"1800000000000", "۱٫۸ هزارمیلیارد"},
		{"1900000000000", "۱٫۹ هزارمیلیارد"},
		{"2000000000000", "۲ هزارمیلیارد"},
		{"3000000000000", "۳ هزارمیلیارد"},
		{"4000000000000", "۴ هزارمیلیارد"},
		{"5000000000000", "۵ هزارمیلیارد"},
		{"6000000000000", "۶ هزارمیلیارد"},
		{"7000000000000", "۷ هزارمیلیارد"},
		{"8000000000000", "۸ هزارمیلیارد"},
		{"9000000000000", "۹ هزارمیلیارد"},
		{"10000000000000", "۱۰ هزارمیلیارد"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.Persian)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeFaOptionShort(t *testing.T) {
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
		{"1000", "۱ هزار"},
		{"2000", "۲ هزار"},
		{"3000", "۳ هزار"},
		{"4000", "۴ هزار"},
		{"5000", "۵ هزار"},
		{"11000", "۱۱ هزار"},
		{"12000", "۱۲ هزار"},
		{"13000", "۱۳ هزار"},
		{"14000", "۱۴ هزار"},
		{"15000", "۱۵ هزار"},
		{"11100", "۱۱٫۱ هزار"},
		{"12100", "۱۲٫۱ هزار"},
		{"13100", "۱۳٫۱ هزار"},
		{"14100", "۱۴٫۱ هزار"},
		{"15100", "۱۵٫۱ هزار"},
		{"10000", "۱۰ هزار"},
		{"12500", "۱۲٫۵ هزار"},
		{"15000", "۱۵ هزار"},
		{"1000000", "۱ میلیون"},
		{"10100000", "۱۰٫۱ میلیون"},
		{"11000000", "۱۱ میلیون"},
		{"12000000", "۱۲ میلیون"},
		{"13000000", "۱۳ میلیون"},
		{"14000000", "۱۴ میلیون"},
		{"15000000", "۱۵ میلیون"},
		{"1000000000", "۱ میلیارد"},
		{"11000000000", "۱۱ میلیارد"},
		{"120000000000", "۱۲۰ میلیارد"},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.Persian)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
