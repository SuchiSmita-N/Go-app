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


func sumall(nos ...[]int) (sums []int) {
	    return 
}
