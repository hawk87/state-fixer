package main

import (
	"encoding/json"
	"fmt"
)

//{ "iN": 1 }, 1<=N<=X

func main() {
	fmt.Printf("Enter upper limit: ")
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Printf("error in scan")
		return
	}

	obj := make(map[string]int)

	for i := 1; i <= input; i++ {
		key := fmt.Sprintf("i%d", i)
		value := 1
		obj[key] = value
	}
	jsonObj, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("error in marshal")
		return
	}
	fmt.Println(string(jsonObj))
}
