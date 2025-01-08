package main

import (
	"fmt"

	"golang.org/x/text/language"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/he"
)

func main() {
	var locales = map[language.Tag]hc.Locale{
		language.Hebrew: locale.Data,
	}
	humanizer := hc.New(
		locales,
		hc.Long,
		func(original string) string {
			return original
		},
	)

	values := []string{"999", "1000", "10000", "1500000", "0.75", "1000000000000"}

	for _, val := range values {
		out, _, err := humanizer.Formatter(val, language.Hebrew)
		if err != nil {
			fmt.Printf("[ERROR] val=%q: %v\n", val, err)
			continue
		}
		fmt.Printf("val=%-15s => %s\n", val, out)
	}
}
