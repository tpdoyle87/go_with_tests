package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	value := getValue(x, fn)

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walk(value.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walk(value.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walk(value.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := value.Recv(); ok {
				walk(v.Interface(), fn)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := value.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
	}

}

func getValue(x interface{}, fn func(input string)) reflect.Value {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}
	return value
}
