package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/chen-huicheng/GSWGo/stl"
)

func SetValue(rv reflect.Value, value string) error {
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		item, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		rv.SetInt(item)

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		item, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		rv.SetUint(item)
	case reflect.String:
		rv.SetString(value)
	case reflect.Float32, reflect.Float64:
		item, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		rv.SetFloat(item)
	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", rv.Type())
	}
	return nil

}
func InitValue(mod interface{}) error {
	fmt.Println(reflect.TypeOf(mod), reflect.TypeOf(mod).Elem())
	modType := reflect.TypeOf(mod).Elem()
	nf := modType.NumField()
	rv := reflect.ValueOf(mod).Elem()
	for i := 0; i < nf; i++ {
		fmt.Println(modType.Field(i).Tag)
		value := modType.Field(i).Tag.Get("init")
		// modType.Name()
		if value == "" {
			log.Printf("InitMod: struct '%s' field '%s' not 'init' tag ", modType.Name(), modType.Field(i).Name)
			continue
		}
		if value == "magicKey" {
			InitValue(rv.Field(i).Addr().Interface())
			continue
		}
		err := SetValue(rv.Field(i), value)
		if err != nil {
			return err
		}
	}
	return nil
	// rv := reflect.ValueOf(mod).Elem()
}
func InitLog() {
	// 创建、追加、读写，777，所有权限
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.SetFlags(log.Llongfile | log.LstdFlags)
}
func main() {
	InitLog()
	user := stl.User{}
	if err := InitValue(&user); err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	stu := stl.Student{}
	if err := InitValue(&stu); err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu)
}
