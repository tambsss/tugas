package main

import (
	"encoding/json"
	"fmt"
) // import json

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var object = []User{{"Hauzan", 20}, {"Shacent", 20}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
