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

// "pluralRule-count-one": "i = 0,1 @integer 0, 1 @decimal 0.0~1.5",
// "pluralRule-count-many": "e = 0 and i != 0 and i % 1000000 = 0 and v = 0 or e != 0..5 @integer 1000000, 1c6, 2c6, 3c6, 4c6, 5c6, 6c6, … @decimal 1.0000001c6, 1.1c6, 2.0000001c6, 2.1c6, 3.0000001c6, 3.1c6, …",
// "pluralRule-count-other": " @integer 2~17, 100, 1000, 10000, 100000, 1c3, 2c3, 3c3, 4c3, 5c3, 6c3, … @decimal 2.0~3.5, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, 1.0001c3, 1.1c3, 2.0001c3, 2.1c3, 3.0001c3, 3.1c3, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	zero, err := decimal.NewFromFloat64(0.0)
	one, err := decimal.NewFromInt64(1, 0, 0)

	if err != nil {
		return "other"
	}

	upper, err := decimal.NewFromFloat64(1.6)
	if err != nil {
		return "other"
	}

	if r.Cmp(one) == 0 && v == "1000" {
		return "1"
	}

	if r.Cmp(zero) == 1 && r.Cmp(upper) == -1 {
		return "one"
	}

	return "other"
}

var Data hc.Locale = Locale{
	localeCode: language.French,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-1":                "mille",
				"1000-count-one":              "0 millier",
				"1000-count-other":            "0 mille",
				"10000-count-one":             "00 mille",
				"10000-count-other":           "00 mille",
				"100000-count-one":            "000 mille",
				"100000-count-other":          "000 mille",
				"1000000-count-one":           "0 million",
				"1000000-count-other":         "0 millions",
				"10000000-count-one":          "00 million",
				"10000000-count-other":        "00 millions",
				"100000000-count-one":         "000 million",
				"100000000-count-other":       "000 millions",
				"1000000000-count-one":        "0 milliard",
				"1000000000-count-other":      "0 milliards",
				"10000000000-count-one":       "00 milliard",
				"10000000000-count-other":     "00 milliards",
				"100000000000-count-one":      "000 milliard",
				"100000000000-count-other":    "000 milliards",
				"1000000000000-count-one":     "0 billion",
				"1000000000000-count-other":   "0 billions",
				"10000000000000-count-one":    "00 billion",
				"10000000000000-count-other":  "00 billions",
				"100000000000000-count-one":   "000 billion",
				"100000000000000-count-other": "000 billions",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 k",
				"1000-count-other":            "0 k",
				"10000-count-one":             "00 k",
				"10000-count-other":           "00 k",
				"100000-count-one":            "000 k",
				"100000-count-other":          "000 k",
				"1000000-count-one":           "0 M",
				"1000000-count-other":         "0 M",
				"10000000-count-one":          "00 M",
				"10000000-count-other":        "00 M",
				"100000000-count-one":         "000 M",
				"100000000-count-other":       "000 M",
				"1000000000-count-one":        "0 Md",
				"1000000000-count-other":      "0 Md",
				"10000000000-count-one":       "00 Md",
				"10000000000-count-other":     "00 Md",
				"100000000000-count-one":      "000 Md",
				"100000000000-count-other":    "000 Md",
				"1000000000000-count-one":     "0 Bn",
				"1000000000000-count-other":   "0 Bn",
				"10000000000000-count-one":    "00 Bn",
				"10000000000000-count-other":  "00 Bn",
				"100000000000000-count-one":   "000 Bn",
				"100000000000000-count-other": "000 Bn",
			},
		},
	},
}
