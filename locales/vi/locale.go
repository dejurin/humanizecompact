package locale

import (
	humanize "github.com/dejurin/humanize-cldr"
	"github.com/govalues/decimal"
)

type Locale struct {
	data       humanize.CldrData
	localeCode string
}

func (l Locale) Data() humanize.CldrData {
	return l.data
}

func (l Locale) Code() string {
	return l.localeCode
}

// "pluralRule-count-other": " @integer 0~15, 100, 1000, 10000, 100000, 1000000, … @decimal 0.0~1.5, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	return "other"
}

var Data humanize.Locale = Locale{
	localeCode: "vi",
	data: humanize.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-other":            "0 nghìn",
				"10000-count-other":           "00 nghìn",
				"100000-count-other":          "000 nghìn",
				"1000000-count-other":         "0 triệu",
				"10000000-count-other":        "00 triệu",
				"100000000-count-other":       "000 triệu",
				"1000000000-count-other":      "0 tỷ",
				"10000000000-count-other":     "00 tỷ",
				"100000000000-count-other":    "000 tỷ",
				"1000000000000-count-other":   "0 nghìn tỷ",
				"10000000000000-count-other":  "00 nghìn tỷ",
				"100000000000000-count-other": "000 nghìn tỷ",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-other":            "0 N",
				"10000-count-other":           "00 N",
				"100000-count-other":          "000 N",
				"1000000-count-other":         "0 Tr",
				"10000000-count-other":        "00 Tr",
				"100000000-count-other":       "000 Tr",
				"1000000000-count-other":      "0 T",
				"10000000000-count-other":     "00 T",
				"100000000000-count-other":    "000 T",
				"1000000000000-count-other":   "0 NT",
				"10000000000000-count-other":  "00 NT",
				"100000000000000-count-other": "000 NT",
			},
		},
	},
}
