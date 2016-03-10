package config

// test the yaml config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil" // 读取指定文件的数据出来，返回一个字节数组
)

type Config struct {
	User  string      `yaml:"user"`
	Sex   string      `yaml:"sex"`
	Age   string      `yaml:"age"`
	Mysql MySQLConfig `yaml:"mysql"`
	Mongo MySQLConfig `yaml:"mongo"`
}

type MySQLConfig struct {
	URL   string `yaml:"url"`
	SQL   string `yaml:"sql"`
	DB    string `yaml:"dbname"`
	Table string `yaml:"tablename"`
}

type MongoConfig struct {
	URL        string `yaml:"url"`
	DB         string `yaml:"dbname"`
	Collection string `yaml:"collection"`
}

func ParseConfigfile(filename string) (MongoConfig, err) {
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		return err
	}
	var cnf MongoConfig
	if err := yaml.Unmarshal(data, &cnf); err != nil {
		return err
	}
	return cnf, nil
}
