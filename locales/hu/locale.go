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

// "pluralRule-count-one": "n = 1 @integer 1 @decimal 1.0, 1.00, 1.000, 1.0000",
// "pluralRule-count-other": " @integer 0, 2~16, 100, 1000, 10000, 100000, 1000000, … @decimal 0.0~0.9, 1.1~1.6, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
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
	localeCode: language.Hungarian,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 ezer",
				"1000-count-other":            "0 ezer",
				"10000-count-one":             "00 ezer",
				"10000-count-other":           "00 ezer",
				"100000-count-one":            "000 ezer",
				"100000-count-other":          "000 ezer",
				"1000000-count-one":           "0 millió",
				"1000000-count-other":         "0 millió",
				"10000000-count-one":          "00 millió",
				"10000000-count-other":        "00 millió",
				"100000000-count-one":         "000 millió",
				"100000000-count-other":       "000 millió",
				"1000000000-count-one":        "0 milliárd",
				"1000000000-count-other":      "0 milliárd",
				"10000000000-count-one":       "00 milliárd",
				"10000000000-count-other":     "00 milliárd",
				"100000000000-count-one":      "000 milliárd",
				"100000000000-count-other":    "000 milliárd",
				"1000000000000-count-one":     "0 billió",
				"1000000000000-count-other":   "0 billió",
				"10000000000000-count-one":    "00 billió",
				"10000000000000-count-other":  "00 billió",
				"100000000000000-count-one":   "000 billió",
				"100000000000000-count-other": "000 billió",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 E",
				"1000-count-other":            "0 E",
				"10000-count-one":             "00 E",
				"10000-count-other":           "00 E",
				"100000-count-one":            "000 E",
				"100000-count-other":          "000 E",
				"1000000-count-one":           "0 M",
				"1000000-count-other":         "0 M",
				"10000000-count-one":          "00 M",
				"10000000-count-other":        "00 M",
				"100000000-count-one":         "000 M",
				"100000000-count-other":       "000 M",
				"1000000000-count-one":        "0 Mrd",
				"1000000000-count-other":      "0 Mrd",
				"10000000000-count-one":       "00 Mrd",
				"10000000000-count-other":     "00 Mrd",
				"100000000000-count-one":      "000 Mrd",
				"100000000000-count-other":    "000 Mrd",
				"1000000000000-count-one":     "0 B",
				"1000000000000-count-other":   "0 B",
				"10000000000000-count-one":    "00 B",
				"10000000000000-count-other":  "00 B",
				"100000000000000-count-one":   "000 B",
				"100000000000000-count-other": "000 B",
			},
		},
	},
}
