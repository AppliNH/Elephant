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

	// // Create Box with 30% width of current screen, and height of 20 lines
	// box1 := tm.NewBox(30|tm.PCT, 20, 0)
	// box2 := tm.NewBox(30|tm.PCT, 20, 0)
	// // Add some content to the box
	// // Note that you can add ANY content, even tables

	// fmt.Fprint(box1, "ok !")
	// fmt.Fprint(box2, "ok2 !")
	// // Move Box to approx center of the screen
	// tm.Print(tm.MoveTo(box1.String(), 0|tm.PCT, 40|tm.PCT))
	// tm.Print(tm.MoveTo(box2.String(), 40|tm.PCT, 40|tm.PCT))
	// tm.Flush()
	abs, err := filepath.Abs(a[1])
	if err == nil {
		var elephantName string
		fmt.Println("Here's the path of your stack: " + abs)
		fmt.Println("Give your elephant a name please: ")
		fmt.Scanln(&elephantName)

		t := ReadDCfile(abs)
		StartStack(t, elephantName)

		//CreateNewContainer(abs)
	}
}
