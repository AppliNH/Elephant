package main

import (
	. "applinh/elephant/commands"
	"fmt"
	"os"

	. "applinh/elephant/kvdb"
	"path/filepath"
)

func main() {

	a := os.Args

	fmt.Println("Hi, and welcome to Elephant 🐘")
	// tx,err := InitDB()
	// if err == nil {
	// 	WriteData(tx, "ok", "okiii")
	// }

	// fmt.Println(ReadData(tx, "ok"))
	

	

	switch a[1] {
	case "walk":
		if abs, err := filepath.Abs(a[2]); err == nil {
			Walk(abs)
		}
	default:
		fmt.Println("An error has occured")

	}

}
