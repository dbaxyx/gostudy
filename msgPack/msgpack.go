package main

import (
	"encoding/json"
	"fmt"
)

type s struct {
	data map[string]interface{}
}

func main() {

}

func jsonDemo() {
	var s1 = s{
		data: make(map[string]interface{}, 8),
	}
	s1.data["count"] = 1
	ret, err := json.Marshal(s1.data)
	if err != nil {
		fmt.Println("marshal failed", err)
	}
}
