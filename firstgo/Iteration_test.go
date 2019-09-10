package iterarion

import "testing"

func Testiterate(t *testing.T) {
	iteration := iterate("S")
	result := "SSSSS"

	if iteration != result {
		t.Errorf("expected '%s' but result is '%s'", result, iteration)
	}
}

// #Add a Benchmark (To run it go test -bench=".")
func Benchmarkiterate(b *testing.B) {
	for i := 0; i<b.N; i++{
		iterate("a")
	}
	
}