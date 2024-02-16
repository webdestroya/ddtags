package ddtags

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	tagKey = "ddtag"

	tagFmt = "fmt"
)

// Extracts a list of datadog tags using the "ddtag" struct member tags
func Extract(struc any) []string {

	if struc == nil {
		return []string{}
	}

	vStruc := reflect.ValueOf(struc)

	if vStruc.Kind() != reflect.Interface && vStruc.Kind() != reflect.Pointer {
		return []string{}
	}

	v := vStruc.Elem()

	if v.Kind() != reflect.Struct {
		return []string{}
	}

	sType := v.Type()

	tags := make([]string, 0, 10)

	for i := 0; i < v.NumField(); i++ {
		sTag := sType.Field(i).Tag
		ddTagVal, hasTag := sTag.Lookup(tagKey)

		if !hasTag || ddTagVal == "" {
			continue
		}

		ddTag, tagExtras, _ := strings.Cut(ddTagVal, ",")
		if ddTag == "" || ddTag == "-" {
			continue
		}

		field := v.Field(i)

		// ignore the zero value
		if field.IsZero() {
			continue
		}

		if field.Kind() == reflect.Pointer {
			field = field.Elem()
		}

		var tagValue string

		switch f := field.Interface().(type) {
		case string:
			tagValue = f

		case bool:
			tagValue = strconv.FormatBool(f)

		case int:
			tagValue = fmtInteger(f, tagExtras)
		case int8:
			tagValue = fmtInteger(f, tagExtras)
		case int16:
			tagValue = fmtInteger(f, tagExtras)
		case int32:
			tagValue = fmtInteger(f, tagExtras)
		case int64:
			tagValue = fmtInteger(f, tagExtras)

		case uint:
			tagValue = fmtInteger(f, tagExtras)
		case uint8:
			tagValue = fmtInteger(f, tagExtras)
		case uint16:
			tagValue = fmtInteger(f, tagExtras)
		case uint32:
			tagValue = fmtInteger(f, tagExtras)
		case uint64:
			tagValue = fmtInteger(f, tagExtras)

		case float32:
			tagValue = fmtFloat(f, tagExtras)

		case float64:
			tagValue = fmtFloat(f, tagExtras)

		default:
			continue
		}

		if tagValue != "" {
			tags = append(tags, ddTag+":"+tagValue)
		}
	}

	return tags
}

func fmtFloat[T float32 | float64](value T, tagExtras string) string {
	strFmt := defaultConfig.FloatFormat
	if tagExtras != "" {
		for _, extra := range strings.Split(tagExtras, ",") {
			nameValue := strings.SplitN(extra, "=", 2)
			if len(nameValue) == 2 {
				name, val := nameValue[0], nameValue[1]
				switch name {
				case tagFmt:
					strFmt = val
				}
			}
		}
	}

	return fmt.Sprintf(strFmt, float64(value))
}

func fmtInteger[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](value T, tagExtras string) string {
	strFmt := defaultConfig.IntegerFormat
	if tagExtras != "" {
		for _, extra := range strings.Split(tagExtras, ",") {
			nameValue := strings.SplitN(extra, "=", 2)
			if len(nameValue) == 2 {
				name, val := nameValue[0], nameValue[1]
				switch name {
				case tagFmt:
					strFmt = val
				}
			}
		}
	}

	return fmt.Sprintf(strFmt, value)
}
