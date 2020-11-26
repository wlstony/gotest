package main

import (
	"fmt"
	"reflect"
)

//type Person struct {
//	Age  int
//	name string
//}
//
//func main() {
//	p := Person{
//		Age:  1210,
//		name: "aa",
//	}
//
//	//fmt.Println("TypeOf author:", reflect.TypeOf(p).Field(0).Name)
//	//fmt.Println("ValueOf author:", reflect.ValueOf(p))
//	v := reflect.ValueOf(&p)
//	s := v.Elem()
//	fmt.Println("s.Kind():", s.Kind())
//	if s.Kind() == reflect.Struct {
//		field := s.FieldByName("Age")
//		if field.IsValid() {
//			fmt.Println("field.CanSet():", field.CanSet())
//			if field.CanSet() {
//				if field.Kind() == reflect.Int {
//					fmt.Println("set")
//					field.SetInt(12)
//				}
//			}
//		}
//	}
//	fmt.Printf("p%+v\n", p)
//	fmt.Printf("v%+v\n", v)
//	fmt.Printf("s%+v\n", s)
//}


func Add(a, b int) (int, error) { return a + b, nil }

func main() {
	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	argv := make([]reflect.Value, t.NumIn())
	for k := range argv {
		if t.In(k).Kind() != reflect.Int {
			return
		}
		argv[k] = reflect.ValueOf(k * 10)
	}
	fmt.Println(argv)
	result := v.Call(argv)
	if len(result) != 2 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println("result:", result[0].Int()) // #=> 1
}
