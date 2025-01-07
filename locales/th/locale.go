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

// "pluralRule-count-other": " @integer 0~15, 100, 1000, 10000, 100000, 1000000, … @decimal 0.0~1.5, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, …"
func (l Locale) PluralForm(r decimal.Decimal, v string) string {
	return "other"
}

var Data hc.Locale = Locale{
	localeCode: language.Thai,
	data: hc.CldrData{
		Long: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-other":            "0 พัน",
				"10000-count-other":           "0 หมื่น",
				"100000-count-other":          "0 แสน",
				"1000000-count-other":         "0 ล้าน",
				"10000000-count-other":        "00 ล้าน",
				"100000000-count-other":       "000 ล้าน",
				"1000000000-count-other":      "0 พันล้าน",
				"10000000000-count-other":     "0 หมื่นล้าน",
				"100000000000-count-other":    "0 แสนล้าน",
				"1000000000000-count-other":   "0 ล้านล้าน",
				"10000000000000-count-other":  "00 ล้านล้าน",
				"100000000000000-count-other": "000 ล้านล้าน",
			},
		},
		Short: struct{ DecimalFormat map[string]string }{
			DecimalFormat: map[string]string{
				"1000-count-other":            "0K",
				"10000-count-other":           "00K",
				"100000-count-other":          "000K",
				"1000000-count-other":         "0M",
				"10000000-count-other":        "00M",
				"100000000-count-other":       "000M",
				"1000000000-count-other":      "0B",
				"10000000000-count-other":     "00B",
				"100000000000-count-other":    "000B",
				"1000000000000-count-other":   "0T",
				"10000000000000-count-other":  "00T",
				"100000000000000-count-other": "000T",
			},
		},
	},
}
