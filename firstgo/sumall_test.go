package main 

import "testing"
import "reflect"

func Testsumalltails(t *testing.T) {
	get := sumall([]int{1,2,3}, []int{1,2})
	result := []int{6,3}

	if !reflect.DeepEqual(get, result){
		t.Errorf("Get %v result %v",get, result)
	}
}

func sumall(nostosum ...[]int) []int {
	var sums []int
	for _, s := range nostosum{
		if len(s) == 0{
		sums = append(sums, 0)	
		}
		tail := s[1:]
	 	sums = append(sums, sum(tail))
	 	// append function which takes a slice and a new value, 
	 	// returning a new slice with all the items in it.
	 } 
	 return sums
}
func sum(nos []int) int {
	sum := 0
	for _,nos := range nos{
		sum += nos
	}
    return sum
}

