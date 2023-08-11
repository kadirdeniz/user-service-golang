package pkg

import (
	"io/ioutil"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

func ReadConfigs() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "./..")

	err := readYamlFile(filepath.Join(basepath, "./config/application.yaml"))
	if err != nil {
		panic(err)
	}

	err = readYamlFile(filepath.Join(basepath, "./config/database.yaml"))
	if err != nil {
		panic(err)
	}

	err = readYamlFile(filepath.Join(basepath, "./config/redis.yaml"))
	if err != nil {
		panic(err)
	}

	err = readYamlFile(filepath.Join(basepath, "./config/jwt.yaml"))
	if err != nil {
		panic(err)
	}
}

func readYamlFile(path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &Configs)
	if err != nil {
		return err
	}

	return nil
}
