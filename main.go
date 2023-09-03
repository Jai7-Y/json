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
	EncodeJson()
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
