package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`{
    "name":"John Wick",
    "born": 1964,
    "hobbies":[
        "martial arts",
        "breakfast foods",
        "piano"
    ]
    }`)

	p := map[string]any{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		fmt.Println("Error unmarshal json,", err)
	}

	for k, v := range p {
		//begin
		switch c := v.(type) {
		case string:
			fmt.Printf("Item %v is a string, containing %q\n", k, c)
		case float64:
			fmt.Printf("Item %v is a number, specifically %f\n", k, c)
		default:
			fmt.Printf("Type of item %v might be %T\n", k, c)
		}
		//end
	}
}
