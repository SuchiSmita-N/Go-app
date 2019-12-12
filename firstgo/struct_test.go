package main

import "testing"

type rectangle struct {
	width float64
	height float64
}

func (r Rectangle) Area() float64  {
    return 0
}

func Testperimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := perimeter(rectangle)
	expect := 40.0	

	if result != expect {
		t.Errorf("result %.2f expect %.2f", result, expect)
	}
}

func perimeter(rectangle Rectangle) float64 {
    return 2 * (rectangle.width + rectangle.height)
}
