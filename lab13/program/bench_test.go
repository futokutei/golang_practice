package program

import (
	"testing"
)

type BenchmarkUser struct {
	Name    string   `json:"user_name"`
	Age     int      `json:"age"`
	IsAdmin bool     `json:"is_admin"`
	Tags    []string `json:"tags"`
}

var benchInputData = BenchmarkUser{
	Name:    "Олександр",
	Age:     28,
	IsAdmin: true,
	Tags:    []string{"go", "reflect", "benchmark", "json", "yaml"},
}

var (
	sinkString string
	sinkBytes  []byte
)

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
