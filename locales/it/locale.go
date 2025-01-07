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
// "pluralRule-count-many": "e = 0 and i != 0 and i % 1000000 = 0 and v = 0 or e != 0..5 @integer 1000000, 1c6, 2c6, 3c6, 4c6, 5c6, 6c6, … @decimal 1.0000001c6, 1.1c6, 2.0000001c6, 2.1c6, 3.0000001c6, 3.1c6, …",
// "pluralRule-count-other": " @integer 0, 2~16, 100, 1000, 10000, 100000, 1c3, 2c3, 3c3, 4c3, 5c3, 6c3, … @decimal 0.0~1.5, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, 1.0001c3, 1.1c3, 2.0001c3, 2.1c3, 3.0001c3, 3.1c3, …"
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
	localeCode: language.Italian,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "mille",
				"1000-count-other":            "0 mila",
				"10000-count-one":             "00 mila",
				"10000-count-other":           "00 mila",
				"100000-count-one":            "000 mila",
				"100000-count-other":          "000 mila",
				"1000000-count-one":           "0 milione",
				"1000000-count-other":         "0 milioni",
				"10000000-count-one":          "00 milioni",
				"10000000-count-other":        "00 milioni",
				"100000000-count-one":         "000 milioni",
				"100000000-count-other":       "000 milioni",
				"1000000000-count-one":        "0 miliardo",
				"1000000000-count-other":      "0 miliardi",
				"10000000000-count-one":       "00 miliardi",
				"10000000000-count-other":     "00 miliardi",
				"100000000000-count-one":      "000 miliardi",
				"100000000000-count-other":    "000 miliardi",
				"1000000000000-count-one":     "mille miliardi",
				"1000000000000-count-other":   "0 mila miliardi",
				"10000000000000-count-one":    "00 mila miliardi",
				"10000000000000-count-other":  "00 mila miliardi",
				"100000000000000-count-one":   "000 mila miliardi",
				"100000000000000-count-other": "000 mila miliardi",
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
				"1000000-count-one":           "0 Mln",
				"1000000-count-other":         "0 Mln",
				"10000000-count-one":          "00 Mln",
				"10000000-count-other":        "00 Mln",
				"100000000-count-one":         "000 Mln",
				"100000000-count-other":       "000 Mln",
				"1000000000-count-one":        "0 Mld",
				"1000000000-count-other":      "0 Mld",
				"10000000000-count-one":       "00 Mld",
				"10000000000-count-other":     "00 Mld",
				"100000000000-count-one":      "000 Mld",
				"100000000000-count-other":    "000 Mld",
				"1000000000000-count-one":     "0 Bln",
				"1000000000000-count-other":   "0 Bln",
				"10000000000000-count-one":    "00 Bln",
				"10000000000000-count-other":  "00 Bln",
				"100000000000000-count-one":   "000 Bln",
				"100000000000000-count-other": "000 Bln",
			},
		},
	},
}
