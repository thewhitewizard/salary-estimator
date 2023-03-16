// Package model contains model engine
package model

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sajari/regression"
)

const (
	minDataModelLength = 10
	minCSVRowLength    = 4
)

var (
	errInvalidDataSet       = errors.New("invalid dataset format")
	errWhileOpenningDataSet = errors.New("error while openning dataset")
	errNotEnoughData        = errors.New("sorry, but no enough data to make a consistent prediction")
	errUnexpected           = errors.New("sorry, a really unexpected error has occurred")
)

type SalaryModel struct {
	engine *regression.Regression
}

func NewSalaryModel() *SalaryModel {
	s := &SalaryModel{
		engine: new(regression.Regression),
	}
	s.engine.SetObserved("Salary Prediction")
	s.engine.SetVar(0, "YearExperience")

	return s
}

func (s *SalaryModel) Add(salary, yearsExp float64) {
	s.engine.Train(regression.DataPoint(salary, []float64{yearsExp}))
}

func (s *SalaryModel) Predict(yearsExp float64) (float64, error) {
	return s.engine.Predict([]float64{yearsExp})
}

func (s *SalaryModel) LoadModelFromDataSetAndApplyFilter(datasetFilePath, filterTecho, filterCity, filterEducation string) error {
	file, err := os.OpenFile(datasetFilePath, os.O_RDONLY, os.ModePerm) //nolint:gosec

	if err != nil {
		return errWhileOpenningDataSet
	}

	defer func() {
		cerr := file.Close()
		if cerr != nil {
			log.Fatal(errUnexpected.Error())
		}
	}()

	sc := bufio.NewScanner(file)

	countDataAfterFilter := 0

	// read line by line and filter to prevent loading a huge dataset in memory
	for sc.Scan() {
		rowData := sc.Text()
		if rowData != "" {
			records := strings.Split(rowData, ",")
			if len(records) > minCSVRowLength { // Techno,YearsExperience,City,Education,SalaryMin,SalaryAverage
				techno := records[0]
				city := records[2]
				education := records[3]

				if strings.EqualFold(techno, filterTecho) &&
					strings.EqualFold(city, filterCity) &&
					strings.EqualFold(education, filterEducation) {
					yearsExp, errExp := strconv.ParseFloat(records[1], 64)
					salary, errSalary := strconv.ParseFloat(records[4], 64)

					if errExp != nil || errSalary != nil {
						return errInvalidDataSet
					}

					s.Add(salary, yearsExp)
					countDataAfterFilter++
				}
			} else {
				return errInvalidDataSet
			}
		}
	}

	if countDataAfterFilter < minDataModelLength {
		return errNotEnoughData
	}

	return s.engine.Run()
}
