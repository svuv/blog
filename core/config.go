package core

import (
	"blog/config"
	"blog/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

/**
*@Author: whh
*@CreateTime: 2023-03-04  19:06
 */

func InitCore() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error :%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init unmarshal: %s", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}
