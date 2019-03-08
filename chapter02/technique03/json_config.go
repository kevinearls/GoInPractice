package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled	bool
	Path string
}

func main() {
	file, error := os.Open("/Users/kearls/go/src/github.com/kevinearls/GoInPractice/chapter02/technique03/conf.json")
	if error != nil {
		fmt.Printf("%v\n", error)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(conf.Path)


}
