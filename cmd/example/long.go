package main

import (
	"fmt"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/en"
	"golang.org/x/text/language"
)

func main() {
    var locales = map[string]hc.Locale{
        "en": locale.Data,
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
        out, err := humanizer.Formatter(val, language.English)
        if err != nil {
            fmt.Printf("[ERROR] val=%q: %v\n", val, err)
            continue
        }
        fmt.Printf("val=%-15s => %s\n", val, out)
    }
}