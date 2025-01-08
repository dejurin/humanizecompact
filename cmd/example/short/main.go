package main

import (
	"fmt"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/es"
	"golang.org/x/text/language"
)

func main() {
	var locales = map[language.Tag]hc.Locale{
		language.Spanish: locale.Data,
	}
	humanizer := hc.New(
		locales,
		hc.Short,
		func(original string) string {
			return original
		},
	)

	values := []string{"9900000000000"}

	for _, val := range values {
		out, _, err := humanizer.Formatter(val, language.Spanish)
		if err != nil {
			fmt.Printf("[ERROR] val=%q: %v\n", val, err)
			continue
		}
		fmt.Printf("val=%-15s => %s\n", val, out)
	}
}
