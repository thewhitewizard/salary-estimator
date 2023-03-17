package main

import (
	"errors"
	"fmt"
	"os"
	"salary-estimator/internal/model"
	"salary-estimator/internal/utils"
	"strconv"
	"strings"
)

const (
	nbArgs = 4
)

var (
	inDirectory  string
	datasetName  string
	outDirectory string

	errMissingArguments = errors.New("missing arguments provided")
	errBadArguments     = errors.New("bad arguments provided")
)

func init() {
	inDirectory = os.Getenv("IEXEC_IN")
	datasetName = os.Getenv("IEXEC_DATASET_FILENAME")
	outDirectory = os.Getenv("IEXEC_OUT")

	if !strings.HasSuffix(inDirectory, "/") {
		inDirectory += "/"
	}

	_, err := os.Stat(inDirectory + datasetName)
	utils.CheckOrRaiseError(outDirectory, err)
}

func main() {
	if len(os.Args) < nbArgs {
		utils.CheckOrRaiseError(outDirectory, errMissingArguments)
	}

	job := os.Args[1]
	city := os.Args[2]
	education := os.Args[3]
	exp, err := strconv.ParseFloat(os.Args[4], 64)

	if err != nil {
		utils.CheckOrRaiseError(outDirectory, errBadArguments)
	}

	model := model.NewSalaryModel()
	err = model.LoadModelFromDataSetAndApplyFilter(inDirectory+datasetName, job, city, education)
	utils.CheckOrRaiseError(outDirectory, err)

	prediction, err := model.Predict(exp)
	utils.CheckOrRaiseError(outDirectory, err)
	utils.CompleteTheTask(outDirectory, []byte(fmt.Sprint("you can expect a salary of ", prediction)))
}
