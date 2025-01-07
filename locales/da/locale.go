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

// "pluralRule-count-one": "n = 1 or t != 0 and i = 0,1 @integer 1 @decimal 0.1~1.6",
// "pluralRule-count-other": " @integer 0, 2~16, 100, 1000, 10000, 100000, 1000000, … @decimal 0.0, 2.0~3.4, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {

	one, _ := decimal.NewFromInt64(1, 0, 0)
	if r.Equal(one) {
		return "one"
	}

	if !r.IsInt() {
		lowerBound, _ := decimal.NewFromFloat64(0.1)
		upperBound, _ := decimal.NewFromFloat64(1.6)

		if r.Cmp(lowerBound) == 1 && r.Cmp(upperBound) == -1 {
			return "one"
		}
		return "other"
	}

	return "other"
}

var Data hc.Locale = Locale{
	localeCode: language.Danish,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 tusind",
				"1000-count-other":            "0 tusind",
				"10000-count-one":             "00 tusind",
				"10000-count-other":           "00 tusind",
				"100000-count-one":            "000 tusind",
				"100000-count-other":          "000 tusind",
				"1000000-count-one":           "0 million",
				"1000000-count-other":         "0 millioner",
				"10000000-count-one":          "00 millioner",
				"10000000-count-other":        "00 millioner",
				"100000000-count-one":         "000 millioner",
				"100000000-count-other":       "000 millioner",
				"1000000000-count-one":        "0 milliard",
				"1000000000-count-other":      "0 milliarder",
				"10000000000-count-one":       "00 milliarder",
				"10000000000-count-other":     "00 milliarder",
				"100000000000-count-one":      "000 milliarder",
				"100000000000-count-other":    "000 milliarder",
				"1000000000000-count-one":     "0 billion",
				"1000000000000-count-other":   "0 billioner",
				"10000000000000-count-one":    "00 billioner",
				"10000000000000-count-other":  "00 billioner",
				"100000000000000-count-one":   "000 billioner",
				"100000000000000-count-other": "000 billioner",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 t",
				"1000-count-other":            "0 t",
				"10000-count-one":             "00 t",
				"10000-count-other":           "00 t",
				"100000-count-one":            "000 t",
				"100000-count-other":          "000 t",
				"1000000-count-one":           "0 mio.",
				"1000000-count-other":         "0 mio.",
				"10000000-count-one":          "00 mio.",
				"10000000-count-other":        "00 mio.",
				"100000000-count-one":         "000 mio.",
				"100000000-count-other":       "000 mio.",
				"1000000000-count-one":        "0 mia.",
				"1000000000-count-other":      "0 mia.",
				"10000000000-count-one":       "00 mia.",
				"10000000000-count-other":     "00 mia.",
				"100000000000-count-one":      "000 mia.",
				"100000000000-count-other":    "000 mia.",
				"1000000000000-count-one":     "0 bio.",
				"1000000000000-count-other":   "0 bio.",
				"10000000000000-count-one":    "00 bio.",
				"10000000000000-count-other":  "00 bio.",
				"100000000000000-count-one":   "000 bio.",
				"100000000000000-count-other": "000 bio.",
			},
		},
	},
}
