package devslog

import (
	"reflect"
	"strings"
)

func (h *developHandler) getFileLineFromQuickjs(err error) (fileLines []string) {
	v := reflect.ValueOf(err)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	// stack is string
	v = v.FieldByName("Stack")
	if v.Kind() != reflect.String {
		return nil
	}

	for _, line := range strings.Split(v.String(), "\n") {
		if strings.Contains(line, "at ") {
			fileLines = append(fileLines, strings.TrimPrefix(strings.Trim(line, " "), "at "))
		}
	}

	return fileLines
}
