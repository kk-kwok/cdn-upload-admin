package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/toolkits/file"
)

// HttpConfig : define http config
type HttpConfig struct {
	Listen   string `json:"listen"`
	RootPath string `json:"root_path"`
}

// DBConfig : define db config
type DBConfig struct {
	Addr string `json:"addr"`
	Idle int    `json:"idle"`
	Max  int    `json:"max"`
}

// GlobalConfig : define global config
type GlobalConfig struct {
	Debug bool        `json:"debug"`
	HTTP  *HttpConfig `json:"http"`
	DB    *DBConfig   `json:"db"`
}

var (
	// ConfigFile : define config file
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

// Config : get config
func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

// Parse : parse config
func Parse(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("configuration file %s is nonexistent", cfg)
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		return fmt.Errorf("read configuration file %s fail %s", cfg, err.Error())
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		return fmt.Errorf("parse configuration file %s fail %s", cfg, err.Error())
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = &c

	log.Println("load configuration file", cfg, "successfully")
	return nil
}
