package dcfile

import (
	. "applinh/elephant/models"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ReadDCfile(path string) DockerCompose {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	t := DockerCompose{}
	erro := yaml.Unmarshal([]byte(dat), &t)
	if erro != nil {
		fmt.Println(err)
	}
	fmt.Print(t)
	return t
	//return t, err
}
