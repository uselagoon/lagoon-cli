package app

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type LagoonProject struct {
	Dir               string
	Name              string
	DockerComposeYaml string `yaml:"docker-compose-yaml"`
}

func (project *LagoonProject) ReadConfig() error {
	var err error

	source, err := ioutil.ReadFile(filepath.Join(project.Dir, ".lagoon.yml"))
	err = yaml.Unmarshal(source, project)
	if err != nil {
		return fmt.Errorf("unable to load config file %s/: %v", filepath.Join(project.Dir, ".lagoon.yml"), err)
	}

	return nil
}

// GetLocalProject returns the current Lagoon app detected.
func GetLocalProject() (*LagoonProject, error) {
	app := &LagoonProject{}
	var err error

	appDir, err := os.Getwd()
	if err != nil {
		return app, fmt.Errorf("error determining the current directory: %s", err)
	}
	appDir, err = findLocalProjectRoot(appDir)
	if err != nil {
		return app, err
	}
	app.Name = filepath.Base(appDir)
	app.Dir = appDir

	return app, nil
}

func findLocalProjectRoot(path string) (string, error) {
	var check = filepath.Join(path, ".lagoon.yml")
	if fileExists(check) {
		return path, nil
	}
	for filepath.Dir(path) != path {
		path = filepath.Dir(path)
		if fileExists(filepath.Join(path, ".lagoon.yml")) {
			return path, nil
		}
	}
	return "", fmt.Errorf("no %s file was found in this directory or any parent", filepath.Join(".lagoon.yml"))
}

// FileExists checks a file's existence
func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
