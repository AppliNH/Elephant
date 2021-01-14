package dcfile

import (
	"io/ioutil"

	"github.com/applinh/elephant/models"
	"gopkg.in/yaml.v2"
)

// ReadDCfile reads a compose file and return a DockerCompose obj to interact with
func ReadDCfile(path string) (models.DockerCompose, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return models.DockerCompose{}, err
	}
	t := models.DockerCompose{}
	erro := yaml.Unmarshal([]byte(dat), &t)

	if erro != nil {
		return models.DockerCompose{}, err
	}

	return t, nil

}
