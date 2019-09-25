package instruction

import (
	"fmt"
	"reflect"
	"strconv"
)

// InspectStruct accepts a struct and prints info about that struct
func InspectStruct(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), "")
}

func display(path string, v reflect.Value, structTag string) {
	template := "%-70s = %-10s - %-40s\n"
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), structTag)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			currentStructTag := ""
			if yamlTag := v.Type().Field(i).Tag.Get("yaml"); yamlTag != "" {
				currentStructTag = currentStructTag + "yaml:" + yamlTag
			}
			if jsonTag := v.Type().Field(i).Tag.Get("json"); jsonTag != "" {
				currentStructTag = currentStructTag + "json:" + jsonTag
			}
			display(fieldPath, v.Field(i), currentStructTag)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key), structTag)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf(template, path, "nil", structTag)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), structTag)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf(template, path, "nil", structTag)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), structTag)
		}
	default: // basic types, channels, funcs
		fmt.Printf(template, path, formatAtom(v), structTag)
	}
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
