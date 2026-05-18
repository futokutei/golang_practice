package program

import (
	"testing"
)

// Спільна структура для тестів продуктивності
type BenchmarkUser struct {
	Name    string   `json:"user_name"`
	Age     int      `json:"age"`
	IsAdmin bool     `json:"is_admin"`
	Tags    []string `json:"tags"`
}

// Загальні вхідні дані (запобігає алокаціям під час самого бенчмарку)
var benchInputData = BenchmarkUser{
	Name:    "Олександр",
	Age:     28,
	IsAdmin: true,
	Tags:    []string{"go", "reflect", "benchmark", "json", "yaml"},
}

// Глобальні змінні, щоб уникнути оптимізації компілятора (compiler optimization)
var (
	sinkString string
	sinkBytes  []byte
)

// 1. Бенчмарк для твоєї попередньої функції ToYAML
func BenchmarkToYAML(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := ToYAML(benchInputData)
		if err != nil {
			b.Fatal(err)
		}
		sinkString = res
	}
}

// 2. Бенчмарк для твоєї нової функції ToJSON
func BenchmarkToJSON(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := ToJSON(benchInputData)
		if err != nil {
			b.Fatal(err)
		}
		sinkString = res
	}
}
