package master

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	ApiPort int `json:"apiPort"`
	ApiReadTimeout int `json:"apiReadTimeout"`
	ApiWriteTimeout int `json:"apiWriteTimeout"`
	EtcdEndpoints []string `json:"etcdEndpoints"`
	EtcdDialTimeout int `json:"etcdDialTimeout"`
}

var (
	G_config *Config
)

func InitConfig(filename string) (err error) {

	var (
		content []byte
	)

	//read config file
	if content, err = ioutil.ReadFile(filename); err != nil {
		return 
	}

	//deserilize json
	if err = json.Unmarshal(content,&conf); err != nil {
		return
	}

	G_config = &conf

	return
}