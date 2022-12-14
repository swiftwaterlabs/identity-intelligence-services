package core

import (
	"errors"
	"fmt"
	"strings"
)

func ConsolidateErrorMap(errorMap map[string]error) error {
	if len(errorMap) == 0 {
		return nil
	}

	errorMessages := make([]string, 0)
	for _, value := range errorMap {
		errorMessages = append(errorMessages, value.Error())
	}

	return errors.New(strings.Join(errorMessages, "\n"))
}

func ConsolidateErrors(toMap []error) error {
	combinedMessage := ""
	for _, item := range toMap {
		combinedMessage = fmt.Sprintf("%v%v;", combinedMessage, item.Error())
	}

	return errors.New(combinedMessage)
}
