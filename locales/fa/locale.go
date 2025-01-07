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

// "pluralRule-count-one": "i = 0 or n = 1 @integer 0, 1 @decimal 0.0~1.0, 0.00~0.04",
// "pluralRule-count-other": " @integer 2~17, 100, 1000, 10000, 100000, 1000000, … @decimal 1.1~2.6, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	one, err := decimal.NewFromInt64(1, 0, 0)
	if err != nil {
		return "other"
	}

	if r.Equal(one) {
		return "one"
	}

	i64, _, ok := r.Int64(0)
	if !ok {
		return "other"
	}

	if i64 == 0 {
		return "one"
	}

	return "other"
}

var Data hc.Locale = Locale{
	localeCode: "fa",
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 هزار",
				"1000-count-other":            "0 هزار",
				"10000-count-one":             "00 هزار",
				"10000-count-other":           "00 هزار",
				"100000-count-one":            "000 هزار",
				"100000-count-other":          "000 هزار",
				"1000000-count-one":           "0 میلیون",
				"1000000-count-other":         "0 میلیون",
				"10000000-count-one":          "00 میلیون",
				"10000000-count-other":        "00 میلیون",
				"100000000-count-one":         "000 میلیون",
				"100000000-count-other":       "000 میلیون",
				"1000000000-count-one":        "0 میلیارد",
				"1000000000-count-other":      "0 میلیارد",
				"10000000000-count-one":       "00 میلیارد",
				"10000000000-count-other":     "00 میلیارد",
				"100000000000-count-one":      "000 میلیارد",
				"100000000000-count-other":    "000 میلیارد",
				"1000000000000-count-one":     "0 هزارمیلیارد",
				"1000000000000-count-other":   "0 هزارمیلیارد",
				"10000000000000-count-one":    "00 هزارمیلیارد",
				"10000000000000-count-other":  "00 هزارمیلیارد",
				"100000000000000-count-one":   "000 هزارمیلیارد",
				"100000000000000-count-other": "000 هزارمیلیارد",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-one":              "0 هزار",
				"1000-count-other":            "0 هزار",
				"10000-count-one":             "00 هزار",
				"10000-count-other":           "00 هزار",
				"100000-count-one":            "000 هزار",
				"100000-count-other":          "000 هزار",
				"1000000-count-one":           "0 میلیون",
				"1000000-count-other":         "0 میلیون",
				"10000000-count-one":          "00 میلیون",
				"10000000-count-other":        "00 میلیون",
				"100000000-count-one":         "000 میلیون",
				"100000000-count-other":       "000 میلیون",
				"1000000000-count-one":        "0 میلیارد",
				"1000000000-count-other":      "0 میلیارد",
				"10000000000-count-one":       "00 میلیارد",
				"10000000000-count-other":     "00 میلیارد",
				"100000000000-count-one":      "000 میلیارد",
				"100000000000-count-other":    "000 میلیارد",
				"1000000000000-count-one":     "0 تریلیون",
				"1000000000000-count-other":   "0 تریلیون",
				"10000000000000-count-one":    "00 تریلیون",
				"10000000000000-count-other":  "00 تریلیون",
				"100000000000000-count-one":   "000 تریلیون",
				"100000000000000-count-other": "000 تریلیون",
			},
		},
	},
}
