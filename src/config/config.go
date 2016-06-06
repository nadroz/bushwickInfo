package config

import(
"github.com/BurntSushi/toml"
"io/ioutil"
"fmt"
)

type DbConfig struct {
		BushwickConn	string `toml:"bushwick_con_string"`
		LogConn		string `toml:"log_con_string"`
}

func LoadDbConfig() (*DbConfig, error){
	configPath := "/usr/local/etc/bushwickInfo/config"

	data, loadConfigErr := ioutil.ReadFile(configPath)
	if loadConfigErr != nil{
		fmt.Println("failed to load config file")
		fmt.Println(loadConfigErr)
		return nil, loadConfigErr
	}

	var config DbConfig
	if _, parseConfigErr := toml.Decode(string(data), &config); parseConfigErr != nil {
		fmt.Println("failed to parse config file")
		fmt.Println(loadConfigErr)
		return nil, parseConfigErr
	}

	return &config, nil
}
