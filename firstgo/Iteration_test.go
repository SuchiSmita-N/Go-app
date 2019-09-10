package iterarion

import "testing"

func Testiterate(t *testing.T) {
	iteration := iterate("s")
	result := "sssss"

	if iteration != result {
		t.Errorf("expected '%s' but result is '%s'", result, iteration)
	}
}

func iterate(character string) string {
    var iterated string
    for i := 0; i < 5; i++ {
    	iterated += character
    }
    return iterated
}