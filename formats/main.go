package formats

import (
	"fmt"
	"reflect"
	"strings"
)

func Sprintf(args ...any) (string, error) {
	result, err := Formats(args...)
	if err != nil {
		return "", err
	}
	return strings.Join(result, " "), nil
}

func Formats(args ...any) (result []string, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%s", rec)
		}
	}()
	list := make([]string, len(args))[0:0]
	for i := 0; i < len(args); i++ {
		target := args[i]
		str, ok := target.(string)
		if !ok {
			return nil, fmt.Errorf("expect \"string\", got \"%s\"", reflect.TypeOf(target))
		}
		count := CountFormatArgs(str)
		list = append(list, fmt.Sprintf(str, args[i+1:i+count+1]...))
		i = i + count
	}
	result = make([]string, len(list))
	copy(result, list)
	return result, nil
}

func CountFormatArgs(str string) int {
	before := -2
	count := 0
	for idx, char := range str {
		if char == '%' && before != (idx-1) {
			count++
			before = idx
		}
	}
	return count
}
