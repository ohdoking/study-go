package main

import (
	"fmt"
	"callback2/lib"
)


type Data struct{

}

func (d *Data)Callback(a, b int) int{
	return a+b;
}

func (d *Data)Init() int{
	return common.Init(10, d)
}

func main(){
	a := new(Data)
	fmt.Println(a.Init())
}

