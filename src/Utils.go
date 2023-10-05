package src

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type GhostTask struct {
    Name string
    State string
}

func (g *GhostTask) Progress() {
    fmt.Printf("%s\t%s\r", g.State, g.Name);
}

func (g *GhostTask) Complete() {
    fmt.Printf("\033[0;32m%s\t%s: %s\033[0m\n", "👻", g.Name, "Complete");
    
}

func (g *GhostTask) Fail(err error) {
    fmt.Printf("\033[0;31m%s\t%s: %s\033[0m\n", "🛑", g.Name, err.Error());
}

type GhostConfig struct {
    Name string `yaml:"name"`
    Version string `yaml:"version"`
    Description string `yaml:"description"`
    Port int `yaml:"port"`
    SurrealDB struct {
        URL string `yaml:"surrealdb-url"`
        Username string `yaml:"surrealdb-username"`
        Password string `yaml:"surrealdb-password"`
        Database string `yaml:"surrealdb-database"`
        Collection string `yaml:"surrealdb-collection"`
    } `yaml:"surrealdb"`
    TailwindCSS struct {
        Input string `yaml:"input"`
        Output string `yaml:"output"`
    } `yaml:"tailwindcss"`
}

// LoadGhostConfig loads a ghost config from the current working directory 
// and returns a GhostConfig struct. Default path is ./ghost.yaml
func LoadGhostConfig() (GhostConfig, error) {
    // load the ghost.yaml file 
    ghostConfig := GhostConfig{};
    ghostConfigFile, err := ioutil.ReadFile("./ghost.yaml");
    if err != nil {
        err = fmt.Errorf("ghost encountered an error loading the ghost.yaml file: %s", err.Error());
        return ghostConfig, err;
    }
    err = yaml.Unmarshal(ghostConfigFile, &ghostConfig);
    if err != nil {
        err = fmt.Errorf("ghost encountered an error loading the ghost.yaml file: %s", err.Error());
        return ghostConfig, err;
    }
    return ghostConfig, nil;
}

// LoadGhostConfigFromPath loads a ghost config from a path 
// and returns a GhostConfig struct. must be a valid path 
// to a valid ghost.yaml file
func LoadGhostConfigFromPath(path string) (GhostConfig, error) {
    ghostConfig := GhostConfig{};
    ghostConfigFile, err := ioutil.ReadFile(path);
    if err != nil { 
        err = fmt.Errorf("ghost encountered an error loading the ghost config: %s", err.Error());
        return ghostConfig, err;
    }

    err = yaml.Unmarshal(ghostConfigFile, &ghostConfig); 
    if err != nil { 
        err = fmt.Errorf("ghost encountered an error loading the ghost config: %s", err.Error());
        return ghostConfig, err;
    }
    return ghostConfig, nil;
}

func Ghost(config string) (GhostConfig, error) {
    if config == "" {
        return LoadGhostConfig();
    }
    return LoadGhostConfigFromPath(config);
}
