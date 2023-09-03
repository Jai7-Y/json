package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"`
	Price    int
	Platform string
	Password string `json:"-"`
	Tags     []string
}

func main() {
	fmt.Println("It is a JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"TeactJs", 299, "LearnCodeonline.in", "abc1233", []string{"web-dev", "js"}},
		{"MERN", 400, "LearnCodeonline.in", "abc233", []string{"web-dev", "Full Stack"}},
		{"SpringBoot", 200, "LearnCodeonline.in", "a33", []string{"web-dev", "java"}},
	}

	// package this

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"name": "TeactJs",
		"Price": 299,
		"Platform": "LearnCodeonline.in",
		"Tags": ["web-dev","js"]
    }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is Valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		// fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Not valid json")
	}

	// some cases key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	// fmt.Println("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v", k, v)
	}
}
