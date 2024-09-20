package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// Convert string to float
	s := "3.14159"
	f, _ := strconv.ParseFloat(s, 8)
	fmt.Println(f)

	// Convert string to bool
	str := "true"
	b, _ := strconv.ParseBool(str)
	fmt.Println(b)

	// Convert float to string
	var flo float64 = 3.14159
	var strFlo string = strconv.FormatFloat(flo, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(strFlo))

	// Convert float to int
	var f32 float32 = 3.14159
	fmt.Println(reflect.TypeOf(f32))
	var i32 int32 = int32(f32)
	fmt.Println(reflect.TypeOf(i32))
}
