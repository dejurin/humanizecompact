// Package humanize provides functionality for converting numeric strings into
// more human-friendly representations according to CLDR data and locale rules.
package humanizecompact

import (
	"fmt"
	"sort"
	"strings"

	"github.com/govalues/decimal"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Locale defines the methods needed by Humanizer to format numbers
// for a particular locale. Implementations must provide both the
// locale code (e.g., "en", "ja") and CLDR data.
type Locale interface {
	// Data returns locale-specific CLDR formatting data.
	Data() CldrData

	// PluralForm determines the appropriate plural category (e.g. "one",
	// "other") based on the decimal value and a string form of that value.
	PluralForm(r decimal.Decimal, v string) string

	// Code returns a BCP-47 locale tag like "en", "ja", etc.
	Code() language.Tag
}

// CldrData contains the relevant decimal formats for short and long forms
// according to the CLDR specification.
type CldrData struct {
	Long struct {
		DecimalFormat map[string]string
	}
	Short struct {
		DecimalFormat map[string]string
	}
}

// Option indicates whether Humanizer should use long or short
// CLDR patterns (e.g., "1 thousand" vs. "1K").
type Option int

const (
	// Long indicates long-form patterns, e.g., "1 thousand".
	Long Option = iota

	// Short indicates short-form patterns, e.g., "1K".
	Short
)

// FallbackFunc is a user-supplied function invoked when the input string
// cannot be humanized (for instance, if the input is not an integer).
type FallbackFunc func(original string) string

// InvalidNumberError is returned when the provided string cannot be parsed
// as a valid number by the decimal library.
type InvalidNumberError struct {
	Value string
	Err   error
}

// Error implements the error interface.
func (e InvalidNumberError) Error() string {
	return fmt.Sprintf("invalid number %q: %v", e.Value, e.Err)
}

// Humanizer is responsible for converting a numeric string into a
// more human-friendly representation (e.g., "1K") according to its
// configured Locale, Option, and fallback strategy.
type Humanizer struct {
	locales  map[language.Tag]Locale
	opt      Option
	fallback FallbackFunc
}

// New returns a pointer to a new Humanizer with the given locale,
// form option (long or short), and fallback function. The fallback
// function is called whenever the input string cannot be
// humanized (e.g., non-integer or missing CLDR data).
func New(locales map[language.Tag]Locale, opt Option, fb FallbackFunc) *Humanizer {
	return &Humanizer{
		locales:  locales,
		opt:      opt,
		fallback: fb,
	}
}

// Humanize attempts to produce a locale-appropriate, human-friendly
// version of the given numeric string. If the string cannot be parsed
// as a decimal integer or the available data is insufficient, the
// configured fallback function is used instead.
func (h *Humanizer) Formatter(value string, locale language.Tag) (string, bool, error) {
	locCode := locale
	loc, exists := h.locales[locCode]
	if !exists {
		return "", false, fmt.Errorf("locale %q not found", locCode)
	}

	valDec, err := decimal.Parse(value)
	if err != nil {
		return "", false, InvalidNumberError{Value: value, Err: err}
	}
	if !valDec.IsInt() {
		return h.fallback(value), true, nil
	}

	var df map[string]string
	if h.opt == Long {
		df = loc.Data().Long.DecimalFormat
	} else {
		df = loc.Data().Short.DecimalFormat
	}
	if len(df) == 0 {
		return h.fallback(value), true, nil
	}

	groupScales := parseGroupScales(df)
	if len(groupScales) == 0 {
		return h.fallback(value), true, nil
	}

	sortedScales := sortGroupScales(groupScales)

	one, _ := decimal.New(1, 0)
	thousand, _ := decimal.New(1000, 0)

	var best groupScale
	var bestRatio decimal.Decimal

	for _, gs := range sortedScales {
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
			if !has0or1DecimalExactly(ratio, loc.Code()) {
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
		return h.fallback(value), true, nil
	}

	pluralForm := loc.PluralForm(bestRatio, value)
	key := fmt.Sprintf("%d-count-%s", best.scale, pluralForm)
	tmpl := df[key]

	if tmpl == "" && pluralForm != "other" {
		altKey := fmt.Sprintf("%d-count-other", best.scale)
		altTmpl := df[altKey]
		if altTmpl != "" {
			tmpl = altTmpl
		}
	}

	if tmpl == "" {
		return h.fallback(value), true, nil
	}

	p := message.NewPrinter(locale)
	floatVal, _ := bestRatio.Float64()

	return replacePlaceholder(tmpl, p.Sprintf("%v", floatVal)), false, nil
}

// groupScale is an internal struct for capturing a scale name
// (e.g. "thousand") and its integer-based scale factor (e.g. 1000).
type groupScale struct {
	name  string
	scale int64
}

// cutCountSuffix removes the "-count-" suffix from a key, returning
// the prefix and a boolean indicating whether the suffix was found.
func cutCountSuffix(k string) (prefix string, ok bool) {
	idx := strings.Index(k, "-count-")
	if idx < 0 {
		return k, false
	}
	return k[:idx], true
}

// parseGroupScales generates a map of group scales (e.g. "thousand" => 1000)
// from the CLDR data.
func parseGroupScales(df map[string]string) map[string]decimal.Decimal {
	nameToMin := make(map[string]decimal.Decimal)

	for k, tmpl := range df {
		prefix, found := cutCountSuffix(k)
		if !found {
			continue
		}

		scaleVal, err := decimal.Parse(prefix)
		if err != nil {
			continue
		}

		gname := extractName(tmpl)
		if gname == "" {
			continue
		}

		if old, ok := nameToMin[gname]; !ok || scaleVal.Cmp(old) < 0 {
			nameToMin[gname] = scaleVal
		}
	}

	return nameToMin
}

// sortGroupScales converts the map of group scales to a slice, then sorts
// by descending scale (largest first).
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

// extractName removes '0' placeholders from the template string, leaving
// the abbreviated or symbolic portion to identify the scale name.
func extractName(s string) string {
	return strings.ReplaceAll(s, "0", "")
}

// has0or1DecimalExactly checks if the given decimal has exactly zero
// or one decimal place, depending on certain locale constraints.
func has0or1DecimalExactly(d decimal.Decimal, localeCode language.Tag) bool {
	// Some East Asian locales require slightly different logic.
	if localeCode.String() != "ja" && localeCode.String() != "ko" {
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

// isAllowedRatio returns true if the ratio is within an acceptable range
// for "short" or "long" transformations (i.e. < 100 if not an integer).
func isAllowedRatio(r decimal.Decimal) bool {
	if r.IsInt() {
		return true
	}
	floor := r.Floor(0)
	hundred, _ := decimal.New(100, 0)
	return floor.Cmp(hundred) < 0
}

// replacePlaceholder replaces the first occurrence of "000", "00", or "0"
// in tmpl with ratio, matching common CLDR patterns. If none found,
// tmpl is returned unchanged.
func replacePlaceholder(tmpl, ratio string) string {
	pats := []string{"000", "00", "0"}
	for _, p := range pats {
		if idx := strings.Index(tmpl, p); idx >= 0 {
			return tmpl[:idx] + ratio + tmpl[idx+len(p):]
		}
	}
	return tmpl
}
