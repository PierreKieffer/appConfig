package appConfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func InitConfig(configFile string, configStruct interface{}) (err error) {
	if len(configFile) == 0 {
		return
	}

	jsonFile, err := os.Open(configFile)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	jsonFileBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(jsonFileBytes), &configStruct)
	if err != nil {
		return
	}
	return
}
