package clients

import "reflect"

func MapToStruct(m map[string]interface{}, result interface{}) error {
	v := reflect.ValueOf(result).Elem()
	for key, value := range m {
		field := v.FieldByName(key)
		if !field.IsValid() {
			continue
		}
		fieldType := field.Type()
		val := reflect.ValueOf(value)

		if val.Type().ConvertibleTo(fieldType) {
			field.Set(val.Convert(fieldType))
		}
	}
	return nil
}
