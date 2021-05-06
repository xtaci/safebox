package extras

import (
	"reflect"
	"strconv"
)

func Number(n interface{}) (num float64) {
	if s, ok := n.(int); ok {
		num = float64(s)
	}
	if s, ok := n.(string); ok {
		num, _ = strconv.ParseFloat(s, 64)
	}
	if s, ok := n.(int64); ok {
		num = float64(s)
	}
	if s, ok := n.(float64); ok {
		num = s
	}
	return
}

// Check if struct is empty
func isEmpty(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool() == false

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return v.Float() == 0

	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0

	case reflect.Ptr, reflect.Interface:
		return isEmpty(v.Elem())

	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isEmpty(v.Index(i)) {
				return false
			}
		}
		return true

	case reflect.Slice, reflect.String, reflect.Map:
		return v.Len() == 0

	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if !isEmpty(v.Field(i)) {
				return false
			}
		}
		return true
		// reflect.Chan, reflect.UnsafePointer, reflect.Func
	default:
		return v.IsNil()
	}
}

// IsEmpty reports whether v is empty struct
// Does not support cycle pointers for performance, so as json
func IsEmpty(v interface{}) bool {
	return isEmpty(reflect.ValueOf(v))
}
