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

// "pluralRule-count-one": "i = 1 and v = 0 @integer 1",
// "pluralRule-count-other": " @integer 0, 2~16, 100, 1000, 10000, 100000, 1000000, … @decimal 0.0~1.5, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	if !r.IsInt() {
		return "other"
	}

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
	localeCode: "sv",
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 tusen",
				"1000-count-other":            "0 tusen",
				"10000-count-one":             "00 tusen",
				"10000-count-other":           "00 tusen",
				"100000-count-one":            "000 tusen",
				"100000-count-other":          "000 tusen",
				"1000000-count-one":           "0 miljon",
				"1000000-count-other":         "0 miljoner",
				"10000000-count-one":          "00 miljon",
				"10000000-count-other":        "00 miljoner",
				"100000000-count-one":         "000 miljoner",
				"100000000-count-other":       "000 miljoner",
				"1000000000-count-one":        "0 miljard",
				"1000000000-count-other":      "0 miljarder",
				"10000000000-count-one":       "00 miljarder",
				"10000000000-count-other":     "00 miljarder",
				"100000000000-count-one":      "000 miljarder",
				"100000000000-count-other":    "000 miljarder",
				"1000000000000-count-one":     "0 biljon",
				"1000000000000-count-other":   "0 biljoner",
				"10000000000000-count-one":    "00 biljoner",
				"10000000000000-count-other":  "00 biljoner",
				"100000000000000-count-one":   "000 biljoner",
				"100000000000000-count-other": "000 biljoner",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 tn",
				"1000-count-other":            "0 tn",
				"10000-count-one":             "00 tn",
				"10000-count-other":           "00 tn",
				"100000-count-one":            "000 tn",
				"100000-count-other":          "000 tn",
				"1000000-count-one":           "0 mn",
				"1000000-count-other":         "0 mn",
				"10000000-count-one":          "00 mn",
				"10000000-count-other":        "00 mn",
				"100000000-count-one":         "000 mn",
				"100000000-count-other":       "000 mn",
				"1000000000-count-one":        "0 md",
				"1000000000-count-other":      "0 md",
				"10000000000-count-one":       "00 md",
				"10000000000-count-other":     "00 md",
				"100000000000-count-one":      "000 md",
				"100000000000-count-other":    "000 md",
				"1000000000000-count-one":     "0 bn",
				"1000000000000-count-other":   "0 bn",
				"10000000000000-count-one":    "00 bn",
				"10000000000000-count-other":  "00 bn",
				"100000000000000-count-one":   "000 bn",
				"100000000000000-count-other": "000 bn",
			},
		},
	},
}
