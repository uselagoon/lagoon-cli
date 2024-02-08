package app

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/integralist/go-findroot/find"
	"gopkg.in/yaml.v3"
)

// LagoonProject represets the information of a Lagoon project.
type LagoonProject struct {
	Dir               string
	Name              string
	Environment       string
	DockerComposeYaml string `yaml:"docker-compose-yaml"`
}

// LagoonDockerCompose represents the docker-compose.yml file contents for a Lagoon project.
type LagoonDockerCompose struct {
	LagoonProject string `yaml:"x-lagoon-project"`
}

// ReadConfig reads the configuration files of a Lagoon project.
func (project *LagoonProject) ReadConfig() error {
	var err error

	source, err := os.ReadFile(filepath.Join(project.Dir, ".lagoon.yml"))
	if err != nil {
		return fmt.Errorf("unable to load config file %s/: %v", filepath.Join(project.Dir, ".lagoon.yml"), err)
	}
	err = yaml.Unmarshal(source, project)
	if err != nil {
		return fmt.Errorf("unable to load config file %s/: %v", filepath.Join(project.Dir, ".lagoon.yml"), err)
	}

	dockerComposeFilepath := filepath.Join(project.Dir, project.DockerComposeYaml)
	if !fileExists(dockerComposeFilepath) {
		return fmt.Errorf("Could not load docker-compose.yml at %s", dockerComposeFilepath)
	}
	sourceCompose, _ := os.ReadFile(dockerComposeFilepath)
	var dockerCompose LagoonDockerCompose
	yaml.Unmarshal(sourceCompose, &dockerCompose)
	// Reset the name based on the docker-compose.yml file.
	project.Name = dockerCompose.LagoonProject

	gitCmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	gitBranch, err := gitCmd.Output()
	if err == nil {
		project.Environment = strings.TrimSpace(string(gitBranch))
	}

	return nil
}

// GetLocalProject returns the current Lagoon app detected.
func GetLocalProject() (LagoonProject, error) {
	_, err := os.Getwd()
	if err != nil {
		return LagoonProject{}, fmt.Errorf("error determining the current directory: %s", err)
	}

	// the .lagoon.yml file will likely be in a git repo, we should check if the directory we are in is a repo :)
	root, err := find.Repo()
	if err != nil {
		return LagoonProject{}, err
	}
	appDir := fmt.Sprintf("%+v", root)
	return getProjectFromPath(appDir)
}

func getProjectFromPath(path string) (LagoonProject, error) {
	app := LagoonProject{}
	var err error

	appDir := path
	appDir, err = findLocalProjectRoot(appDir)
	if err != nil {
		return app, err
	}
	app.Name = filepath.Base(appDir)
	app.Dir = appDir
	app.ReadConfig()
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
