package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// Exit outputs error message in standerd error
func Exit(err error, codes ...int) {
	var code int
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 2
	}
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func getDefaultHistoryPath() (path string, err error) {
	var d string
	if d, err = homedir.Dir(); err != nil {
		return "", err
	}
	path = filepath.Join(d, ".bash_history")
	if !existFile(path) {
		return "", fmt.Errorf(".bash_history dose not exist")
	}
	return
}

func existFile(path string) bool {
	src, err := os.Stat(path)
	if err != nil {
		return false
	}

	if !src.IsDir() {
		return true
	}
	return false
}
