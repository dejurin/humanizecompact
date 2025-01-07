package locale

import (
	humanize "github.com/dejurin/humanize-cldr"
	"github.com/govalues/decimal"
)

type Locale struct {
	data       humanize.CldrData
	localeCode string
}

func (l Locale) Data() humanize.CldrData {
	return l.data
}

func (l Locale) Code() string {
	return l.localeCode
}

// "pluralRule-count-one": "i = 1 and v = 0 @integer 1",
// "pluralRule-count-other": " @integer 0, 2~16, 100, 1000, 10000, 100000, 1000000, … @decimal 0.0~1.5, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	if !r.IsInt() {
		return "other"
	}

	i64, scale, ok := r.Int64(0)
	if !ok {
		return "other"
	}

	if scale != 0 {
		return "other"
	}

	if i64 == 1 {
		return "one"
	}

	return "other"
}

var Data humanize.Locale = Locale{
	localeCode: "de",
	data: humanize.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 Tausend",
				"1000-count-other":            "0 Tausend",
				"10000-count-one":             "00 Tausend",
				"10000-count-other":           "00 Tausend",
				"100000-count-one":            "000 Tausend",
				"100000-count-other":          "000 Tausend",
				"1000000-count-one":           "0 Million",
				"1000000-count-other":         "0 Millionen",
				"10000000-count-one":          "00 Millionen",
				"10000000-count-other":        "00 Millionen",
				"100000000-count-one":         "000 Millionen",
				"100000000-count-other":       "000 Millionen",
				"1000000000-count-one":        "0 Milliarde",
				"1000000000-count-other":      "0 Milliarden",
				"10000000000-count-one":       "00 Milliarden",
				"10000000000-count-other":     "00 Milliarden",
				"100000000000-count-one":      "000 Milliarden",
				"100000000000-count-other":    "000 Milliarden",
				"1000000000000-count-one":     "0 Billion",
				"1000000000000-count-other":   "0 Billionen",
				"10000000000000-count-one":    "00 Billionen",
				"10000000000000-count-other":  "00 Billionen",
				"100000000000000-count-one":   "000 Billionen",
				"100000000000000-count-other": "000 Billionen",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0",
				"1000-count-other":            "0",
				"10000-count-one":             "0",
				"10000-count-other":           "0",
				"100000-count-one":            "0",
				"100000-count-other":          "0",
				"1000000-count-one":           "0 Mio.",
				"1000000-count-other":         "0 Mio.",
				"10000000-count-one":          "00 Mio.",
				"10000000-count-other":        "00 Mio.",
				"100000000-count-one":         "000 Mio.",
				"100000000-count-other":       "000 Mio.",
				"1000000000-count-one":        "0 Mrd.",
				"1000000000-count-other":      "0 Mrd.",
				"10000000000-count-one":       "00 Mrd.",
				"10000000000-count-other":     "00 Mrd.",
				"100000000000-count-one":      "000 Mrd.",
				"100000000000-count-other":    "000 Mrd.",
				"1000000000000-count-one":     "0 Bio.",
				"1000000000000-count-other":   "0 Bio.",
				"10000000000000-count-one":    "00 Bio.",
				"10000000000000-count-other":  "00 Bio.",
				"100000000000000-count-one":   "000 Bio.",
				"100000000000000-count-other": "000 Bio.",
			},
		},
	},
}
