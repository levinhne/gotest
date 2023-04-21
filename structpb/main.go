package main

import (
	"log"

	"github.com/mitchellh/mapstructure"
)

type Model struct {
	ID   string
	Name string
}

type Event struct {
	ID string
}

var m = map[string]map[string]interface{}{
	"a": {
		"filter": make(map[string]interface{}),
		"domain": &Model{},
		"event":  &Event{},
	},
}

func main() {
	var A = map[string]interface{}{
		"filter": map[string]interface{}{
			"_id": 1,
		},
		"data": map[string]interface{}{
			"id":   "string-id",
			"name": "string-name",
		},
	}
	var d = m["a"]["domain"]
	mapstructure.Decode(A["data"], d)
	log.Println(d)
}
