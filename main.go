package main

import (
	"errors"
	"fmt"
	"os"
	"salary-estimator/internal/model"
	"salary-estimator/internal/utils"
	"strings"
)

var (
	inDirectory      string
	datasetName      string
	outDirectory     string
	jobSecret        string
	citySecret       string
	educationSecret  string
	experienceSecret float64

	errCantReadSecret error = errors.New("can't read or missing secret")
)

func init() {
	inDirectory = os.Getenv("IEXEC_IN")
	datasetName = os.Getenv("IEXEC_DATASET_FILENAME")
	outDirectory = os.Getenv("IEXEC_OUT")

	jobSecret = utils.GetStringSecret(1)        //nolint: gomnd
	citySecret = utils.GetStringSecret(2)       //nolint: gomnd
	educationSecret = utils.GetStringSecret(3)  //nolint: gomnd
	experienceSecret = utils.GetNumberSecret(4) //nolint: gomnd

	if jobSecret == "" || citySecret == "" || educationSecret == "" || experienceSecret < 0 {
		utils.CheckOrRaiseError(outDirectory, errCantReadSecret)
	}

	if !strings.HasSuffix(inDirectory, "/") {
		inDirectory += "/"
	}

	if !strings.HasSuffix(outDirectory, "/") {
		outDirectory += "/"
	}

	_, err := os.Stat(inDirectory + datasetName)
	utils.CheckOrRaiseError(outDirectory, err)
}

func main() {
	model := model.NewSalaryModel()
	err := model.LoadModelFromDataSetAndApplyFilter(inDirectory+datasetName, jobSecret, citySecret, educationSecret)
	utils.CheckOrRaiseError(outDirectory, err)

	prediction, err := model.Predict(experienceSecret)
	utils.CheckOrRaiseError(outDirectory, err)
	utils.CompleteTheTask(outDirectory, []byte(fmt.Sprint("you can expect a salary of ", prediction)))
}
