package main

import (
	"fmt"
	"log"
	"github.com/supercobra/goreddit"
)

func main() {
	items, err := goreddit.Get("golang")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item)
	}

}
