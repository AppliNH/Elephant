package commands

import (
	. "applinh/elephant/dcfile"
	. "applinh/elephant/dockercontrol"
	"fmt"
)

func Walk(absPath string) {
	var elephantName string
		fmt.Println("Here's the path of your stack: " + absPath)
		fmt.Println("Give your elephant a name please: ")
		fmt.Scanln(&elephantName)

		t := ReadDCfile(absPath)
		
		StartStack(t, elephantName)
}