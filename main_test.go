package main_test

import (
	"fmt"
	"testing"
)

func Test_name(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "map"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			json := map[string] int {
				"age": 8,
			}
			json["weight"] = 9
			fmt.Println("before ", json)
			updateMap(json)
			fmt.Println("after ", json)
		})
	}
}

func updateMap(json map[string] int) {
	json["age"] = 9
}
