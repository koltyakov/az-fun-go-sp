//+build !prod

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

// Debug mode configs exposure
// binds Azure Functions local settings and FUNCTIONS_CUSTOMHANDLER_PORT via .env

func init() {
	log.Println("Debug mode is initiated")
	readDebugDotEnv("./functions/tmp/.env")
	readDebugDotEnv("./tmp/.env")
}

func resolveCnfgPath(relativePath string) string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), relativePath)
}

func readDebugDotEnv(dotEnvPath string) {
	envFilePath := resolveCnfgPath(dotEnvPath)
	envFile, err := os.Open(envFilePath)
	if err != nil {
		return
	}
	defer func() { _ = envFile.Close() }()

	byteValue, _ := ioutil.ReadAll(envFile)
	keyVals := strings.Split(fmt.Sprintf("%s", byteValue), "\n")
	for _, keyVal := range keyVals {
		kv := strings.SplitN(keyVal, "=", 2)
		if len(kv) == 2 {
			if _, ok := os.LookupEnv(kv[0]); !ok {
				_ = os.Setenv(kv[0], kv[1])
			}
		}
	}
}
