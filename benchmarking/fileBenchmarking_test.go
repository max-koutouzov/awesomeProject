package main

import (
	"fmt"
	"testing"
)

func TestFileLen(t *testing.T) {
	result, err := FileLen("/tmp/cover282191582/coverage.html", 1024)
	if err != nil {
		t.Fatal(err)
	}
	if result < 500 {
		t.Error("Expected > 500, got: ", result)
	}
}

var blackhole int
var blackhole2 int

func BenchmarkFileLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		restult, err := FileLen("/tmp/cover282191582/coverage.html", 1)
		if err != nil {
			b.Fatal(err)
		}
		blackhole = restult
	}
}

func BenchmarkFileLen2(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen("/tmp/cover282191582/coverage.html", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole2 = result
			}
		})
	}
}
