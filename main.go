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
	db, err := InitDB()
	if err == nil {
		//WriteData(tx, "ok", "boi")
		//fmt.Println("main is: " + ReadData(db, "ok"))

		switch a[1] {
		case "walk":
			if abs, err := filepath.Abs(a[2]); err == nil {
				Walk(db, abs)
			}

		case "stomp":
			if a[2] != "" {
				Stomp(db, a[2])
			} else {
				fmt.Println("You haven't provided any elephant name")
			}

		default:
			fmt.Println("An error has occured")

		}
	}
}
