package locale

import (
	hc "github.com/dejurin/humanizecompact"
	"github.com/govalues/decimal"
	"golang.org/x/text/language"
)

type Locale struct {
	data       hc.CldrData
	localeCode language.Tag
}

func (l Locale) Data() hc.CldrData {
	return l.data
}

func (l Locale) Code() language.Tag {
	return l.localeCode
}

// "pluralRule-count-one": "n = 1 @integer 1 @decimal 1.0, 1.00, 1.000, 1.0000",
// "pluralRule-count-other": " @integer 0, 2~16, 100, 1000, 10000, 100000, 1000000, … @decimal 0.0~0.9, 1.1~1.6, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	i64, _, ok := r.Int64(0)
	if !ok {
		return "other"
	}

	if i64 == 1 {
		return "one"
	}
	return "other"
}

var Data hc.Locale = Locale{
	localeCode: language.Bulgarian,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 хил.",
				"1000-count-other":            "0 хиляди",
				"10000-count-one":             "00 хиляди",
				"10000-count-other":           "00 хиляди",
				"100000-count-one":            "000 хиляди",
				"100000-count-other":          "000 хиляди",
				"1000000-count-one":           "0 милион",
				"1000000-count-other":         "0 милиона",
				"10000000-count-one":          "00 милиона",
				"10000000-count-other":        "00 милиона",
				"100000000-count-one":         "000 милиона",
				"100000000-count-other":       "000 милиона",
				"1000000000-count-one":        "0 милиард",
				"1000000000-count-other":      "0 милиарда",
				"10000000000-count-one":       "00 милиарда",
				"10000000000-count-other":     "00 милиарда",
				"100000000000-count-one":      "000 милиарда",
				"100000000000-count-other":    "000 милиарда",
				"1000000000000-count-one":     "0 трилион",
				"1000000000000-count-other":   "0 трилиона",
				"10000000000000-count-one":    "00 трилиона",
				"10000000000000-count-other":  "00 трилиона",
				"100000000000000-count-one":   "000 трилиона",
				"100000000000000-count-other": "000 трилиона",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 хил.",
				"1000-count-other":            "0 хил.",
				"10000-count-one":             "00 хил.",
				"10000-count-other":           "00 хил.",
				"100000-count-one":            "000 хил.",
				"100000-count-other":          "000 хил.",
				"1000000-count-one":           "0 млн.",
				"1000000-count-other":         "0 млн.",
				"10000000-count-one":          "00 млн.",
				"10000000-count-other":        "00 млн.",
				"100000000-count-one":         "000 млн.",
				"100000000-count-other":       "000 млн.",
				"1000000000-count-one":        "0 млрд.",
				"1000000000-count-other":      "0 млрд.",
				"10000000000-count-one":       "00 млрд.",
				"10000000000-count-other":     "00 млрд.",
				"100000000000-count-one":      "000 млрд.",
				"100000000000-count-other":    "000 млрд.",
				"1000000000000-count-one":     "0 трлн.",
				"1000000000000-count-other":   "0 трлн.",
				"10000000000000-count-one":    "00 трлн.",
				"10000000000000-count-other":  "00 трлн.",
				"100000000000000-count-one":   "000 трлн.",
				"100000000000000-count-other": "000 трлн.",
			},
		},
	},
}
