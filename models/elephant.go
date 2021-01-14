package models

import (
	"strings"

	"github.com/applinh/elephant/kvdb"
	"github.com/boltdb/bolt"
)

// Elephant defines an elephant in the db
type Elephant struct {
	Name       string
	Containers []string
}

// NewElephantArmy outputs an array of elephants from a map[ElephantName]="ContainerA,ContainerB"
func NewElephantArmy(elephantsMapList map[string]string) []Elephant {
	var elephants []Elephant
	for k, v := range elephantsMapList {
		containerIDs := strings.Split(v, ",")

		elephant := Elephant{Name: k, Containers: containerIDs}

		elephants = append(elephants, elephant)
	}

	return elephants
}

// RemContainerByID removes a container from the elephant obj
func (elephant *Elephant) RemContainerByID(containerID string) {
	for k, v := range elephant.Containers {
		if v == containerID {
			if k+1 > len(elephant.Containers) {
				if elephant.Containers[:k] == nil {
					elephant.Containers = nil
				} else {
					elephant.Containers = elephant.Containers[:k]
				}

			} else {
				elephant.Containers = append(elephant.Containers[:k], elephant.Containers[k+1:]...)
			}

		}
	}

}

// WriteToDB writes elephant to db
func (elephant *Elephant) WriteToDB(db *bolt.DB) error {
	containersString := strings.Join(elephant.Containers, ",")
	return kvdb.WriteData(db, elephant.Name, containersString)
}

// DelFromDB deletes elephant from db
func (elephant *Elephant) DelFromDB(db *bolt.DB) error {
	return kvdb.DeleteKey(db, elephant.Name)
}
