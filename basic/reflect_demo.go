package basic

import (
	"fmt"
	"reflect"
)

type X struct {
	y byte
	z complex128
}

func (x *X) String() string {
	return fmt.Sprintf("%+v", *x)
}

func reflect_demo() {
	fmt.Println("Hello World!")
	var x X
	fmt.Println(x)

	v := reflect.TypeOf(x)
	fmt.Println("type: ", v)
	fmt.Println("Align: ", v.Align())
	fmt.Println("FieldAlign: ", v.FieldAlign())
	for i := 0; i < v.NumMethod(); i++ {
		fmt.Println("Method", i, v.Method(i).Name)
	}

	fmt.Println("Name: ", v.Name())
	fmt.Println("pkgPath : ", v.PkgPath())
	fmt.Println("kind: ", v.Kind())
	for i := 0; i < v.NumField(); i++ {
		fmt.Println("Field : ", i, v.Field(i).Name)
	}
	fmt.Println("Size : ", v.Size())

	fmt.Println("-------------")
	var x2 float64 = 3.14
	v2 := reflect.ValueOf(x2)
	fmt.Println("type :", v2.Type())
	fmt.Println("kind is float64", v2.Kind() == reflect.Float64)
	fmt.Println("value ", v2.Float())

	fmt.Println("----------")
	var x3 float64 = 3.14
	fmt.Println("x3", x3)
	v3 := reflect.ValueOf(&x3)
	fmt.Println("v.CanSet:", v3.CanSet())
	i := v3.Elem()
	fmt.Println("v.CanSet: ", i.CanSet())
	i.SetFloat(3.1415)
	fmt.Println("i:", i.Float())
	fmt.Println("x: ", x3)

}
