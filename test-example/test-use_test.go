package example

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T){
	arr := []int{4,5,76,2}
	ret := Min(arr)
	fmt.Println(ret)
}