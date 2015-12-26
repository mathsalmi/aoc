package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

func main() {
	data, _ := ioutil.ReadFile("input")
	var u interface{}

	err := json.Unmarshal(data, &u)
	if err != nil {
		log.Fatalln(err)
	}

	n, o := parse(u)

	fmt.Println("Sum: ", n, o)
}

func parse(v interface{}) (n float64, hasStr bool) {
	switch v.(type) {
	case string:
		// ignore
		return 0, v.(string) == "red"
	case float64:
		return v.(float64), false
	case []interface{}:
		lsum := 0.0
		for _, v2 := range v.([]interface{}) {
			n, _ := parse(v2)
			lsum += n
		}

		return lsum, false
	case map[string]interface{}:
		lsum := 0.0
		for _, v2 := range v.(map[string]interface{}) {
			n, hasStr := parse(v2)
			if hasStr {
				return 0, false
			}

			lsum += n
		}

		return lsum, false
	default:
		x := reflect.ValueOf(v)
		fmt.Println("unknown ", v, x.Kind(), x.Type())
	}

	return 0, true
}
