package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

var sum = 0.0

func main() {
	data, _ := ioutil.ReadFile("input")
	var u interface{}

	err := json.Unmarshal(data, &u)
	if err != nil {
		log.Fatalln(err)
	}

	parse(u)

	fmt.Println("Sum: ", sum)
}

func parse(v interface{}) {
	switch v.(type) {
	case string:
		// ignore
	case float64:
		sum += v.(float64)
	case []interface{}:
		for _, v2 := range v.([]interface{}) {
			parse(v2)
		}
	case map[string]interface{}:
		for _, v2 := range v.(map[string]interface{}) {
			parse(v2)
		}
	default:
		x := reflect.ValueOf(v)
		fmt.Println("unknown ", v, x.Kind(), x.Type())
	}
}
