package main

import (
	. "applinh/elephant/dcfile"
	. "applinh/elephant/dockercontrol"
	"fmt"
	"os"

	"path/filepath"
)

func main() {

	a := os.Args

	fmt.Println("Hi, and welcome to Elephant 🐘")
	abs, err := filepath.Abs(a[1])
	if err == nil {
		var elephantName string
		fmt.Println("Here's the path of your stack: " + abs)
		fmt.Println("Give your elephant a name please: ")
		fmt.Scanln(&elephantName)

		t := ReadDCfile(abs)
		fmt.Println(t)
		StartStack(t, elephantName)

		//CreateNewContainer(abs)
	}
}
