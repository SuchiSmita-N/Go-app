package main

import "testing"

func Testsum(t *testing.T) {
	nos := [5]int{1,2,3,4,5}

	get := sum(nos)
	result:= 15

    if result != get {
		t.Errorf ("get %d expected %d, Given %v ", get, result, nos)
    }
	
}

func sum(nos [5]int) int {
	sum := 0
	for _,nos := range nos{
		sum += nos
	}
    return sum
}
