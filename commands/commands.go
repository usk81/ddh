package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "ddh",
		Short: "delete duplicates in .bash_history",
		Long:  "delete duplicates in .bash_history",
		Run:   deleteCommand,
	}
)

// Run executes commands
func Run() {
	RootCmd.Execute()
}

func deleteCommand(cmd *cobra.Command, args []string) {
	path, err := getDefaultHistoryPath()
	if err != nil {
		Exit(err, 1)
	}
	if err = removeDuplicateHistory(path); err != nil {
		Exit(err, 1)
	}
}

func removeDuplicateHistory(path string) (err error) {
	var contents []byte
	if contents, err = ioutil.ReadFile(path); err != nil {
		return
	}
	if contents, err = mergeDuplicateHistoryData(contents); err != nil {
		return
	}
	return ioutil.WriteFile(path, contents, os.ModePerm)
}

func mergeDuplicateHistoryData(contents []byte) (result []byte, err error) {
	lines := regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(contents), -1)
	et := regexp.MustCompile(`#[0-9]+`)
	cnt := len(lines)
	if cnt < 2 {
		return nil, fmt.Errorf("history doesn't exists")
	}
	cnt = cnt - 1
	exists := map[string]string{}
	var ts string
	var buffer bytes.Buffer
	for k, v := range lines {
		if k == cnt {
			break
		}
		if k%2 == 0 {
			if !et.MatchString(v) {
				return nil, fmt.Errorf("invalid format; line: %d, value: %s", k, v)
			}
			ts = v
		} else {
			if _, ok := exists[v]; !ok {
				exists[v] = ts
				buffer.WriteString(ts + "\n" + v + "\n")
			}
		}
	}
	result = buffer.Bytes()
	return
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
