package config

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

const (
	basePath       string = "./cmd/config/resources/"
	configFileProd string = "prod.yml"
	configFileQA   string = "qa.yml"
	configFileDev  string = "dev.yml"
	configFileBase string = "base.yml"
	configFile     string = "./config.yml"
)

const (
	RootPath     string = "./"
	RootPathTest string = "./../"
	Production   string = "prod"
	Testing      string = "qa"
	Develop      string = "dev"
)

func InitializeBasicConfig() {
	generateConfigFile()
	ReadConfigFile(RootPath)
}

func ReadConfigFile(path string) {

	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	// viper.AutomaticEnv() # chanca los env del os

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}

	for _, element := range viper.AllKeys() {
		os.Setenv(element, viper.GetString(element))
	}

}

func generateConfigFile() {
	var err error
	flag.Parse()

	// create target directory if not exist
	err = os.MkdirAll(basePath, 0775)
	if err != nil {
		fmt.Printf("create %s error: %s\n", basePath, err)
		panic(err)
	}

	appEnv := os.Getenv("APP_ENV")
	fmt.Printf("Environment %s \n", appEnv)
	if strings.EqualFold(appEnv, Production) {
		mergeFiles(basePath+configFileProd, basePath+configFileBase, configFile)
	}
	if strings.EqualFold(appEnv, Develop) {
		mergeFiles(basePath+configFileDev, basePath+configFileBase, configFile)
	}
	if strings.EqualFold(appEnv, Testing) {
		mergeFiles(basePath+configFileQA, basePath+configFileBase, configFile)
	}
	if len(appEnv) == 0 {
		mergeFiles(basePath+configFileBase, basePath+configFileBase, configFile)
	}
}

func mergeFiles(masterFile string, baseFile string, output string) {
	fmt.Printf("Master File: %s\n", masterFile)
	fmt.Printf("Base File: %s\n", baseFile)

	var master map[string]interface{}
	file := readFile(baseFile)
	if err := yaml.Unmarshal(file, &master); err != nil {
		panic(err)
	}

	var override map[string]interface{}
	file = readFile(masterFile)
	if err := yaml.Unmarshal(file, &override); err != nil {
		panic(err)
	}

	for k, v := range override {
		master[k] = v
	}

	file, err := yaml.Marshal(master)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(output, file, 0644); err != nil {
		fmt.Printf("Merge error: %s\n", err)
		panic(err)
	}

	fmt.Printf("Override and generate %s ok", output)

}

func validateExist(file string) {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		fmt.Printf("%s not exist \n", file)
		panic(err)
	}
}

func readFile(file string) []byte {
	validateExist(file)
	read, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("%s not read file \n", file)
		panic(err)
	}
	return read
}

func IsProduction() bool {
	appEnv := os.Getenv("APP_ENV")
	if strings.EqualFold(appEnv, Production) {
		return true
	}
	return false
}

func IsTesting() bool {
	appEnv := os.Getenv("APP_ENV")
	if strings.EqualFold(appEnv, Testing) {
		return true
	}
	return false
}

func IsDevelop() bool {
	appEnv := os.Getenv("APP_ENV")
	if strings.EqualFold(appEnv, Develop) {
		return true
	}
	return false
}

func IsLocal() bool {
	appEnv := os.Getenv("APP_ENV")
	if len(appEnv) == 0 {
		return true
	}
	return false
}
