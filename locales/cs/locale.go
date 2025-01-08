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
// "pluralRule-count-few": "i = 2..4 and v = 0 @integer 2~4",
// "pluralRule-count-many": "v != 0   @decimal 0.0~1.5, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …",
// "pluralRule-count-other": " @integer 0, 5~19, 100, 1000, 10000, 100000, 1000000, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	if !r.IsInt() {
		return "many"
	}

	i64, _, ok := r.Int64(0)
	if !ok {
		return "other"
	}

	switch i64 {
	case 1:
		return "one"
	case 2, 3, 4:
		return "few"
	default:
		return "other"
	}
}

var Data hc.Locale = Locale{
	localeCode: language.Czech,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 tisíc",
				"1000-count-few":              "0 tisíce",
				"1000-count-many":             "0 tisíce",
				"1000-count-other":            "0 tisíc",
				"10000-count-one":             "00 tisíc",
				"10000-count-few":             "00 tisíc",
				"10000-count-many":            "00 tisíce",
				"10000-count-other":           "00 tisíc",
				"100000-count-one":            "000 tisíc",
				"100000-count-few":            "000 tisíc",
				"100000-count-many":           "000 tisíce",
				"100000-count-other":          "000 tisíc",
				"1000000-count-one":           "0 milion",
				"1000000-count-few":           "0 miliony",
				"1000000-count-many":          "0 milionu",
				"1000000-count-other":         "0 milionů",
				"10000000-count-one":          "00 milionů",
				"10000000-count-few":          "00 milionů",
				"10000000-count-many":         "00 milionu",
				"10000000-count-other":        "00 milionů",
				"100000000-count-one":         "000 milionů",
				"100000000-count-few":         "000 milionů",
				"100000000-count-many":        "000 milionu",
				"100000000-count-other":       "000 milionů",
				"1000000000-count-one":        "0 miliarda",
				"1000000000-count-few":        "0 miliardy",
				"1000000000-count-many":       "0 miliardy",
				"1000000000-count-other":      "0 miliard",
				"10000000000-count-one":       "00 miliard",
				"10000000000-count-few":       "00 miliard",
				"10000000000-count-many":      "00 miliardy",
				"10000000000-count-other":     "00 miliard",
				"100000000000-count-one":      "000 miliard",
				"100000000000-count-few":      "000 miliard",
				"100000000000-count-many":     "000 miliardy",
				"100000000000-count-other":    "000 miliard",
				"1000000000000-count-one":     "0 bilion",
				"1000000000000-count-few":     "0 biliony",
				"1000000000000-count-many":    "0 bilionu",
				"1000000000000-count-other":   "0 bilionů",
				"10000000000000-count-one":    "00 bilionů",
				"10000000000000-count-few":    "00 bilionů",
				"10000000000000-count-many":   "00 bilionu",
				"10000000000000-count-other":  "00 bilionů",
				"100000000000000-count-one":   "000 bilionů",
				"100000000000000-count-few":   "000 bilionů",
				"100000000000000-count-many":  "000 bilionu",
				"100000000000000-count-other": "000 bilionů",
			},
		},

		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 tis.",
				"1000-count-other":            "0 tis.",
				"10000-count-one":             "00 tis.",
				"10000-count-other":           "00 tis.",
				"100000-count-one":            "000 tis.",
				"100000-count-other":          "000 tis.",
				"1000000-count-one":           "0 mil.",
				"1000000-count-other":         "0 mil.",
				"10000000-count-one":          "00 mil.",
				"10000000-count-other":        "00 mil.",
				"100000000-count-one":         "000 mil.",
				"100000000-count-other":       "000 mil.",
				"1000000000-count-one":        "0 mld.",
				"1000000000-count-other":      "0 mld.",
				"10000000000-count-one":       "00 mld.",
				"10000000000-count-other":     "00 mld.",
				"100000000000-count-one":      "000 mld.",
				"100000000000-count-other":    "000 mld.",
				"1000000000000-count-one":     "0 bil.",
				"1000000000000-count-other":   "0 bil.",
				"10000000000000-count-one":    "00 bil.",
				"10000000000000-count-other":  "00 bil.",
				"100000000000000-count-one":   "000 bil.",
				"100000000000000-count-other": "000 bil.",
			},
		},
	},
}
