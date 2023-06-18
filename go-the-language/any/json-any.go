package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//begin
	data := []byte(`{
    "name":"John Wick",
    "born": 1964,
    "hobbies":[
        "martial arts",
        "piano"
    ]
    }`)

	p := map[string]any{}
	err := json.Unmarshal(data, &p)
	fmt.Printf("Johns name: %v", p["name"])
	//end
	if err != nil {
		fmt.Println("Error unmarshal json,", err)
	}

}
