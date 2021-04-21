package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)

type Env struct {
	Gobin  string
	Gopath string
}

func main() {
	var env Env
	_ = envconfig.Process("", &env)

	if env.Gobin != "" {
		_ = printFiles(env.Gobin)
		return
	}

	cmd, _ := exec.Command("go", "env", "GOBIN").Output()
	gobin := strings.TrimSuffix(string(cmd), "\n")
	if gobin != "" {
		_ = printFiles(gobin)
		return
	}

	if env.Gopath != "" {
		_ = printFiles(filepath.Join(env.Gopath, "bin"))
		return
	}

	cmd, _ = exec.Command("go", "env", "GOPATH").Output()
	gopath := strings.TrimSuffix(string(cmd), "\n")
	if gopath != "" {
		_ = printFiles(filepath.Join(gopath, "bin"))
		return
	}
}

func printFiles(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return nil
}
