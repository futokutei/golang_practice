package program

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ToYAML(v any) (string, error) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return "", fmt.Errorf("Only struct support, received: %s", val.Kind())
	}

	var parts []string

	for i := 0; i < val.NumField(); i++ {
		fieldType := typ.Field(i)
		fieldValue := val.Field(i)

		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = fieldType.Name
		}

		var jsonValue string
		switch fieldValue.Kind() {
		case reflect.String:
			jsonValue = fmt.Sprintf(`"%s"`, fieldValue.String())

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			jsonValue = strconv.FormatInt(fieldValue.Int(), 10)

		case reflect.Bool:
			jsonValue = strconv.FormatBool(fieldValue.Bool())

		case reflect.Slice:

			var sliceParts []string
			for j := 0; j < fieldValue.Len(); j++ {
				elem := fieldValue.Index(j)
				if elem.Kind() == reflect.String {
					sliceParts = append(sliceParts, fmt.Sprintf(`"%s"`, elem.String()))
				} else {
					sliceParts = append(sliceParts, fmt.Sprintf("%v", elem.Interface()))
				}
			}
			jsonValue = "\n   - " + strings.Join(sliceParts, "\n   - ") + "\n"

		default:
			jsonValue = fmt.Sprintf(`"%v"`, fieldValue.Interface())
		}
		parts = append(parts, fmt.Sprintf(`%s: %s`, jsonTag, jsonValue))
	}

	return strings.Join(parts, "\n"), nil
}
