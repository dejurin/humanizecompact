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
	i64, _, ok := r.Int64(0)
	if !ok {
		return "other"
	}

	if i64 == 1 {
		return "one"
	}
	return "other"
}

var Data humanize.Locale = Locale{
	localeCode: "en",
	data: humanize.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 thousand",
				"1000-count-other":            "0 thousand",
				"10000-count-one":             "00 thousand",
				"10000-count-other":           "00 thousand",
				"100000-count-one":            "000 thousand",
				"100000-count-other":          "000 thousand",
				"1000000-count-one":           "0 million",
				"1000000-count-other":         "0 million",
				"10000000-count-one":          "00 million",
				"10000000-count-other":        "00 million",
				"100000000-count-one":         "000 million",
				"100000000-count-other":       "000 million",
				"1000000000-count-one":        "0 billion",
				"1000000000-count-other":      "0 billion",
				"10000000000-count-one":       "00 billion",
				"10000000000-count-other":     "00 billion",
				"100000000000-count-one":      "000 billion",
				"100000000000-count-other":    "000 billion",
				"1000000000000-count-one":     "0 trillion",
				"1000000000000-count-other":   "0 trillion",
				"10000000000000-count-one":    "00 trillion",
				"10000000000000-count-other":  "00 trillion",
				"100000000000000-count-one":   "000 trillion",
				"100000000000000-count-other": "000 trillion",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0K",
				"1000-count-other":            "0K",
				"10000-count-one":             "00K",
				"10000-count-other":           "00K",
				"100000-count-one":            "000K",
				"100000-count-other":          "000K",
				"1000000-count-one":           "0M",
				"1000000-count-other":         "0M",
				"10000000-count-one":          "00M",
				"10000000-count-other":        "00M",
				"100000000-count-one":         "000M",
				"100000000-count-other":       "000M",
				"1000000000-count-one":        "0B",
				"1000000000-count-other":      "0B",
				"10000000000-count-one":       "00B",
				"10000000000-count-other":     "00B",
				"100000000000-count-one":      "000B",
				"100000000000-count-other":    "000B",
				"1000000000000-count-one":     "0T",
				"1000000000000-count-other":   "0T",
				"10000000000000-count-one":    "00T",
				"10000000000000-count-other":  "00T",
				"100000000000000-count-one":   "000T",
				"100000000000000-count-other": "000T",
			},
		},
	},
}
