package main

import (
	. "applinh/elephant/commands"
	"fmt"
	"os"
	"strings"

	. "applinh/elephant/kvdb"
	"path/filepath"
)

func main() {

	a := os.Args

	fmt.Println("Hi, and welcome to Elephant 🐘")
	fmt.Println(strings.Repeat("_", 25))
	fmt.Println()
	db, err := InitDB()
	if err == nil {
		//WriteData(tx, "ok", "boi")
		//fmt.Println("main is: " + ReadData(db, "ok"))

		switch a[1] {
		case "walk":
			if len(a) > 3 {
				if abs, err := filepath.Abs(a[2]); err == nil {
					Walk(db, abs)
				}
			} else {
				fmt.Println("You must provide the path to your docker-compose file like this : ~/garbageDir/poop.yml or ./garbageDir/poop.yml")
			}
		case "ls":
			List(db)
		case "stomp":
			if len(a) > 3 {
				if a[2] != "" {
					Stomp(db, a[2])
				}
			} else {
				fmt.Println("You haven't provided any elephant name")
			}
		case "help":
			Help()
		default:
			fmt.Println("An error has occured")

		}
	}
}
