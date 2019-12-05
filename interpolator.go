package sqlinterpolator

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Interpolate(query string, args []interface{}) (string, error) {
	if strings.Count(query, "?") != len(args) {
		return "", ErrInvalidArgsCount
	}

	result := query

	for _, arg := range args {
		value, err := driver.DefaultParameterConverter.ConvertValue(arg)
		if err != nil {
			return "", err
		}

		var new string
		switch v := value.(type) {
		case int64:
			new = strconv.FormatInt(v, 10)
		case float64:
			new = strconv.FormatFloat(v, 'g', -1, 64)
		case bool:
			if v {
				new = "1"
			} else {
				new = "0"
			}
		case []byte:
			if v == nil {
				new = "NULL"
			} else {
				return "", ErrInvalidBytes
			}
		case string:
			new = v
		case time.Time:
			if v.IsZero() {
				new = "'0000-00-00'"
			} else {
				new = fmt.Sprintf("'%s'", v.Format("2006-01-02 15:04:05.9999"))
			}
		default:
			return "", errors.New(fmt.Sprintf("Invalid arg type: %s", reflect.TypeOf(v)))
		}
		result = strings.Replace(result, "?", new, 1)
	}

	return result, nil
}
