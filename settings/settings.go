package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var environments = map[string]string{
	"production"  : "settings/prod.json",
	"development" :	"settings/dev.json",
	"tests"		  :	"../../settings/test.json"
}

type Settings struct {
	PrivateKeyPath string
	PublicKeyPath  string
	JWTExpirationDelta int
}

var settings Settings = Settings{}

var env = "development"

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Alerta, modo de desarrollo activado por falta de GO_ENV")
		env = "development"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error al leer archivo de configuracion", err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error al parsear config", jsonErr)
	}
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}

func IsTestEnvironment() bool {
	return env == "tests"
}
