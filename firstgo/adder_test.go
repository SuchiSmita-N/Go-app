package integers	

import "testing"

#Adding 2 values & writing atest for it

func addtest(t *testing.T) {
	sum := Add(3,4)
	result := 7

  if sum !=	result {
  	t.Errorf("expected '%d' but result is '%d'", result,sum)
  }
}