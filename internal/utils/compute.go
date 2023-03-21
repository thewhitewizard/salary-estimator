// Package utils contains some method to respect iExec expected output
package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	rwmode                  = 0666
	iExecSecretBaseVariable = "IEXEC_REQUESTER_SECRET_" // nolint: gosec
)

// CompleteTheTask complete the task by writing content in result file
func CompleteTheTask(outDirectory string, content []byte) {
	writeResultFile(outDirectory, content)
	writeComputedFile(outDirectory)
	exit()
}

func GetStringSecret(index int) string {
	if index < 1 {
		return ""
	}

	return strings.TrimSpace(os.Getenv(iExecSecretBaseVariable + strconv.Itoa(index)))
}

func GetNumberSecret(index int) float64 {
	value := GetStringSecret(index)
	if value == "" {
		return -1
	}

	fValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return -1
	}

	return fValue
}

// CheckOrRaiseError check if err is nil, if not write the error as result and stop the task
func CheckOrRaiseError(outDirectory string, err error) {
	if err != nil {
		CompleteTheTask(outDirectory, []byte(err.Error()))
	}
}

func writeResultFile(outDirectory string, content []byte) {
	fmt.Println("writing the result in", outDirectory+"result.txt")
	err := os.WriteFile(outDirectory+"result.txt", content, rwmode) //nolint:gosec

	if err != nil {
		log.Fatalln(err)
	}
}

func writeComputedFile(outDirectory string) {
	fmt.Println("writing the proof of calculation", outDirectory+"computed.json")
	err := os.WriteFile(outDirectory+"computed.json",
		[]byte(`{"deterministic-output-path": "`+outDirectory+`result.txt"}`), rwmode) //nolint:gosec

	if err != nil {
		log.Fatalln(err)
	}
}

func exit() {
	os.Exit(0)
}
