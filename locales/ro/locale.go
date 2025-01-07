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

// "pluralRule-count-one": "i = 1 and v = 0 @integer 1",
// "pluralRule-count-few": "v != 0 or n = 0 or n != 1 and n % 100 = 1..19 @integer 0, 2~16, 101, 1001, … @decimal 0.0~1.5, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …",
// "pluralRule-count-other": " @integer 20~35, 100, 1000, 10000, 100000, 1000000, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	if !r.IsInt() {
		// Here you might want to check if v != 0 for "few" if fractional parts are relevant
		return "few"
	}

	i64, _, ok := r.Int64(0)
	if !ok {
		return "other"
	}

	rem := i64 % 100

	switch {
	case i64 == 1 && r.Scale() == 0:
		return "one"

	case i64 == 0 || (i64 != 1 && rem >= 1 && rem <= 19):
		return "few"

	default:
		return "other"
	}
}

var Data hc.Locale = Locale{
	localeCode: language.Romanian,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 mie",
				"1000-count-few":              "0 mii",
				"1000-count-other":            "0 de mii",
				"10000-count-one":             "00 mie",
				"10000-count-few":             "00 mii",
				"10000-count-other":           "00 de mii",
				"100000-count-one":            "000 mie",
				"100000-count-few":            "000 mii",
				"100000-count-other":          "000 de mii",
				"1000000-count-one":           "0 milion",
				"1000000-count-few":           "0 milioane",
				"1000000-count-other":         "0 de milioane",
				"10000000-count-one":          "00 milion",
				"10000000-count-few":          "00 milioane",
				"10000000-count-other":        "00 de milioane",
				"100000000-count-one":         "000 milion",
				"100000000-count-few":         "000 milioane",
				"100000000-count-other":       "000 de milioane",
				"1000000000-count-one":        "0 miliard",
				"1000000000-count-few":        "0 miliarde",
				"1000000000-count-other":      "0 de miliarde",
				"10000000000-count-one":       "00 miliard",
				"10000000000-count-few":       "00 miliarde",
				"10000000000-count-other":     "00 de miliarde",
				"100000000000-count-one":      "000 miliard",
				"100000000000-count-few":      "000 miliarde",
				"100000000000-count-other":    "000 de miliarde",
				"1000000000000-count-one":     "0 trilion",
				"1000000000000-count-few":     "0 trilioane",
				"1000000000000-count-other":   "0 de trilioane",
				"10000000000000-count-one":    "00 trilion",
				"10000000000000-count-few":    "00 trilioane",
				"10000000000000-count-other":  "00 de trilioane",
				"100000000000000-count-one":   "000 trilion",
				"100000000000000-count-few":   "000 trilioane",
				"100000000000000-count-other": "000 de trilioane",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 K",
				"1000-count-other":            "0 K",
				"10000-count-one":             "00 K",
				"10000-count-other":           "00 K",
				"100000-count-one":            "000 K",
				"100000-count-other":          "000 K",
				"1000000-count-one":           "0 mil.",
				"1000000-count-other":         "0 mil.",
				"10000000-count-one":          "00 mil.",
				"10000000-count-other":        "00 mil.",
				"100000000-count-one":         "000 mil.",
				"100000000-count-other":       "000 mil.",
				"1000000000-count-one":        "0 mld.",
				"1000000000-count-other":      "0 mld.",
				"10000000000-count-one":       "00 mld.",
				"10000000000-count-other":     "00 mld.",
				"100000000000-count-one":      "000 mld.",
				"100000000000-count-other":    "000 mld.",
				"1000000000000-count-one":     "0 tril.",
				"1000000000000-count-other":   "0 tril.",
				"10000000000000-count-one":    "00 tril.",
				"10000000000000-count-other":  "00 tril.",
				"100000000000000-count-one":   "000 tril.",
				"100000000000000-count-other": "000 tril.",
			},
		},
	},
}
