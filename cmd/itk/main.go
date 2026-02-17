package main

import (
	"fmt"

	"github.com/amaterasutears/itk/config"
)

func main() {
	c, err := config.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
