package integers	

import "testing"

func Add(x, y int) int {
    return x+y
}

func addtest(t *testing.T) {
	sum := Add(3,4)
	result := 7

  if sum !=	result {
  	t.Errorf("expected '%d' but result is '%d'", result,sum)
  }
}