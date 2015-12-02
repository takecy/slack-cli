package slacker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config is config for slack
type Config struct {
	// URL is incoming url by slack
	URL string `json:"url"`
}

var (
	h            = os.Getenv("HOME")
	confDirName  = ".slack_cli"
	confFileName = "conf.json"
	confDir      = fmt.Sprintf("%s/%s", h, confDirName)
	confPath     = fmt.Sprintf("%s/%s", confDir, confFileName)
)

// ExistConfig is check conf finle exists
func ExistConfig() bool {
	_, err := os.Stat(confPath)
	return err == nil
}

// CreateConfig is initialize conf file
func CreateConfig() (err error) {
	err = os.MkdirAll(confDir, 0755)
	if err != nil {
		return
	}

	_, err = os.Create(confPath)
	return
}

// ReadConfig is read conf file
func ReadConfig() (c *Config, err error) {
	b, err := ioutil.ReadFile(confPath)
	if err != nil {
		return
	}

	c = new(Config)
	err = json.Unmarshal(b, c)
	return
}

// WriteConfig is write to conf file
func WriteConfig(c *Config) (err error) {
	b, err := json.Marshal(c)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(confPath, b, 0755)
	return
}

// RemoveConfig is remove conf file
func RemoveConfig() (err error) {
	err = os.RemoveAll(confPath)
	return
}
