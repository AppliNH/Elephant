package commands

import (
	. "applinh/elephant/dcfile"
	. "applinh/elephant/dockercontrol"
	"fmt"

	"github.com/boltdb/bolt"
)

func Walk(db *bolt.DB, absPath string) {
	var elephantName string
	fmt.Println("Here's the path of your stack: " + absPath)
	fmt.Println("Give your elephant a name please: ")
	fmt.Scanln(&elephantName)

	t := ReadDCfile(absPath)

	StartStack(db, t, elephantName)
}
