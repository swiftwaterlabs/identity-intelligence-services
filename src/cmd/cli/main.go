package main

import (
	"flag"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/configuration"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/orchestration"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/repositories"
	"log"
	"strings"
)

var (
	directoryArgument = flag.String("directory", "", "Directory to Search")
	objectArgument    = flag.String("object", "user", "Type of object to search.  Default is user")
)

func main() {
	configurationService := configuration.NewConfigurationService()
	directoryRepository := repositories.NewDirectoryRepository(configurationService)

	switch strings.ToLower(*objectArgument) {
	case "user":
		err := orchestration.ProcessUsers(*directoryArgument, configurationService, directoryRepository)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalln("Unrecognized object")
	}
}
