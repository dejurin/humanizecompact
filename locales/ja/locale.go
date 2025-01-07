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
	localeCode: "ja",
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-other":                 "0",
				"10000-count-other":                "0万",
				"100000-count-other":               "00万",
				"1000000-count-other":              "000万",
				"10000000-count-other":             "0000万",
				"100000000-count-other":            "0億",
				"1000000000-count-other":           "00億",
				"10000000000-count-other":          "000億",
				"100000000000-count-other":         "0000億",
				"1000000000000-count-other":        "0兆",
				"10000000000000-count-other":       "00兆",
				"100000000000000-count-other":      "000兆",
				"1000000000000000-count-other":     "0000兆",
				"10000000000000000-count-other":    "0京",
				"100000000000000000-count-other":   "00京",
				"1000000000000000000-count-other":  "000京",
				"10000000000000000000-count-other": "0000京",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-other":                 "0",
				"10000-count-other":                "0万",
				"100000-count-other":               "00万",
				"1000000-count-other":              "000万",
				"10000000-count-other":             "0000万",
				"100000000-count-other":            "0億",
				"1000000000-count-other":           "00億",
				"10000000000-count-other":          "000億",
				"100000000000-count-other":         "0000億",
				"1000000000000-count-other":        "0兆",
				"10000000000000-count-other":       "00兆",
				"100000000000000-count-other":      "000兆",
				"1000000000000000-count-other":     "0000兆",
				"10000000000000000-count-other":    "0京",
				"100000000000000000-count-other":   "00京",
				"1000000000000000000-count-other":  "000京",
				"10000000000000000000-count-other": "0000京",
			},
		},
	},
}
