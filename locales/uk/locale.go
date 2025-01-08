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

// "pluralRule-count-one": "v = 0 and i % 10 = 1 and i % 100 != 11 @integer 1, 21, 31, 41, 51, 61, 71, 81, 101, 1001, …",
// "pluralRule-count-few": "v = 0 and i % 10 = 2..4 and i % 100 != 12..14 @integer 2~4, 22~24, 32~34, 42~44, 52~54, 62, 102, 1002, …",
// "pluralRule-count-many": "v = 0 and i % 10 = 0 or v = 0 and i % 10 = 5..9 or v = 0 and i % 100 = 11..14 @integer 0, 5~19, 100, 1000, 10000, 100000, 1000000, …",
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
	case i10 == 1 && i100 != 11:
		return "one"

	case i10 >= 2 && i10 <= 4 && !(i100 >= 12 && i100 <= 14):
		return "few"

	case (i10 == 0 || (i10 >= 5 && i10 <= 9)) || (i100 >= 11 && i100 <= 14):
		return "many"

	default:
		return "other"
	}
}

var Data hc.Locale = Locale{
	localeCode: language.Ukrainian,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 тисяча",
				"1000-count-few":              "0 тисячі",
				"1000-count-many":             "0 тисяч",
				"1000-count-other":            "0 тисячі",
				"10000-count-one":             "00 тисяча",
				"10000-count-few":             "00 тисячі",
				"10000-count-many":            "00 тисяч",
				"10000-count-other":           "00 тисячі",
				"100000-count-one":            "000 тисяча",
				"100000-count-few":            "000 тисячі",
				"100000-count-many":           "000 тисяч",
				"100000-count-other":          "000 тисячі",
				"1000000-count-one":           "0 мільйон",
				"1000000-count-few":           "0 мільйони",
				"1000000-count-many":          "0 мільйонів",
				"1000000-count-other":         "0 мільйона",
				"10000000-count-one":          "00 мільйон",
				"10000000-count-few":          "00 мільйони",
				"10000000-count-many":         "00 мільйонів",
				"10000000-count-other":        "00 мільйона",
				"100000000-count-one":         "000 мільйон",
				"100000000-count-few":         "000 мільйони",
				"100000000-count-many":        "000 мільйонів",
				"100000000-count-other":       "000 мільйона",
				"1000000000-count-one":        "0 мільярд",
				"1000000000-count-few":        "0 мільярди",
				"1000000000-count-many":       "0 мільярдів",
				"1000000000-count-other":      "0 мільярда",
				"10000000000-count-one":       "00 мільярд",
				"10000000000-count-few":       "00 мільярди",
				"10000000000-count-many":      "00 мільярдів",
				"10000000000-count-other":     "00 мільярда",
				"100000000000-count-one":      "000 мільярд",
				"100000000000-count-few":      "000 мільярди",
				"100000000000-count-many":     "000 мільярдів",
				"100000000000-count-other":    "000 мільярда",
				"1000000000000-count-one":     "0 трильйон",
				"1000000000000-count-few":     "0 трильйони",
				"1000000000000-count-many":    "0 трильйонів",
				"1000000000000-count-other":   "0 трильйона",
				"10000000000000-count-one":    "00 трильйон",
				"10000000000000-count-few":    "00 трильйони",
				"10000000000000-count-many":   "00 трильйонів",
				"10000000000000-count-other":  "00 трильйона",
				"100000000000000-count-one":   "000 трильйон",
				"100000000000000-count-few":   "000 трильйони",
				"100000000000000-count-many":  "000 трильйонів",
				"100000000000000-count-other": "000 трильйона",
			},
		},

		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 тис.",
				"1000-count-other":            "0 тис.",
				"10000-count-one":             "00 тис.",
				"10000-count-other":           "00 тис.",
				"100000-count-one":            "000 тис.",
				"100000-count-other":          "000 тис.",
				"1000000-count-one":           "0 млн",
				"1000000-count-other":         "0 млн",
				"10000000-count-one":          "00 млн",
				"10000000-count-other":        "00 млн",
				"100000000-count-one":         "000 млн",
				"100000000-count-other":       "000 млн",
				"1000000000-count-one":        "0 млрд",
				"1000000000-count-other":      "0 млрд",
				"10000000000-count-one":       "00 млрд",
				"10000000000-count-other":     "00 млрд",
				"100000000000-count-one":      "000 млрд",
				"100000000000-count-other":    "000 млрд",
				"1000000000000-count-one":     "0 трлн",
				"1000000000000-count-other":   "0 трлн",
				"10000000000000-count-one":    "00 трлн",
				"10000000000000-count-other":  "00 трлн",
				"100000000000000-count-one":   "000 трлн",
				"100000000000000-count-other": "000 трлн",
			},
		},
	},
}
