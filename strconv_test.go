package strconv

import (
	"strings"
	"testing"
)

type MyStruct struct {
	A int
	B string
}

func TestString(t *testing.T) {
	interfaces := []interface{}{
		true,                          // bool
		int(1),                        // int
		int8(2),                       // int8
		int16(3),                      // int16
		int32(4),                      // int32
		int64(5),                      // int64
		uint(6),                       // uint
		uint8(7),                      // uint8
		uint16(8),                     // uint16
		uint32(9),                     // uint32
		uint64(10),                    // uint64
		float32(11.1),                 // float32
		float64(12.2),                 // float64
		complex64(13.3 + 14.4i),       // complex64
		complex128(15.5 + 16.6i),      // complex128
		"hello",                       // string
		MyStruct{},                    // struct
		[]interface{}{},               // slice
		[1]interface{}{},              // array
		map[interface{}]interface{}{}, // map
	}
	expected := []string{
		"true",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"11.1",
		"12.2",
		"(13.3+14.4i)",
		"(15.5+16.6i)",
		"\"hello\"",
		"{\n\t\"A\": 0,\n\t\"B\": \"\"\n}",
		"[]",
		"[\n\tnull\n]",
		"{}",
	}
	actual := []string{}

	// Set actual values
	for _, i := range interfaces {
		actual = append(actual, String(i))
	}

	// Test actual values against expected
	for i := 0; i < len(interfaces); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("\nExpected:\t%s\nActual:\t\t%s\n", expected[i], actual[i])
		}
	}
}

func TestMapToString(t *testing.T) {
	m := map[interface{}]interface{}{
		"A": 0,
		"B": "",
		"C": map[interface{}]interface{}{
			0:       "0",
			"":      "",
			true:    "true",
			nil:     "nil",
			"slice": []interface{}{"value1", "value2"},
			"struct": MyStruct{
				A: 425,
				B: "test",
			},
		},
		"D": true,
		"E": nil,
	}
	type result struct {
		expected      string
		matchesActual bool
	}
	results := []result{
		{"{", false},
		{"\t\"A\": 0", false},
		{"\t\"B\": \"\"", false},
		{"\t\"C\": {", false},
		{"\t\t0: \"0\"", false},
		{"\t\t\"\": \"\"", false},
		{"\t\ttrue: \"true\"", false},
		{"\t\t<nil>: \"nil\"", false},
		{"\t\t\"slice\": [", false},
		{"\t\t\t\"value1\",", false},
		{"\t\t\t\"value2\"", false},
		{"\t\t]", false},
		{"\t\t\"struct\": {", false},
		{"\t\t\t\"A\": 425,", false},
		{"\t\t\t\"B\": \"test\"", false},
		{"\t\t}", false},
		{"\t}", false},
		{"\t\"D\": true", false},
		{"\t\"E\": <nil>", false},
		{"}", false},
	}
	actualLines := strings.Split(String(m), "\n")

	for i := 0; i < len(results); i++ {
		for j := 0; j < len(actualLines); j++ {
			if results[i].expected == actualLines[j] {
				results[i].matchesActual = true
				break
			}
		}
	}

	for _, result := range results {
		if !result.matchesActual {
			t.Fatalf("\nExpected:\t%t, the line '%s' is generated\nActual:\t\t%t, the line '%s' is NOT generated\n",
				true,
				result.expected,
				result.matchesActual,
				result.expected,
			)
		}
	}
}
