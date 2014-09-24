package conf

import (
	"encoding/json"
	"github.com/macococo/go-gamereviews/utils"
	"io/ioutil"
)

type AppConfig struct {
	Port            int
	Runmode         string
	Datasource      string
	MaxIdleConns    int
	MaxOpenConns    int
	NewrelicLicense string
	MemcachedServer string
}

type RunmodeEnums struct {
	DEV  *utils.StringEnum
	PROD *utils.StringEnum
}

var (
	Config  *AppConfig
	Runmode *RunmodeEnums
)

func init() {
	Config = &AppConfig{}

	jsonString, err := ioutil.ReadFile("./conf/conf.json")
	utils.HandleError(err)

	err = json.Unmarshal(jsonString, &Config)
	utils.HandleError(err)

	Runmode = &RunmodeEnums{DEV: &utils.StringEnum{Value: "dev"}, PROD: &utils.StringEnum{Value: "prod"}}
}

func IsDev() bool {
	return Config.Runmode == Runmode.DEV.Value
}

func IsProd() bool {
	return Config.Runmode == Runmode.PROD.Value
}
