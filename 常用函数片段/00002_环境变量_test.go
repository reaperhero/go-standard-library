package command

import (
	"log"
	"os"
	"testing"
)

// get env
func GetEnvWithDefault(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}

func Test_Get_Env(t *testing.T) {
	logpath := GetEnvWithDefault("log", "/data")
	log.Println(logpath)
}
