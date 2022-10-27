package orchestration

import (
	"errors"
	"fmt"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/configuration"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/core"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/repositories"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/services"
	"log"
	"strings"
	"sync"
)

func ExtractUsers(directoryName string,
	configurationService configuration.ConfigurationService,
	directoryRepository repositories.DirectoryRepository) error {

	directories, err := getDirectories(directoryName, directoryRepository)
	if err != nil {
		return err
	}
	if len(directories) == 0 {
		return errors.New(fmt.Sprintf("no directory found with identifier '%v'", directoryName))
	}

	processingErrors := make(map[string]error, 0)
	var awaiter sync.WaitGroup
	for _, item := range directories {
		awaiter.Add(1)

		go func(directory *models.Directory, config configuration.ConfigurationService) {
			defer awaiter.Done()
			processingErr := processDirectoryUsers(directory, config)
			if processingErr != nil {
				processingErrors[item.Name] = processingErr
			}
		}(item, configurationService)
	}
	awaiter.Wait()

	return core.ConsolidateErrorMap(processingErrors)

}

func getDirectories(directoryName string,
	directoryRepository repositories.DirectoryRepository) ([]*models.Directory, error) {
	if strings.EqualFold(directoryName, "") {
		return directoryRepository.GetAll()
	}

	result, err := directoryRepository.Get(directoryName)
	if err != nil {
		return make([]*models.Directory, 0), err
	}

	return []*models.Directory{result}, nil
}

func processDirectoryUsers(directory *models.Directory, configuration configuration.ConfigurationService) error {
	directoryService, err := services.NewDirectoryService(directory, configuration)
	if err != nil {
		return err
	}

	defer directoryService.Close()

	counter := 0
	handler := func(data []*models.User) {
		for _, item := range data {
			core.MapToJson(item)
			counter++
		}
		log.Printf("Prorcessed %v users in %s", counter, directory.Name)
	}

	err = directoryService.HandleUsers(handler)
	return err

}
