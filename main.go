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
		fmt.Println(abs)
		t := ReadDCfile(abs)
		StartStack(t)

		//CreateNewContainer(abs)
	}
}
