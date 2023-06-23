package main

import (
  "fmt"
  "sort"
)

func main() {
	//begin
	heroes := map[string]bool{
		"peter": true,
		"gwen":  false,
		"bruce": true,
	}
	keys := make([]string, 0, len(heroes))
	for k := range heroes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, heroes[k])
	}
	//end
}
