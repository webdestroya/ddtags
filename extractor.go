package ddtags

import (
	"reflect"
	"strconv"
	"strings"
)

const (
	tagKey = "ddtag"

	tagPrecision = "precision"
	tagBitSize   = "bitsize"
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
			if field.IsNil() {
				continue
			}
			field = field.Elem()
		}

		var tagValue string

		switch f := field.Interface().(type) {
		case string:
			tagValue = f

		case bool:
			tagValue = strconv.FormatBool(f)

		case int:
			tagValue = strconv.FormatInt(int64(f), 10)
		case int8:
			tagValue = strconv.FormatInt(int64(f), 10)
		case int16:
			tagValue = strconv.FormatInt(int64(f), 10)
		case int32:
			tagValue = strconv.FormatInt(int64(f), 10)
		case int64:
			tagValue = strconv.FormatInt(f, 10)

		case uint:
			tagValue = strconv.FormatUint(uint64(f), 10)
		case uint8:
			tagValue = strconv.FormatUint(uint64(f), 10)
		case uint16:
			tagValue = strconv.FormatUint(uint64(f), 10)
		case uint32:
			tagValue = strconv.FormatUint(uint64(f), 10)
		case uint64:
			tagValue = strconv.FormatUint(f, 10)

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

	precision := defaultConfig.FloatPrecision
	bitSize := defaultConfig.FloatBitSize

	if tagExtras != "" {
		for _, extra := range strings.Split(tagExtras, ",") {
			nameValue := strings.SplitN(extra, "=", 2)
			if len(nameValue) == 2 {
				name, val := nameValue[0], nameValue[1]
				switch name {
				case tagPrecision:
					if nv, err := strconv.ParseInt(val, 10, 64); err == nil {
						precision = int(nv)
					}
				case tagBitSize:
					if nv, err := strconv.ParseInt(val, 10, 64); err == nil {
						bitSize = int(nv)
					}
				}
			}
		}
	}

	// if v, ok := sTag.Lookup(tagPrecision); ok {
	// 	if nv, err := strconv.ParseInt(v, 10, 64); err == nil {
	// 		precision = int(nv)
	// 	}
	// }

	// if v, ok := sTag.Lookup(tagBitSize); ok {
	// 	if nv, err := strconv.ParseInt(v, 10, 64); err == nil {
	// 		bitSize = int(nv)
	// 	}
	// }

	return strconv.FormatFloat(float64(value), 'f', precision, bitSize)
}
