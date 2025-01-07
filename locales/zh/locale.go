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

// "pluralRule-count-other": " @integer 0~15, 100, 1000, 10000, 100000, 1000000, … @decimal 0.0~1.5, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	return "other"
}

var Data hc.Locale = Locale{
	localeCode: "zh",
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-other":            "0",
				"10000-count-other":           "0万",
				"100000-count-other":          "00万",
				"1000000-count-other":         "000万",
				"10000000-count-other":        "0000万",
				"100000000-count-other":       "0亿",
				"1000000000-count-other":      "00亿",
				"10000000000-count-other":     "000亿",
				"100000000000-count-other":    "0000亿",
				"1000000000000-count-other":   "0万亿",
				"10000000000000-count-other":  "00万亿",
				"100000000000000-count-other": "000万亿",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-other":            "0",
				"10000-count-other":           "0万",
				"100000-count-other":          "00万",
				"1000000-count-other":         "000万",
				"10000000-count-other":        "0000万",
				"100000000-count-other":       "0亿",
				"1000000000-count-other":      "00亿",
				"10000000000-count-other":     "000亿",
				"100000000000-count-other":    "0000亿",
				"1000000000000-count-other":   "0万亿",
				"10000000000000-count-other":  "00万亿",
				"100000000000000-count-other": "000万亿",
			},
		},
	},
}
