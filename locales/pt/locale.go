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

// "pluralRule-count-one": "i = 0..1 @integer 0, 1 @decimal 0.0~1.5",
// "pluralRule-count-many": "e = 0 and i != 0 and i % 1000000 = 0 and v = 0 or e != 0..5 @integer 1000000, 1c6, 2c6, 3c6, 4c6, 5c6, 6c6, … @decimal 1.0000001c6, 1.1c6, 2.0000001c6, 2.1c6, 3.0000001c6, 3.1c6, …",
// "pluralRule-count-other": " @integer 2~17, 100, 1000, 10000, 100000, 1c3, 2c3, 3c3, 4c3, 5c3, 6c3, … @decimal 2.0~3.5, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, 1.0001c3, 1.1c3, 2.0001c3, 2.1c3, 3.0001c3, 3.1c3, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	zero, err := decimal.NewFromFloat64(0)
	if err != nil {
		return "other"
	}

	upper, err := decimal.NewFromFloat64(1.6)
	if err != nil {
		return "other"
	}

	if r.Cmp(zero) == 1 && r.Cmp(upper) == -1 {
		return "one"
	}

	return "other"
}

var Data hc.Locale = Locale{
	localeCode: language.Portuguese,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 mil",
				"1000-count-other":            "0 mil",
				"10000-count-one":             "00 mil",
				"10000-count-other":           "00 mil",
				"100000-count-one":            "000 mil",
				"100000-count-other":          "000 mil",
				"1000000-count-one":           "0 milhão",
				"1000000-count-other":         "0 milhões",
				"10000000-count-one":          "00 milhão",
				"10000000-count-other":        "00 milhões",
				"100000000-count-one":         "000 milhão",
				"100000000-count-other":       "000 milhões",
				"1000000000-count-one":        "0 bilhão",
				"1000000000-count-other":      "0 bilhões",
				"10000000000-count-one":       "00 bilhão",
				"10000000000-count-other":     "00 bilhões",
				"100000000000-count-one":      "000 bilhão",
				"100000000000-count-other":    "000 bilhões",
				"1000000000000-count-one":     "0 trilhão",
				"1000000000000-count-other":   "0 trilhões",
				"10000000000000-count-one":    "00 trilhão",
				"10000000000000-count-other":  "00 trilhões",
				"100000000000000-count-one":   "000 trilhão",
				"100000000000000-count-other": "000 trilhões",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 mil",
				"1000-count-other":            "0 mil",
				"10000-count-one":             "00 mil",
				"10000-count-other":           "00 mil",
				"100000-count-one":            "000 mil",
				"100000-count-other":          "000 mil",
				"1000000-count-one":           "0 mi",
				"1000000-count-other":         "0 mi",
				"10000000-count-one":          "00 mi",
				"10000000-count-other":        "00 mi",
				"100000000-count-one":         "000 mi",
				"100000000-count-other":       "000 mi",
				"1000000000-count-one":        "0 bi",
				"1000000000-count-other":      "0 bi",
				"10000000000-count-one":       "00 bi",
				"10000000000-count-other":     "00 bi",
				"100000000000-count-one":      "000 bi",
				"100000000000-count-other":    "000 bi",
				"1000000000000-count-one":     "0 tri",
				"1000000000000-count-other":   "0 tri",
				"10000000000000-count-one":    "00 tri",
				"10000000000000-count-other":  "00 tri",
				"100000000000000-count-one":   "000 tri",
				"100000000000000-count-other": "000 tri",
			},
		},
	},
}
