package omap

import (
	"errors"
	"reflect"
)

// Facade is 包装真实的值 each的时候 如果iterator返回Facade 则会将Real替换iterator返回值
type Facade struct {
	Real reflect.Value
}

var (
	// ErrorRt is 错误类型
	ErrorRt = reflect.TypeOf(errors.New(""))
	// FacadeRt is 门面类型
	FacadeRt = reflect.TypeOf(Facade{})
	// NullRv is 反射值
	NullRv = reflect.ValueOf(nil)
	// NullRvOfRt is nil反射值类型
	NullRvOfRt = reflect.TypeOf(NullRv)
)

func Map(source, selector interface{}) interface{} {
	var arrRV reflect.Value
	each(source, selector, func(resRV, valueRV, _ reflect.Value) bool {
		if !arrRV.IsValid() {
			arrRT := reflect.SliceOf(resRV.Type())
			arrRV = reflect.MakeSlice(arrRT, 0, 0)
		}

		arrRV = reflect.Append(arrRV, resRV)
		return false
	})
	if arrRV.IsValid() {
		return arrRV.Interface()
	}

	return nil

}

func each(source interface{}, iterator interface{}, predicate func(reflect.Value, reflect.Value, reflect.Value) bool) {
	length, getKeyValue := parseSource(source)
	if length == 0 {
		return
	}

	if predicate == nil {
		predicate = func(resRV, _, _ reflect.Value) bool {
			if resRV.Kind() == reflect.Bool {
				return resRV.Bool()
			}

			return false
		}
	}

	iteratorRV := reflect.ValueOf(iterator)
	for i := 0; i < length; i++ {
		valueRV, keyRV := getKeyValue(i)
		returnRVs := iteratorRV.Call(
			[]reflect.Value{valueRV, keyRV},
		)
		if len(returnRVs) > 0 {
			resRV := returnRVs[0]
			if resRV.Type() == FacadeRt {
				resRV = resRV.Interface().(Facade).Real
			}

			if predicate(resRV, valueRV, keyRV) {
				break
			}
		}
	}
}

func parseSource(source interface{}) (int, func(i int) (reflect.Value, reflect.Value)) {
	if source != nil {
		sourceRV := reflect.ValueOf(source)
		switch sourceRV.Kind() {
		case reflect.Array, reflect.Slice:
			return sourceRV.Len(), func(i int) (reflect.Value, reflect.Value) {
				return sourceRV.Index(i), reflect.ValueOf(i)
			}
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			return len(keyRVs), func(i int) (reflect.Value, reflect.Value) {
				return sourceRV.MapIndex(keyRVs[i]), keyRVs[i]
			}
		}
	}
	return 0, nil
}
