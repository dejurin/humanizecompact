package main

import (
	"fmt"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/en"
)

func main() {
    humanizer := hc.New(
        locale.Data, 
        hc.Long,
        func(original string) string {
            return original
        },
    )

    values := []string{"999", "1000", "10000", "1500000", "0.75", "1000000000000"}

    for _, val := range values {
        out, err := humanizer.Humanize(val)
        if err != nil {
            fmt.Printf("[ERROR] val=%q: %v\n", val, err)
            continue
        }
        fmt.Printf("val=%-15s => %s\n", val, out)
    }
}