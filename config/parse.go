package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type repoConfig struct {
	RepoURL string `yaml:"repo"`
	Path    string `yaml:"path"`
	Command string `yaml:"cmd"`
}

func ReadConfig(configFile string) []repoConfig {
	c := &[]repoConfig{}
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return *c
}
