package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type: ", reflect.TypeOf(x))
	fmt.Println("value: ", reflect.ValueOf(x))
	fmt.Println("----------------------")
	v := reflect.ValueOf(x)
	fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
	fmt.Println("type: ", v.Type())
	fmt.Println("value: ", v.Float())

	//
	var num float64 = 1.2345
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)
	fmt.Println(convertPointer)
	fmt.Println(convertValue)

	//
	p1 := Person{"cxy", 25, "man"}
	DoFileAndMethod(p1)

	//
	var num1 float64 = 1.23456
	fmt.Println("old value of point: ", num1)
	pointer1 := reflect.ValueOf(&num1)
	newValue := pointer1.Elem()
	fmt.Println("type of pointer: ", newValue.Type())
	fmt.Println("settability of pointer: ", newValue.CanSet())
	newValue.SetFloat(77)
	fmt.Println("new value of pointer: ", num1)

	//
	p2 := Person{"cxy", 25, "man"}
	getValue := reflect.ValueOf(p2)
	methodValue1 := getValue.MethodByName("PrintInfo")
	fmt.Printf("Kind: %s, Type: %s\n", methodValue1.Kind(), methodValue1.Type())
	methodValue1.Call(nil)
	arg1 := make([]reflect.Value, 0)
	methodValue1.Call(arg1)
	methodValue2 := getValue.MethodByName("Say")
	fmt.Printf("Kind: %s, Type: %s\n", methodValue2.Kind(), methodValue2.Type())
	arg2 := []reflect.Value{reflect.ValueOf("reflect")}
	methodValue2.Call(arg2)
	methodValue3 := getValue.MethodByName("Test")
	fmt.Printf("Kind: %s, Type: %s\n", methodValue3.Kind(), methodValue3.Type())
	arg3 := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200), reflect.ValueOf("hello")}
	methodValue3.Call(arg3)

	//
	f1 := fun1
	value1 := reflect.ValueOf(f1)
	fmt.Printf("Kind: %s, Type: %s\n", value1.Kind(), value1.Type())
	value2 := reflect.ValueOf(fun2)
	fmt.Printf("Kind: %s, Type: %s\n", value2.Kind(), value2.Type())
	value1.Call(nil)
	value2.Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf("hello")})
}

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello, ", msg)
}
func (p Person) PrintInfo() {
	fmt.Printf("name: %s, age: %d, sex: %s\n",
		p.Name, p.Age, p.Sex)
}
func (p Person) Test(i, j int, s string) {
	fmt.Println(i, j, s)
}
func DoFileAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get type is: ", getType.Name())
	fmt.Println("get kind is: ", getType.Kind())
	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is: ", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("field name: %s, field type: %s, field value: %v \n",
			field.Name, field.Type, value)
	}

	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("method name: %s, method type: %v \n",
			method.Name, method.Type)
	}
}

func fun1() {
	fmt.Println("i am fun1(), no args...")
}
func fun2(i int, s string) {
	fmt.Println("i am fun2(), with args...", i, s)
}
