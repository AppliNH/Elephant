package commands

import (
	. "applinh/elephant/dockercontrol"
	. "applinh/elephant/kvdb"
	"fmt"
	"regexp"
	"strings"

	"github.com/boltdb/bolt"
)

func List(db *bolt.DB) {

	list := make(map[string]string)
	list, _ = ReadAll(db)

	for k, v := range list {
		var containNumber = regexp.MustCompile(`\d`)
		if containNumber.MatchString(v) {
			fmt.Println("🐘 " + k + ":")
			containerIDs := strings.Split(v, ",")
			for _, c := range containerIDs {
				cont, err := InspectContainer(c)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(cont.Name)
				}

			}
		} else {
			fmt.Println("No big boi found🙅")
		}
	}

}
