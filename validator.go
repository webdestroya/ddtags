package ddtags

import (
	"fmt"
	"reflect"
	"strings"
)

// Returns an error if any fields in the struct (tagged with ddtag) are invalid and will be ignored
// You can use this inside your test suite
func Validate(struc any) error {
	v := reflect.ValueOf(struc).Elem()
	sType := v.Type()
	for i := 0; i < v.NumField(); i++ {
		ddTagVal, hasTag := sType.Field(i).Tag.Lookup(tagKey)

		if !hasTag || ddTagVal == "" {
			continue
		}

		ddTag, _, _ := strings.Cut(ddTagVal, ",")
		if ddTag == "" || ddTag == "-" {
			continue
		}

		field := v.Field(i)

		switch f := field.Interface().(type) {
		case string:
		case bool:
		case int, int8, int16, int32, int64:
		case uint, uint8, uint16, uint32, uint64:
		case float32:
		case float64:

		// case *string:
		// case *bool:
		// case *float32:
		// case *float64:
		// case *int, *int8, *int16, *int32, *int64:
		// case *uint, *uint8, *uint16, *uint32, *uint64:
		default:
			return fmt.Errorf("invalid field type: %s (%T)", sType.Field(i).Name, f)
		}
	}

	return nil
}
