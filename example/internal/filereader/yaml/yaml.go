package yaml

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func FromYAMLFile(path string, data interface{}) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(file, data)
	if err != nil {
		return fmt.Errorf("Unmarshal: %v", err)
	}

	return nil
}
