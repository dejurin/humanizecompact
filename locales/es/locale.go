package locale

import (
	hc "github.com/dejurin/humanizecompact"
	"github.com/govalues/decimal"
)

type Locale struct {
	data       hc.CldrData
	localeCode string
}

func (l Locale) Data() hc.CldrData {
	return l.data
}

func (l Locale) Code() string {
	return l.localeCode
}

// "pluralRule-count-one": "n = 1 @integer 1 @decimal 1.0, 1.00, 1.000, 1.0000",
// "pluralRule-count-many": "e = 0 and i != 0 and i % 1000000 = 0 and v = 0 or e != 0..5 @integer 1000000, 1c6, 2c6, 3c6, 4c6, 5c6, 6c6, … @decimal 1.0000001c6, 1.1c6, 2.0000001c6, 2.1c6, 3.0000001c6, 3.1c6, …",
// "pluralRule-count-other": " @integer 0, 2~16, 100, 1000, 10000, 100000, 1c3, 2c3, 3c3, 4c3, 5c3, 6c3, … @decimal 0.0~0.9, 1.1~1.6, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, 1.0001c3, 1.1c3, 2.0001c3, 2.1c3, 3.0001c3, 3.1c3, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	one, err := decimal.New(1, 0)
	if err != nil {
		return "other"
	}

	if r.Equal(one) {
		return "one"
	}

	if r.IsInt() {
		i64, scale, ok := r.Int64(0)
		if ok && scale == 0 && i64 != 0 && (i64%1_000_000 == 0) {
			return "many"
		}
	}

	absVal := r.Abs()

	// >= 1e6
	val, err := decimal.New(1, 6)
	if err == nil && absVal.Cmp(val) >= 0 {
		return "many"
	}
	val, err = decimal.New(1, -6)
	if err == nil && absVal.Sign() != 0 && absVal.Cmp(val) <= 0 {
		return "many"
	}

	return "other"
}

var Data hc.Locale = Locale{
	localeCode: "es",
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 mil",
				"1000-count-other":            "0 mil",
				"10000-count-one":             "00 mil",
				"10000-count-other":           "00 mil",
				"100000-count-one":            "000 mil",
				"100000-count-other":          "000 mil",
				"1000000-count-one":           "0 millón",
				"1000000-count-other":         "0 millones",
				"10000000-count-one":          "00 millones",
				"10000000-count-other":        "00 millones",
				"100000000-count-one":         "000 millones",
				"100000000-count-other":       "000 millones",
				"1000000000-count-one":        "0 mil millones",
				"1000000000-count-other":      "0 mil millones",
				"10000000000-count-one":       "00 mil millones",
				"10000000000-count-other":     "00 mil millones",
				"100000000000-count-one":      "000 mil millones",
				"100000000000-count-other":    "000 mil millones",
				"1000000000000-count-one":     "0 billón",
				"1000000000000-count-other":   "0 billones",
				"10000000000000-count-one":    "00 billones",
				"10000000000000-count-other":  "00 billones",
				"100000000000000-count-one":   "000 billones",
				"100000000000000-count-other": "000 billones",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 mil",
				"1000-count-other":            "0 mil",
				"10000-count-one":             "00 mil",
				"10000-count-other":           "00 mil",
				"100000-count-one":            "000 mil",
				"100000-count-other":          "000 mil",
				"1000000-count-one":           "0 M",
				"1000000-count-other":         "0 M",
				"10000000-count-one":          "00 M",
				"10000000-count-other":        "00 M",
				"100000000-count-one":         "000 M",
				"100000000-count-other":       "000 M",
				"1000000000-count-one":        "0000 M",
				"1000000000-count-other":      "0000 M",
				"10000000000-count-one":       "00 mil M",
				"10000000000-count-other":     "00 mil M",
				"100000000000-count-one":      "000 mil M",
				"100000000000-count-other":    "000 mil M",
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
