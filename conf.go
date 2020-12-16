package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

func init() {
	// Debug mode configs exposure
	readDebugDotEnv()
}

func resolveCnfgPath(relativePath string) string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), relativePath)
}

func readDebugDotEnv() {
	envFilePath := resolveCnfgPath("./tmp/.env")
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
