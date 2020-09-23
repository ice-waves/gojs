package conf

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Mysql struct {
	DSN string `yaml:"dsn"`
}

type Database struct {
	Mysql Mysql `yaml:"mysql"`
}

func (Database *Database) GetConf() *Database {
	basePath,err := os.Getwd()
	if err != nil {
		panic("conf path error!")
	}

	fileName := filepath.Join(basePath, "configs/conf", "database.yaml")
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil{
		panic("load conf failed!")
	}

	err = yaml.Unmarshal(yamlFile, Database)

	if err != nil {
		panic("parse conf file failed!")
	}

	return Database
}