package strconv

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func String(i interface{}) string {
	return stringRecursive(i, "", 0)
}

func stringRecursive(i interface{}, out string, depth int) string {
	if isMap(i) {
		return mapToString(i.(map[interface{}]interface{}), out, depth)
	} else if isJsonIncompatible(i) {
		return fmt.Sprintf("%v", i)
	} else {
		return toJson(i, depth)
	}
}

func isJsonIncompatible(v interface{}) bool {
	jsonIncompatibleTypes := []reflect.Type{
		reflect.TypeOf(complex64(0)),
		reflect.TypeOf(complex128(0)),
	}

	for _, t := range jsonIncompatibleTypes {
		if reflect.TypeOf(v).AssignableTo(t) {
			return true
		}
	}
	return false
}

func isMap(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Map
}

func mapToString(m map[interface{}]interface{}, out string, depth int) string {
	if len(m) == 0 {
		return "{}"
	}

	out += "{\n"

	for k, v := range m {
		for i := 0; i < depth+1; i++ {
			out += "\t"
		}

		if k == nil {
			out += "<nil>"
		} else {
			out += stringRecursive(k, "", depth+1)
		}

		out += ": "

		if v == nil {
			out += "<nil>"
		} else {
			out += stringRecursive(v, "", depth+1)
		}

		out += "\n"
	}

	for i := 0; i < depth; i++ {
		out += "\t"
	}

	out += "}"

	return out
}

func toJson(i interface{}, depth int) string {
	jsonStr, err := json.MarshalIndent(i, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	return indentToDepth(string(jsonStr), depth)
}

func indentToDepth(s string, depth int) string {
	lines := strings.Split(s, "\n")

	if len(lines) > 1 {
		for i := 1; i < len(lines); i++ {
			indentedLine := ""

			for i := 0; i < depth; i++ {
				indentedLine += "\t"
			}

			indentedLine += lines[i]
			lines[i] = indentedLine
		}
	}

	return strings.Join(lines, "\n")
}