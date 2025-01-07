package humanizecompact_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	hc "github.com/dejurin/humanizecompact"
	locale "github.com/dejurin/humanizecompact/locales/en"
)

// generateTestCases создает массив строковых чисел для тестирования.
func generateTestCases() []string {
	var testCases []string

	for i := -1000; i <= 1000; i++ {
		testCases = append(testCases, fmt.Sprintf("%d", i))
	}

	for i := -1000; i <= 1000; i++ {
		testCases = append(testCases, fmt.Sprintf("%.6f", float64(i)/10))
	}

	for exp := 1; exp <= 1000; exp++ {
		value := int64(math.Pow10(exp))
		testCases = append(testCases, fmt.Sprintf("%d", value))
		testCases = append(testCases, fmt.Sprintf("%d", -value))
	}

	for i := 0; i < 100000; i++ {
		randomValue := rand.Float64() * math.Pow10(rand.Intn(20))
		testCases = append(testCases, fmt.Sprintf("%.6f", randomValue))
	}

	return testCases
}

func BenchmarkHumanizeGenerated(b *testing.B) {
	h := hc.New(locale.Data, hc.Short, func(s string) string {
		return s
	})

	inputs := generateTestCases()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			_, _ = h.Humanize(input)
		}
	}
}
