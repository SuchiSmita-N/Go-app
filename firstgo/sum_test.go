package main

import "testing"


func TestSum(t *testing.T) {
	nos := [5]int{1,2,3,4,5}

	get := Sum(nos)
	result:= 15

    if result != get {
		t.Errorf ("get %d expected %d, Given %v ", get, result, nos)
    }
	
}
func Sum(nos [5]int) int {
	sum := 0
	for i := 0; i<5; i++ {
		sum += nos[i]
	}
    return sum
}
