package myutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configf string = "./config.json"

// Environment describe environment of the app
type Environment struct {
	ProjectID string
	Logging   struct {
		Type string
	}
}

// GetEnvNameFromArgs will return the environment name. Default is "dev"
func GetEnvNameFromArgs() string {
	args := os.Args[1:]

	if len(args) <= 0 {
		return "dev"
	}

	envArg := args[0]

	return envArg
}

// GetEnvFromArgs will return the Environment.
// Because it is a main process. It would panic if anything fail during the process.
func GetEnvFromArgs() (*Environment, string) {
	envName := GetEnvNameFromArgs()

	envs := GetConfigMap(configf)

	env, envExist := (*envs)[envName]
	if !envExist {
		panic("Can't find envName:" + envName + " among all environments in " + configf)
	}
	return &env, envName
}

// GetConfigMap will return a map of Environment
func GetConfigMap(path string) *map[string]Environment {
	rawConfig, errRead := ioutil.ReadFile(path)
	if errRead != nil {
		fmt.Println(errRead)
		panic("Can't find " + configf + ". Quiting.")
	}

	envs := map[string]Environment{}
	if err := json.Unmarshal(rawConfig, &envs); err != nil {
		fmt.Println(err)
		panic("Can't parse content from " + configf + " to Environment. Quiting.")
	}
	return &envs
}
