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
// "pluralRule-count-few": "v = 0 and i % 10 = 2..4 and i % 100 != 12..14 @integer 2~4, 22~24, 32~34, 42~44, 52~54, 62, 102, 1002, …",
// "pluralRule-count-many": "v = 0 and i != 1 and i % 10 = 0..1 or v = 0 and i % 10 = 5..9 or v = 0 and i % 100 = 12..14 @integer 0, 5~19, 100, 1000, 10000, 100000, 1000000, …",
// "pluralRule-count-other": "   @decimal 0.0~1.5, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	if !r.IsInt() {
		return "other"
	}

	i64, _, ok := r.Int64(0)
	if !ok {
		return "other"
	}

	i10 := i64 % 10
	i100 := i64 % 100

	switch {
	case i64 == 1:
		return "one"

	case (i10 >= 2 && i10 <= 4) && !(i100 >= 12 && i100 <= 14):
		return "few"

	case (i64 != 1 && (i10 == 0 || i10 == 1)) ||
		(i10 >= 5 && i10 <= 9) ||
		(i100 >= 12 && i100 <= 14):
		return "many"

	default:
		return "other"
	}
}

var Data hc.Locale = Locale{
	localeCode: language.Polish,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 tysiąc",
				"1000-count-few":              "0 tysiące",
				"1000-count-many":             "0 tysięcy",
				"1000-count-other":            "0 tysiąca",
				"10000-count-one":             "00 tysiąc",
				"10000-count-few":             "00 tysiące",
				"10000-count-many":            "00 tysięcy",
				"10000-count-other":           "00 tysiąca",
				"100000-count-one":            "000 tysiąc",
				"100000-count-few":            "000 tysiące",
				"100000-count-many":           "000 tysięcy",
				"100000-count-other":          "000 tysiąca",
				"1000000-count-one":           "0 milion",
				"1000000-count-few":           "0 miliony",
				"1000000-count-many":          "0 milionów",
				"1000000-count-other":         "0 miliona",
				"10000000-count-one":          "00 milion",
				"10000000-count-few":          "00 miliony",
				"10000000-count-many":         "00 milionów",
				"10000000-count-other":        "00 miliona",
				"100000000-count-one":         "000 milion",
				"100000000-count-few":         "000 miliony",
				"100000000-count-many":        "000 milionów",
				"100000000-count-other":       "000 miliona",
				"1000000000-count-one":        "0 miliard",
				"1000000000-count-few":        "0 miliardy",
				"1000000000-count-many":       "0 miliardów",
				"1000000000-count-other":      "0 miliarda",
				"10000000000-count-one":       "00 miliard",
				"10000000000-count-few":       "00 miliardy",
				"10000000000-count-many":      "00 miliardów",
				"10000000000-count-other":     "00 miliarda",
				"100000000000-count-one":      "000 miliard",
				"100000000000-count-few":      "000 miliardy",
				"100000000000-count-many":     "000 miliardów",
				"100000000000-count-other":    "000 miliarda",
				"1000000000000-count-one":     "0 bilion",
				"1000000000000-count-few":     "0 biliony",
				"1000000000000-count-many":    "0 bilionów",
				"1000000000000-count-other":   "0 biliona",
				"10000000000000-count-one":    "00 bilion",
				"10000000000000-count-few":    "00 biliony",
				"10000000000000-count-many":   "00 bilionów",
				"10000000000000-count-other":  "00 biliona",
				"100000000000000-count-one":   "000 bilion",
				"100000000000000-count-few":   "000 biliony",
				"100000000000000-count-many":  "000 bilionów",
				"100000000000000-count-other": "000 biliona",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 tys.",
				"1000-count-other":            "0 tys.",
				"10000-count-one":             "00 tys.",
				"10000-count-other":           "00 tys.",
				"100000-count-one":            "000 tys.",
				"100000-count-other":          "000 tys.",
				"1000000-count-one":           "0 mln",
				"1000000-count-other":         "0 mln",
				"10000000-count-one":          "00 mln",
				"10000000-count-other":        "00 mln",
				"100000000-count-one":         "000 mln",
				"100000000-count-other":       "000 mln",
				"1000000000-count-one":        "0 mld",
				"1000000000-count-other":      "0 mld",
				"10000000000-count-one":       "00 mld",
				"10000000000-count-other":     "00 mld",
				"100000000000-count-one":      "000 mld",
				"100000000000-count-other":    "000 mld",
				"1000000000000-count-one":     "0 bln",
				"1000000000000-count-other":   "0 bln",
				"10000000000000-count-one":    "00 bln",
				"10000000000000-count-other":  "00 bln",
				"100000000000000-count-one":   "000 bln",
				"100000000000000-count-other": "000 bln",
			},
		},
	},
}
