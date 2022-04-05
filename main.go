package main

import (
	"fmt"

	"optiz-consulting.com/project/model"
)

func main() {
	c := model.Custom{
		ID: "ASdasdasd",
	}

	fmt.Println(c.ToString())
}
