package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"

	"github.com/yurytaranau/adventofcode-2022/days"
)

const (
	partOne = "Day%dp1"
	partTwo = "Day%dp2"
	source  = "sources/day%d"
)

func main() {
	d := days.Days{}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("incorrect day number provided")
	}
	if day < 1 || day > 31 {
		panic("day number must be in range (1..31)")
	}
	fmt.Printf("Day %d called\n", day)

	// prepare source data
	src, err := os.Open(fmt.Sprintf(source, day))
	if err != nil {
		panic(err)
	}
	defer src.Close()

	buf := &bytes.Buffer{}
	tee := io.TeeReader(src, buf)

	m1 := reflect.ValueOf(d).MethodByName(fmt.Sprintf(partOne, day))
	if m1.IsValid() {
		args1 := []reflect.Value{reflect.ValueOf(tee)}
		m1.Call(args1)
	} else {
		fmt.Printf("method %s not found\n", m1)
	}

	m2 := reflect.ValueOf(d).MethodByName(fmt.Sprintf(partTwo, day))
	if m1.IsValid() {
		args2 := []reflect.Value{reflect.ValueOf(buf)}
		m2.Call(args2)
	} else {
		fmt.Printf("method %s not found\n", m2)
	}

}
