package fillSeting

import (
	"errors"
	"reflect"
)

func FillSeting(st interface{}, settings map[string]interface{}) error {
	if reflect.ValueOf(st).Kind() != reflect.Ptr ||
		reflect.ValueOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("第一个参数必须为struc指针")
	}
	if settings == nil {
		return errors.New("setting 不能为空")
	}

	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range settings {
		if field, ok = reflect.ValueOf(st).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st).Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}
