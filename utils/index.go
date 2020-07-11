package utils

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GenerateUuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Println(uuid)
	return uuid
}

func WriteDockerComposeFile(uuid string, data string) (string, error) {

	_, erro := os.Stat("./composefiles/" + uuid)

	if os.IsNotExist(erro) {
		errDir := os.MkdirAll("./composefiles/"+uuid, 0755)
		if errDir != nil {
			log.Fatal(erro)
		}
	}
	if err := ioutil.WriteFile("./composefiles/"+uuid+"/docker-compose.yml", []byte(data), 0644); err == nil {
		fmt.Println("Wrote in docker-compose.yml")
		return uuid, nil
	} else {
		fmt.Println(err)
		return "", err
	}

}
