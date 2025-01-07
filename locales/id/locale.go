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
	localeCode: "id",
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-other":            "0 ribu",
				"10000-count-other":           "00 ribu",
				"100000-count-other":          "000 ribu",
				"1000000-count-other":         "0 juta",
				"10000000-count-other":        "00 juta",
				"100000000-count-other":       "000 juta",
				"1000000000-count-other":      "0 miliar",
				"10000000000-count-other":     "00 miliar",
				"100000000000-count-other":    "000 miliar",
				"1000000000000-count-other":   "0 triliun",
				"10000000000000-count-other":  "00 triliun",
				"100000000000000-count-other": "000 triliun",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-other":            "0 rb",
				"10000-count-other":           "00 rb",
				"100000-count-other":          "000 rb",
				"1000000-count-other":         "0 jt",
				"10000000-count-other":        "00 jt",
				"100000000-count-other":       "000 jt",
				"1000000000-count-other":      "0 M",
				"10000000000-count-other":     "00 M",
				"100000000000-count-other":    "000 M",
				"1000000000000-count-other":   "0 T",
				"10000000000000-count-other":  "00 T",
				"100000000000000-count-other": "000 T",
			},
		},
	},
}
