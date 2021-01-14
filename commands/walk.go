package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/applinh/elephant/dcfile"
	dctrl "github.com/applinh/elephant/dockercontrol"
	dmng "github.com/applinh/elephant/dockermng"
	"github.com/applinh/elephant/models"

	"github.com/boltdb/bolt"
)

// Walk runs a stack from a compose file, save the elephant to db, and connect to logs
func Walk(db *bolt.DB, absPath string, elName string) {
	var elephantName = elName

	fmt.Println("Here's the path of your stack: " + absPath)

	if len(elephantName) == 0 {
		fmt.Println("Give your elephant a name please: ")
		fmt.Scanln(&elephantName)
	}

	t, err := dcfile.ReadDCfile(absPath)
	if err != nil {
		log.Fatal(err)
	}

	containers, cli, err := dctrl.StartStack(db, t, elephantName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All good ! Containers have been packed on your big boi 🐘")

	// Write to DB
	contIDs := []string{}
	for k := range containers {
		contIDs = append(contIDs, k)
	}
	elephant := models.Elephant{Name: elephantName, Containers: contIDs}
	elephant.WriteToDB(db)

	// Logs
	fmt.Println(strings.Repeat("=", 25))
	fmt.Println("Logs")
	fmt.Println(strings.Repeat("=", 25))
	fmt.Println()
	dmng.ReadLogs(cli, containers)

}
