package main

import (
	"fmt"
	"inherit/lib"
)

type implement interface{
	people.Common
}

type asian struct{
	people.People
}

type europian struct{
	people.People
}

func (asian) GetString() string{
	return "oh"
}

func (europian) GetString() string{
	return "DO"
}

func main(){

	var p implement
	p = asian{}

	fmt.Println(p.GetString());

	p = europian{}

	fmt.Println(p.GetString());

	s:= make1("ah : %s", "ohdoking")
	fmt.Println(s)

	b:= make1("ah : %s , bh : %s", "ohdoking", "test")
	fmt.Println(b)

	c:= make1("ah")
	fmt.Println(c)


}

func make1(msg string ,params ...interface{}) string{
	return fmt.Sprintf(msg, params...)
}