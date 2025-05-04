package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {

	data, err := os.ReadFile(fileName)
	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}

	balanceText := string(data)
	value, convertError := strconv.ParseFloat(balanceText, 64)

	if convertError != nil {
		return 1000, errors.New("failed to parse stored  value")
	}

	return value, nil

}

func WriteFloatToFile(value float64, fileName string) {

	valueText := fmt.Sprint(value)

	os.WriteFile(fileName, []byte(valueText), 0644)
}
