package omap

import (
	"testing"
)

type dataStruct struct {
	key   int
	value string
}

func TestMap(t *testing.T) {
	data := []dataStruct{
		{key: 1, value: "123"},
		{key: 4, value: "456"},
		{key: 7, value: "789"},
	}
	result := Map(data, func(d dataStruct, _ int) map[string]interface{} {
		return map[string]interface{}{
			"title":   d.key,
			"content": d.value,
		}
	}).([]map[string]interface{})
	t.Logf("%+v", result)
}
