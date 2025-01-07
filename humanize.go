package humanize

import (
	"fmt"
	"sort"
	"strings"

	"github.com/govalues/decimal"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Locale interface {
	Data() CldrData
	PluralForm(r decimal.Decimal, v string) string
	Code() string
}

type CldrData struct {
	Long struct {
		DecimalFormat map[string]string
	}
	Short struct {
		DecimalFormat map[string]string
	}
}

type Option int

const (
	OptionLong Option = iota
	OptionShort
)

type FallbackFunc func(original string) string

type InvalidNumberError struct {
	Value string
	Err   error
}

func (l InvalidNumberError) Error() string {
	return fmt.Sprintf("invalid number %q: %v", l.Value, l.Err)
}

type Humanizer struct {
	locale   Locale
	opt      Option
	fallback FallbackFunc
}

func NewHumanizer(loc Locale, opt Option, fb FallbackFunc) *Humanizer {
	return &Humanizer{
		locale:   loc,
		opt:      opt,
		fallback: fb,
	}
}

func (h *Humanizer) Humanize(value string) (string, error) {
	valDec, err := decimal.Parse(value)
	if err != nil {
		return "", InvalidNumberError{Value: value, Err: err}
	}
	if !valDec.IsInt() {
		return h.fallback(value), nil
	}

	var df map[string]string
	if h.opt == OptionLong {
		df = h.locale.Data().Long.DecimalFormat
	} else {
		df = h.locale.Data().Short.DecimalFormat
	}
	if len(df) == 0 {
		return h.fallback(value), nil
	}

	groupScales := parseGroupScales(df)
	if len(groupScales) == 0 {
		return h.fallback(value), nil
	}
	sortedGroupScales := sortGroupScales(groupScales)
	var best groupScale
	var bestRatio decimal.Decimal

	one, _ := decimal.New(1, 0)
	thousand, _ := decimal.New(1000, 0)

	for _, gs := range sortedGroupScales {
		scaleDec, _ := decimal.New(gs.scale, 0)
		if valDec.Cmp(scaleDec) >= 0 {
			ratio, divErr := valDec.Quo(scaleDec)
			if divErr != nil {
				continue
			}
			if ratio.Cmp(one) < 0 {
				continue
			}
			if ratio.Cmp(thousand) > 0 {
				continue
			}
			if !has0or1DecimalExactly(ratio, h.locale.Code()) {
				continue
			}
			if !isAllowedRatio(ratio) {
				continue
			}
			if bestRatio.IsZero() || ratio.Cmp(bestRatio) == -1 {
				best = gs
				bestRatio = ratio
			}
		}
	}
	if bestRatio.IsZero() {
		return h.fallback(value), nil
	}

	pluralForm := h.locale.PluralForm(bestRatio, value)
	key := fmt.Sprintf("%d-count-%s", best.scale, pluralForm)
	tmpl := df[key]

	if tmpl == "" && pluralForm != "other" {
		altKey := fmt.Sprintf("%d-count-other", best.scale)
		altTmpl := df[altKey]
		if altTmpl != "" {
			tmpl = altTmpl
		}
	}

	tag := language.MustParse(h.locale.Code())
	p := message.NewPrinter(tag)
	floatVal, _ := bestRatio.Float64()

	return replacePlaceholder(tmpl, p.Sprintf("%v", floatVal)), nil
}

type groupScale struct {
	name  string
	scale int64
}

func parseGroupScales(df map[string]string) map[string]decimal.Decimal {
	nameToMin := make(map[string]decimal.Decimal)
	for k, tmpl := range df {
		parts := strings.Split(k, "-count-")
		if len(parts) != 2 {
			continue
		}
		scaleStr := parts[0]
		scaleVal := decimal.MustParse(scaleStr)

		gname := extractName(tmpl)
		if gname == "" {
			continue
		}

		old, ok := nameToMin[gname]
		if !ok || scaleVal.Cmp(old) == -1 {
			nameToMin[gname] = scaleVal
		}
	}
	return nameToMin
}

func sortGroupScales(m map[string]decimal.Decimal) []groupScale {
	out := make([]groupScale, 0, len(m))
	for n, sc := range m {
		i64, _, _ := sc.Int64(0)
		out = append(out, groupScale{name: n, scale: i64})
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].scale > out[j].scale
	})
	return out
}

func extractName(s string) string {
	return strings.ReplaceAll(s, "0", "")
}

func has0or1DecimalExactly(d decimal.Decimal, localeCode string) bool {
	// TODO: find a better way to handle this
	if localeCode != "ja" && localeCode != "ko" {
		t1 := d.Trunc(1)
		if d.Equal(t1) {
			return true
		}
		r1 := d.Round(1)
		return d.Equal(r1)
	}

	r2 := d.Round(2)
	return d.Equal(r2)
}

func isAllowedRatio(r decimal.Decimal) bool {
	if r.IsInt() {
		return true
	}
	floor := r.Floor(0)
	hundred, _ := decimal.New(100, 0)
	if floor.Cmp(hundred) >= 0 {
		return false
	}
	return true
}

func replacePlaceholder(tmpl, ratio string) string {
	pats := []string{"000", "00", "0"}
	for _, p := range pats {
		if idx := strings.Index(tmpl, p); idx >= 0 {
			return tmpl[:idx] + ratio + tmpl[idx+len(p):]
		}
	}
	return tmpl
}
