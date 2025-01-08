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

// "pluralRule-count-one": "i = 1 and v = 0 or i = 0 and v != 0 @integer 1 @decimal 0.0~0.9, 0.00~0.05",
// "pluralRule-count-two": "i = 2 and v = 0 @integer 2",
// "pluralRule-count-other": " @integer 0, 3~17, 100, 1000, 10000, 100000, 1000000, … @decimal 1.0~2.5, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	if !r.IsInt() {
		return "other"
	}

	i64, _, ok := r.Int64(0)
	if !ok {
		return "other"
	}

	switch {
	case i64 == 1:
		return "one"

	case i64 == 2:
		return "two"

	case (i64 < 0 || i64 > 10) && (i64%10 == 0):
		return "many"

	default:
		return "other"
	}
}

var Data hc.Locale = Locale{
	localeCode: language.Hebrew,
	data: hc.CldrData{
		Long: struct {
			DecimalFormat map[string]string
		}{
			DecimalFormat: map[string]string{
				"1000-count-one":              "‏0 אלף",
				"1000-count-two":              "‏0 אלף",
				"1000-count-many":             "‏0 אלף",
				"1000-count-other":            "‏0 אלף",
				"10000-count-one":             "‏00 אלף",
				"10000-count-two":             "‏00 אלף",
				"10000-count-many":            "‏00 אלף",
				"10000-count-other":           "‏00 אלף",
				"100000-count-one":            "‏000 אלף",
				"100000-count-two":            "‏000 אלף",
				"100000-count-many":           "‏000 אלף",
				"100000-count-other":          "‏000 אלף",
				"1000000-count-one":           "‏0 מיליון",
				"1000000-count-two":           "‏0 מיליון",
				"1000000-count-many":          "‏0 מיליון",
				"1000000-count-other":         "‏0 מיליון",
				"10000000-count-one":          "‏00 מיליון",
				"10000000-count-two":          "‏00 מיליון",
				"10000000-count-many":         "‏00 מיליון",
				"10000000-count-other":        "‏00 מיליון",
				"100000000-count-one":         "‏000 מיליון",
				"100000000-count-two":         "‏000 מיליון",
				"100000000-count-many":        "‏000 מיליון",
				"100000000-count-other":       "‏000 מיליון",
				"1000000000-count-one":        "‏0 מיליארד",
				"1000000000-count-two":        "‏0 מיליארד",
				"1000000000-count-many":       "‏0 מיליארד",
				"1000000000-count-other":      "‏0 מיליארד",
				"10000000000-count-one":       "‏00 מיליארד",
				"10000000000-count-two":       "‏00 מיליארד",
				"10000000000-count-many":      "‏00 מיליארד",
				"10000000000-count-other":     "‏00 מיליארד",
				"100000000000-count-one":      "‏000 מיליארד",
				"100000000000-count-two":      "‏000 מיליארד",
				"100000000000-count-many":     "‏000 מיליארד",
				"100000000000-count-other":    "‏000 מיליארד",
				"1000000000000-count-one":     "‏0 טריליון",
				"1000000000000-count-two":     "‏0 טריליון",
				"1000000000000-count-many":    "‏0 טריליון",
				"1000000000000-count-other":   "‏0 טריליון",
				"10000000000000-count-one":    "‏00 טריליון",
				"10000000000000-count-two":    "‏00 טריליון",
				"10000000000000-count-many":   "‏00 טריליון",
				"10000000000000-count-other":  "‏00 טריליון",
				"100000000000000-count-one":   "‏000 טריליון",
				"100000000000000-count-two":   "‏000 טריליון",
				"100000000000000-count-many":  "‏000 טריליון",
				"100000000000000-count-other": "‏000 טריליון",
			},
		},
		Short: struct {
			DecimalFormat map[string]string
		}{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0K‏",
				"1000-count-other":            "0K‏",
				"10000-count-one":             "00K‏",
				"10000-count-other":           "00K‏",
				"100000-count-one":            "000K‏",
				"100000-count-other":          "000K‏",
				"1000000-count-one":           "0M‏",
				"1000000-count-other":         "0M‏",
				"10000000-count-one":          "00M‏",
				"10000000-count-other":        "00M‏",
				"100000000-count-one":         "000M‏",
				"100000000-count-other":       "000M‏",
				"1000000000-count-one":        "0B‏",
				"1000000000-count-other":      "0B‏",
				"10000000000-count-one":       "00B‏",
				"10000000000-count-other":     "00B‏",
				"100000000000-count-one":      "000B‏",
				"100000000000-count-other":    "000B‏",
				"1000000000000-count-one":     "0T‏",
				"1000000000000-count-other":   "0T‏",
				"10000000000000-count-one":    "00T‏",
				"10000000000000-count-other":  "00T‏",
				"100000000000000-count-one":   "000T‏",
				"100000000000000-count-other": "000T‏",
			},
		},
	},
}
