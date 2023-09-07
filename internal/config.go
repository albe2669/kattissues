package internal

import (
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	Github struct {
		Token     string `yaml:"token"`
		RepoOwner string `yaml:"repo-owner"`
		RepoName  string `yaml:"repo-name"`
	} `yaml:"github"`
}

var Config *config

func ReadConfig() error {
	config := &config{}
	f, err := os.Open("config.yml")
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		return err
	}

	Config = config

	return nil
}
