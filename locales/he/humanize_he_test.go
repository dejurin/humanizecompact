package locale_test

import (
	"testing"

	"github.com/dejurin/humanize-cldr"
	locale "github.com/dejurin/humanize-cldr/locales/he"
)

func fallback(number string) string {
	return number
}

func TestHumanizeHeOptionLong(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1 אלף"},
		{"2000", "2 אלף"},
		{"3000", "3 אלף"},
		{"4000", "4 אלף"},
		{"5000", "5 אלף"},
		{"10000", "10 אלף"},
		{"21000", "21 אלף"},
		{"31000", "31 אלף"},
		{"41000", "41 אלף"},
		{"51000", "51 אלף"},
		{"1000000", "1 מיליון"},
		{"2000000", "2 מיליון"},
		{"3000000", "3 מיליון"},
		{"4000000", "4 מיליון"},
		{"5000000", "5 מיליון"},
		{"1000000000", "1 מיליארד"},
		{"2000000000", "2 מיליארד"},
		{"3000000000", "3 מיליארד"},
		{"4000000000", "4 מיליארד"},
		{"5000000000", "5 מיליארד"},
		{"1000000000000", "1 טריליון"},
		{"2000000000000", "2 טריליון"},
		{"3000000000000", "3 טריליון"},
		{"4000000000000", "4 טריליון"},
		{"5000000000000", "5 טריליון"},
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

func TestHumanizeHeOptionShort(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{"1000", "1K"},
		{"10000", "10K"},
		{"100000", "100K"},
		{"1000000", "1M"},
		{"1100000", "1.1M"},
		{"1200000", "1.2M"},
		{"1300000", "1.3M"},
		{"1400000", "1.4M"},
		{"1500000", "1.5M"},
		{"10000000", "10M"},
		{"100000000", "100M"},
		{"1000000000", "1B"},
		{"10000000000", "10B"},
		{"100000000000", "100B"},
		{"1000000000000", "1T"},
		{"10000000000000", "10T"},
		{"100000000000000", "100T"},
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
