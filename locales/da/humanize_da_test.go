package locale_test

import (
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/da"
	"golang.org/x/text/language"
)

func fallback(number string) string {
	return number
}

var locales = map[language.Tag]hc.Locale{
	language.Danish: locale.Data,
}

func TestHumanizeDaOptionLong(t *testing.T) {
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
		{"1000", "1 tusind"},
		{"1500", "1,5 tusind"},
		{"1600", "1,6 tusind"},
		{"2000", "2 tusind"},
		{"3000", "3 tusind"},
		{"4000", "4 tusind"},
		{"5000", "5 tusind"},
		{"11000", "11 tusind"},
		{"12000", "12 tusind"},
		{"13000", "13 tusind"},
		{"14000", "14 tusind"},
		{"15000", "15 tusind"},
		{"11100", "11,1 tusind"},
		{"12100", "12,1 tusind"},
		{"13100", "13,1 tusind"},
		{"14100", "14,1 tusind"},
		{"15100", "15,1 tusind"},
		{"10000", "10 tusind"},
		{"12500", "12,5 tusind"},
		{"15000", "15 tusind"},
		{"15100", "15,1 tusind"},
		{"99500", "99,5 tusind"},
		{"1000000", "1 million"},
		{"1100000", "1,1 million"},
		{"1200000", "1,2 million"},
		{"1300000", "1,3 million"},
		{"1400000", "1,4 million"},
		{"1500000", "1,5 million"},
		{"1600000", "1,6 millioner"},
		{"1700000", "1,7 millioner"},
		{"1800000", "1,8 millioner"},
		{"1900000", "1,9 millioner"},
		{"10100000", "10,1 millioner"},
		{"2000000", "2 millioner"},
		{"99900000", "99,9 millioner"},
		{"99900000000", "99,9 milliarder"},
		{"99900000000000", "99,9 billioner"},
		{"2300000", "2,3 millioner"},
		{"1700000000", "1,7 milliarder"},
		{"1000000000", "1 milliard"},
		{"1100000000", "1,1 milliard"},
		{"1200000000", "1,2 milliard"},
		{"1300000000", "1,3 milliard"},
		{"1400000000", "1,4 milliard"},
		{"1500000000", "1,5 milliard"},
		{"1600000000", "1,6 milliarder"},
		{"1700000000", "1,7 milliarder"},
		{"1800000000", "1,8 milliarder"},
		{"1900000000", "1,9 milliarder"},
		{"2000000000", "2 milliarder"},
		{"3000000000", "3 milliarder"},
		{"4000000000", "4 milliarder"},
		{"5000000000", "5 milliarder"},
		{"6000000000", "6 milliarder"},
		{"7000000000", "7 milliarder"},
		{"8000000000", "8 milliarder"},
		{"9000000000", "9 milliarder"},
		{"10000000000", "10 milliarder"},
		{"11000000000", "11 milliarder"},
		{"12000000000", "12 milliarder"},
		{"13000000000", "13 milliarder"},
		{"14000000000", "14 milliarder"},
		{"15000000000", "15 milliarder"},
		{"16000000000", "16 milliarder"},
		{"17000000000", "17 milliarder"},
		{"18000000000", "18 milliarder"},
		{"19000000000", "19 milliarder"},
		{"20000000000", "20 milliarder"},
		{"100000000000", "100 milliarder"},
		{"200000000000", "200 milliarder"},
		{"300000000000", "300 milliarder"},
		{"400000000000", "400 milliarder"},
		{"500000000000", "500 milliarder"},
		{"600000000000", "600 milliarder"},
		{"700000000000", "700 milliarder"},
		{"800000000000", "800 milliarder"},
		{"900000000000", "900 milliarder"},
		{"1000000000000", "1 billion"},
		{"1100000000000", "1,1 billion"},
		{"1200000000000", "1,2 billion"},
		{"1300000000000", "1,3 billion"},
		{"1400000000000", "1,4 billion"},
		{"1500000000000", "1,5 billion"},
		{"1600000000000", "1,6 billioner"},
		{"1700000000000", "1,7 billioner"},
		{"1800000000000", "1,8 billioner"},
		{"1900000000000", "1,9 billioner"},
		{"2000000000000", "2 billioner"},
		{"3000000000000", "3 billioner"},
		{"4000000000000", "4 billioner"},
		{"5000000000000", "5 billioner"},
		{"6000000000000", "6 billioner"},
		{"7000000000000", "7 billioner"},
		{"8000000000000", "8 billioner"},
		{"9000000000000", "9 billioner"},
		{"10000000000000", "10 billioner"},
	}

	h := hc.New(locales, hc.Long, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.Danish)
		if err != nil {
			t.Errorf("number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}

func TestHumanizeDaOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 t"},
		{"1100", "1,1 t"},
		{"1200", "1,2 t"},
		{"1300", "1,3 t"},
		{"1400", "1,4 t"},
		{"1500", "1,5 t"},
		{"1600", "1,6 t"},
		{"2000", "2 t"},
		{"3000", "3 t"},
		{"4000", "4 t"},
		{"5000", "5 t"},
		{"1000000", "1 mio."},
		{"1100000", "1,1 mio."},
		{"1200000", "1,2 mio."},
		{"1300000", "1,3 mio."},
		{"1400000", "1,4 mio."},
		{"1500000", "1,5 mio."},
		{"1600000", "1,6 mio."},
		{"1700000", "1,7 mio."},
		{"1800000", "1,8 mio."},
		{"1900000", "1,9 mio."},
		{"2000000", "2 mio."},
		{"1000000000", "1 mia."},
		{"1100000000", "1,1 mia."},
		{"1200000000", "1,2 mia."},
		{"1300000000", "1,3 mia."},
		{"1400000000", "1,4 mia."},
		{"1500000000", "1,5 mia."},
		{"1600000000", "1,6 mia."},
		{"1700000000", "1,7 mia."},
		{"1000000000000", "1 bio."},
		{"1100000000000", "1,1 bio."},
		{"1200000000000", "1,2 bio."},
		{"1300000000000", "1,3 bio."},
		{"1400000000000", "1,4 bio."},
		{"1500000000000", "1,5 bio."},
		{"1600000000000", "1,6 bio."},
		{"1700000000000", "1,7 bio."},
		{"1800000000000", "1,8 bio."},
		{"1900000000000", "1,9 bio."},
		{"2000000000000", "2 bio."},
	}

	h := hc.New(locales, hc.Short, fallback)

	for _, tt := range tests {
		res, err := h.Formatter(tt.number, language.Danish)
		if err != nil {
			t.Errorf("[SHORT] number %q => unexpected error: %v", tt.number, err)
			continue
		}
		if res != tt.expected {
			t.Errorf("[SHORT] number %q => got %q, want %q", tt.number, res, tt.expected)
		}
	}
}
