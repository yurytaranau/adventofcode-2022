package main

import (
	"fmt"
	"os"
	"reflect"
)

type days struct{}

func main() {
	d := days{}
	day := os.Args[1]
	fmt.Println(day, "called")
	m := reflect.ValueOf(d).MethodByName(day)
	m.Call(nil)
}
