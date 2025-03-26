package go_conditionals

import "reflect"

// IsZeroValue checks if the given value is the zero value for its type.
//
// It supports all Go types including booleans, integers, floats, complex numbers,
// strings, pointers, arrays, slices, maps, structs, channels, and functions.
//
// Parameters:
//   - value: interface{} - The value to be checked. Can be of any type.
//
// Returns:
//   - bool: true if value is the zero value for its type, false otherwise.
func IsZeroValue(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return v.IsZero()
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0
	case reflect.String:
		return v.String() == ""
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Chan, reflect.Func:
		return v.IsNil()
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !IsZeroValue(v.Index(i).Interface()) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !IsZeroValue(v.Field(i).Interface()) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// IsEmptySlice checks if the given slice is empty.
//
// This function determines whether the provided interface{} is a slice and if so,
// whether it contains any elements. It returns false for non-slice types.
//
// Parameters:
//   - slice: interface{} - The value to be checked. Should be a slice of any type.
//
// Returns:
//   - bool: true if slice is empty or nil, false otherwise.
func IsEmptySlice(slice interface{}) bool {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return false
	}

	return v.Len() == 0
}

// IsEmptyMap checks if the given map is empty.
//
// This function determines whether the provided interface{} is a map and if so,
// whether it contains any key-value pairs. It returns false for non-map types.
//
// Parameters:
//   - m: interface{} - The value to be checked. Should be a map of any type.
//
// Returns:
//   - bool: true if m is a map and contains no elements, false otherwise.
func IsEmptyMap(m interface{}) bool {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		return false
	}

	return v.Len() == 0
}
