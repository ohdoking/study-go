/**
	
	how to callback in go lang

	@author ohdoking
	
	@refer : https://stackoverflow.com/questions/12655464/can-functions-be-passed-as-parameters-in-go
**/

package main

import (
	"fmt"
	"strconv"
)

type convert func(int) int //interface 

func main (){
	// how to callback in go lang

	//first way
	fmt.Println(test(3, func (i int) int{
		return i*i
	}))

	//second way
	fmt.Println(test(4, square))

	//third way
	fmt.Println(plusX(4)(3))

	//fourth way
	x := func (i int) int{return i*i}
	fmt.Println(test(5, x))
}

func test (i int, fn convert) string{
	j := fn(i)
	return strconv.Itoa(j) + " is square of "+ strconv.Itoa(i)
}

func square(i int) int{
	return i*i
}

func plusX(x int) (func(v int) (int)) {
   return func(v int) (int) {
       return v+x
   }
}

