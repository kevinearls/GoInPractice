package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	config, _ := yaml.ReadFile(pwd + "/chapter02/technique03/eatme.yaml")
	//if err != nil {
//		fmt.Printf("%v\n", err)
//	}
	p, _ := config.Get("path")
	fmt.Printf(">>>> [%v]\n", p)
}
