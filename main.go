package main

import (
	"flag"
	"fmt"
)

func main() {
	description := flag.String("description", "", "description of the task")

	flag.Parse()

	fmt.Println("Description:", *description)
}