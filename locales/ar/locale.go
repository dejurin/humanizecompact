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

// "pluralRule-count-zero": "n = 0 @integer 0 @decimal 0.0, 0.00, 0.000, 0.0000",
// "pluralRule-count-one": "n = 1 @integer 1 @decimal 1.0, 1.00, 1.000, 1.0000",
// "pluralRule-count-two": "n = 2 @integer 2 @decimal 2.0, 2.00, 2.000, 2.0000",
// "pluralRule-count-few": "n % 100 = 3..10 @integer 3~10, 103~110, 1003, … @decimal 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 103.0, 1003.0, …",
// "pluralRule-count-many": "n % 100 = 11..99 @integer 11~26, 111, 1011, … @decimal 11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 111.0, 1011.0, …",
// "pluralRule-count-other": " @integer 100~102, 200~202, 300~302, 400~402, 500~502, 600, 1000, 10000, 100000, 1000000, … @decimal 0.1~0.9, 1.1~1.7, 10.1, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	if !r.IsInt() {
		return "other"
	}

	i64, _, ok := r.Int64(0)
	if !ok {
		return "other"
	}

	rem := i64 % 100

	switch {
	case i64 == 0:
		return "zero"

	case i64 == 1:
		return "one"

	case i64 == 2:
		return "two"

	case rem >= 3 && rem <= 10:
		return "few"

	case rem >= 11 && rem <= 99:
		return "many"

	default:
		return "other"
	}
}

var Data hc.Locale = Locale{
	localeCode: language.Arabic,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-zero":             "0 ألف",
				"1000-count-one":              "0 ألف",
				"1000-count-two":              "0 ألف",
				"1000-count-few":              "0 آلاف",
				"1000-count-many":             "0 ألف",
				"1000-count-other":            "0 ألف",
				"10000-count-zero":            "00 ألف",
				"10000-count-one":             "00 ألف",
				"10000-count-two":             "00 ألف",
				"10000-count-few":             "00 ألف",
				"10000-count-many":            "00 ألف",
				"10000-count-other":           "00 ألف",
				"100000-count-zero":           "000 ألف",
				"100000-count-one":            "000 ألف",
				"100000-count-two":            "000 ألف",
				"100000-count-few":            "000 ألف",
				"100000-count-many":           "000 ألف",
				"100000-count-other":          "000 ألف",
				"1000000-count-zero":          "0 مليون",
				"1000000-count-one":           "0 مليون",
				"1000000-count-two":           "0 مليون",
				"1000000-count-few":           "0 ملايين",
				"1000000-count-many":          "0 مليون",
				"1000000-count-other":         "0 مليون",
				"10000000-count-zero":         "00 مليون",
				"10000000-count-one":          "00 مليون",
				"10000000-count-two":          "00 مليون",
				"10000000-count-few":          "00 ملايين",
				"10000000-count-many":         "00 مليون",
				"10000000-count-other":        "00 مليون",
				"100000000-count-zero":        "000 مليون",
				"100000000-count-one":         "000 مليون",
				"100000000-count-two":         "000 مليون",
				"100000000-count-few":         "000 مليون",
				"100000000-count-many":        "000 مليون",
				"100000000-count-other":       "000 مليون",
				"1000000000-count-zero":       "0 مليار",
				"1000000000-count-one":        "0 مليار",
				"1000000000-count-two":        "0 مليار",
				"1000000000-count-few":        "0 مليار",
				"1000000000-count-many":       "0 مليار",
				"1000000000-count-other":      "0 مليار",
				"10000000000-count-zero":      "00 مليار",
				"10000000000-count-one":       "00 مليار",
				"10000000000-count-two":       "00 مليار",
				"10000000000-count-few":       "00 مليار",
				"10000000000-count-many":      "00 مليار",
				"10000000000-count-other":     "00 مليار",
				"100000000000-count-zero":     "000 مليار",
				"100000000000-count-one":      "000 مليار",
				"100000000000-count-two":      "000 مليار",
				"100000000000-count-few":      "000 مليار",
				"100000000000-count-many":     "000 مليار",
				"100000000000-count-other":    "000 مليار",
				"1000000000000-count-zero":    "0 ترليون",
				"1000000000000-count-one":     "0 ترليون",
				"1000000000000-count-two":     "0 ترليون",
				"1000000000000-count-few":     "0 ترليون",
				"1000000000000-count-many":    "0 ترليون",
				"1000000000000-count-other":   "0 ترليون",
				"10000000000000-count-zero":   "00 ترليون",
				"10000000000000-count-one":    "00 ترليون",
				"10000000000000-count-two":    "00 ترليون",
				"10000000000000-count-few":    "00 ترليون",
				"10000000000000-count-many":   "00 ترليون",
				"10000000000000-count-other":  "00 ترليون",
				"100000000000000-count-zero":  "000 ترليون",
				"100000000000000-count-one":   "000 ترليون",
				"100000000000000-count-two":   "000 ترليون",
				"100000000000000-count-few":   "000 ترليون",
				"100000000000000-count-many":  "000 ترليون",
				"100000000000000-count-other": "000 ترليون",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-zero":             "0 ألف",
				"1000-count-one":              "0 ألف",
				"1000-count-two":              "0 ألف",
				"1000-count-few":              "0 آلاف",
				"1000-count-many":             "0 ألف",
				"1000-count-other":            "0 ألف",
				"10000-count-zero":            "00 ألف",
				"10000-count-one":             "00 ألف",
				"10000-count-two":             "00 ألف",
				"10000-count-few":             "00 ألف",
				"10000-count-many":            "00 ألف",
				"10000-count-other":           "00 ألف",
				"100000-count-zero":           "000 ألف",
				"100000-count-one":            "000 ألف",
				"100000-count-two":            "000 ألف",
				"100000-count-few":            "000 ألف",
				"100000-count-many":           "000 ألف",
				"100000-count-other":          "000 ألف",
				"1000000-count-one":           "0 مليون",
				"1000000-count-other":         "0 مليون",
				"10000000-count-one":          "00 مليون",
				"10000000-count-other":        "00 مليون",
				"100000000-count-one":         "000 مليون",
				"100000000-count-other":       "000 مليون",
				"1000000000-count-one":        "0 مليار",
				"1000000000-count-other":      "0 مليار",
				"10000000000-count-one":       "00 مليار",
				"10000000000-count-other":     "00 مليار",
				"100000000000-count-one":      "000 مليار",
				"100000000000-count-other":    "000 مليار",
				"1000000000000-count-one":     "0 ترليون",
				"1000000000000-count-other":   "0 ترليون",
				"10000000000000-count-one":    "00 ترليون",
				"10000000000000-count-other":  "00 ترليون",
				"100000000000000-count-one":   "000 ترليون",
				"100000000000000-count-other": "000 ترليون",
			},
		},
	},
}
