package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	valueTest()
	typeTest()
}
func valueTest() {
	m := map[string]string{"name": "张三", "age": "20"}
	p := &Person{}
	pt := reflect.ValueOf(p)
	if pt.Kind() != reflect.Ptr {
		return
	}
	pt = pt.Elem()
	if pt.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < pt.NumField(); i++ {
		v, ok := m[reflect.TypeOf(p).Elem().Field(i).Tag.Get("json")]
		if !ok {
			continue
		}
		ptf := reflect.Indirect(pt).Field(i)
		switch ptf.Kind() {
		case reflect.String:
			ptf.SetString(v)
		case reflect.Int32:
			_v, _ := strconv.ParseInt(v, 10, 32)
			ptf.Set(reflect.ValueOf(_v))
		case reflect.Int:
			_v, _ := strconv.Atoi(v)
			ptf.Set(reflect.ValueOf(_v))
		case reflect.Int64:
			_v, _ := strconv.ParseInt(v, 10, 64)
			ptf.Set(reflect.ValueOf(_v))
		case reflect.Bool:
			_v, _ := strconv.ParseBool(v)
			ptf.SetBool(_v)
		case reflect.Ptr:
			switch reflect.TypeOf(p).Elem().Field(i).Type.Elem().Kind() {
			case reflect.String:
				_v := v
				ptf.Set(reflect.ValueOf(&_v))
			case reflect.Int32:
				_v, _ := strconv.ParseInt(v, 10, 32)
				_V := int32(_v)
				ptf.Set(reflect.ValueOf((&(_V))))
			case reflect.Int:
				_v, _ := strconv.Atoi(v)
				ptf.Set(reflect.ValueOf((&(_v))))
			case reflect.Int64:
				_v, _ := strconv.ParseInt(v, 10, 64)
				ptf.Set(reflect.ValueOf((&(_v))))
			case reflect.Bool:
				_v, _ := strconv.ParseBool(v)
				ptf.Set(reflect.ValueOf((&_v)))
			}
		}
	}
	fmt.Println(p)
}

func typeTest() {
	m := map[string]string{"name": "张三", "age": "20"}
	p := Person{}

	pt := reflect.TypeOf(p)
	if pt.Kind() == reflect.Ptr {
		pt = pt.Elem()
	}
	if pt.Kind() != reflect.Struct {
		return
	}
	fmt.Println(pt.Kind().String())
	fmt.Println(reflect.TypeOf(m).Kind().String())
	for i := 0; i < pt.NumField(); i++ {
		ptf := pt.Field(i)
		ptft := ptf.Type
		if ptft.Kind() == reflect.Ptr {
			ptft = ptft.Elem()
		}
		fmt.Println(ptf.Name, ptft.Kind().String(), ptf.Tag.Get("json"))
	}
}
