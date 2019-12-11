package main

import "testing"

func Testperimeter(t *testing.T) {
	result := perimeter(10.0,10.0)
	expect := 40.0	

	if result != expect {
		t.errorf("result %.2f expect %.2f", result, expect)
	}
}

func TestArea(t *testing.T) {

	result := area(5.0, 6.0)
	expect := 30.0

	if result != expect{
		t.errorf("result %.2f expect %.2f" result, expect)
	}	
}

func perimeter(width float64, height float64) float64 {
    return 2 * (width + height)
}

func area(width float64, height float64) float64 {
    return  (width * height)
}