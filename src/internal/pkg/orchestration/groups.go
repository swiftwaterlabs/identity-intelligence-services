package orchestration

import (
	"errors"
	"fmt"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/configuration"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/core"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/messaging"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/repositories"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/services"
	"log"
	"sync"
)

func ExtractGroups(directoryName string,
	configurationService configuration.ConfigurationService,
	directoryRepository repositories.DirectoryRepository,
	userDataHub messaging.MessageHub) error {
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
			processingErr := processDirectoryGroups(directory, config, userDataHub)
			if processingErr != nil {
				processingErrors[item.Name] = processingErr
			}
		}(item, configurationService)
	}
	awaiter.Wait()

	return core.ConsolidateErrorMap(processingErrors)
}

func processDirectoryGroups(directory *models.Directory,
	configuration configuration.ConfigurationService,
	dataHub messaging.MessageHub) error {
	directoryService, err := services.NewDirectoryService(directory, configuration)
	if err != nil {
		return err
	}

	defer directoryService.Close()

	publishingQueue := configuration.GetValue("identity_intelligence_prd_ingestion_queue")

	counter := 0
	log.Printf("Sending groups to %s", publishingQueue)
	handler := func(data []*models.Group) {
		toPublish := core.ToInterfaceSlice(data)
		err := dataHub.SendBulk(toPublish, publishingQueue)
		if err != nil {
			log.Println(err)
		}
		counter += len(data)
		log.Printf("Processed %v groups in %s", counter, directory.Name)
	}

	err = directoryService.HandleGroups(handler)
	return err

}
