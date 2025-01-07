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
	localeCode: language.Russian,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 тысяча",
				"1000-count-few":              "0 тысячи",
				"1000-count-many":             "0 тысяч",
				"1000-count-other":            "0 тысячи",
				"10000-count-one":             "00 тысяча",
				"10000-count-few":             "00 тысячи",
				"10000-count-many":            "00 тысяч",
				"10000-count-other":           "00 тысячи",
				"100000-count-one":            "000 тысяча",
				"100000-count-few":            "000 тысячи",
				"100000-count-many":           "000 тысяч",
				"100000-count-other":          "000 тысячи",
				"1000000-count-one":           "0 миллион",
				"1000000-count-few":           "0 миллиона",
				"1000000-count-many":          "0 миллионов",
				"1000000-count-other":         "0 миллиона",
				"10000000-count-one":          "00 миллион",
				"10000000-count-few":          "00 миллиона",
				"10000000-count-many":         "00 миллионов",
				"10000000-count-other":        "00 миллиона",
				"100000000-count-one":         "000 миллион",
				"100000000-count-few":         "000 миллиона",
				"100000000-count-many":        "000 миллионов",
				"100000000-count-other":       "000 миллиона",
				"1000000000-count-one":        "0 миллиард",
				"1000000000-count-few":        "0 миллиарда",
				"1000000000-count-many":       "0 миллиардов",
				"1000000000-count-other":      "0 миллиарда",
				"10000000000-count-one":       "00 миллиард",
				"10000000000-count-few":       "00 миллиарда",
				"10000000000-count-many":      "00 миллиардов",
				"10000000000-count-other":     "00 миллиарда",
				"100000000000-count-one":      "000 миллиард",
				"100000000000-count-few":      "000 миллиарда",
				"100000000000-count-many":     "000 миллиардов",
				"100000000000-count-other":    "000 миллиарда",
				"1000000000000-count-one":     "0 триллион",
				"1000000000000-count-few":     "0 триллиона",
				"1000000000000-count-many":    "0 триллионов",
				"1000000000000-count-other":   "0 триллиона",
				"10000000000000-count-one":    "00 триллион",
				"10000000000000-count-few":    "00 триллиона",
				"10000000000000-count-many":   "00 триллионов",
				"10000000000000-count-other":  "00 триллиона",
				"100000000000000-count-one":   "000 триллион",
				"100000000000000-count-few":   "000 триллиона",
				"100000000000000-count-many":  "000 триллионов",
				"100000000000000-count-other": "000 триллиона",
			},
		},

		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 тыс.",
				"1000-count-other":            "0 тыс.",
				"10000-count-one":             "00 тыс.",
				"10000-count-other":           "00 тыс.",
				"100000-count-one":            "000 тыс.",
				"100000-count-other":          "000 тыс.",
				"1000000-count-one":           "0 млн",
				"1000000-count-other":         "0 млн",
				"10000000-count-one":          "00 млн",
				"10000000-count-other":        "00 млн",
				"100000000-count-one":         "000 млн",
				"100000000-count-other":       "000 млн",
				"1000000000-count-one":        "0 млрд",
				"1000000000-count-other":      "0 млрд",
				"10000000000-count-one":       "00 млрд",
				"10000000000-count-other":     "00 млрд",
				"100000000000-count-one":      "000 млрд",
				"100000000000-count-other":    "000 млрд",
				"1000000000000-count-one":     "0 трлн",
				"1000000000000-count-other":   "0 трлн",
				"10000000000000-count-one":    "00 трлн",
				"10000000000000-count-other":  "00 трлн",
				"100000000000000-count-one":   "000 трлн",
				"100000000000000-count-other": "000 трлн",
			},
		},
	},
}
