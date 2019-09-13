package main 

import "testing"
import "reflect"

func Testsumall(t *testing.T) {
	get := sumall([]int{1,2,3}, []int{1,2})
	result := []int{6,3}

	if !reflect.DeepEqual(get, result){
		t.Errorf("Get %v result %v",get, result)
	}
}

func sumall(nostosum ...[]int) []int {
	 lenofnos := len(nostosum)
	 sums := make([]int, lenofnos)

	 for i, s := range nostosum{
	 	sums[i] = sum(s)
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

