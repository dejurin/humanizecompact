package locale

import (
	hc "github.com/dejurin/humanizecompact"
	"github.com/govalues/decimal"
)

type Locale struct {
	data       hc.CldrData
	localeCode string
}

func (l Locale) Data() hc.CldrData {
	return l.data
}

func (l Locale) Code() string {
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
	localeCode: "tr",
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 bin",
				"1000-count-other":            "0 bin",
				"10000-count-one":             "00 bin",
				"10000-count-other":           "00 bin",
				"100000-count-one":            "000 bin",
				"100000-count-other":          "000 bin",
				"1000000-count-one":           "0 milyon",
				"1000000-count-other":         "0 milyon",
				"10000000-count-one":          "00 milyon",
				"10000000-count-other":        "00 milyon",
				"100000000-count-one":         "000 milyon",
				"100000000-count-other":       "000 milyon",
				"1000000000-count-one":        "0 milyar",
				"1000000000-count-other":      "0 milyar",
				"10000000000-count-one":       "00 milyar",
				"10000000000-count-other":     "00 milyar",
				"100000000000-count-one":      "000 milyar",
				"100000000000-count-other":    "000 milyar",
				"1000000000000-count-one":     "0 trilyon",
				"1000000000000-count-other":   "0 trilyon",
				"10000000000000-count-one":    "00 trilyon",
				"10000000000000-count-other":  "00 trilyon",
				"100000000000000-count-one":   "000 trilyon",
				"100000000000000-count-other": "000 trilyon",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 B",
				"1000-count-other":            "0 B",
				"10000-count-one":             "00 B",
				"10000-count-other":           "00 B",
				"100000-count-one":            "000 B",
				"100000-count-other":          "000 B",
				"1000000-count-one":           "0 Mn",
				"1000000-count-other":         "0 Mn",
				"10000000-count-one":          "00 Mn",
				"10000000-count-other":        "00 Mn",
				"100000000-count-one":         "000 Mn",
				"100000000-count-other":       "000 Mn",
				"1000000000-count-one":        "0 Mr",
				"1000000000-count-other":      "0 Mr",
				"10000000000-count-one":       "00 Mr",
				"10000000000-count-other":     "00 Mr",
				"100000000000-count-one":      "000 Mr",
				"100000000000-count-other":    "000 Mr",
				"1000000000000-count-one":     "0 Tn",
				"1000000000000-count-other":   "0 Tn",
				"10000000000000-count-one":    "00 Tn",
				"10000000000000-count-other":  "00 Tn",
				"100000000000000-count-one":   "000 Tn",
				"100000000000000-count-other": "000 Tn",
			},
		},
	},
}
