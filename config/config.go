package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type conf struct {
	Server struct {
		Address string
	}
	Context struct {
		Timeout int8
	}
	Database struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
}

func GetConf(filename string) (*conf, error) {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	c := &conf{}

	err = yaml.Unmarshal(buf, c)

	if err != nil {
		return nil, fmt.Errorf("error on file %q: %v", filename, err)
	}

	return c, nil
}
