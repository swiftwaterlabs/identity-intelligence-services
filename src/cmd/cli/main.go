package main

import (
	"flag"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/configuration"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/messaging"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/orchestration"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/repositories"
	"log"
	"os"
	"strings"
)

var (
	directoryArgument = flag.String("directory", "", "Directory to Search")
	objectArgument    = flag.String("object", "user", "Type of object to search.  Default is user")
)

func main() {
	appConfig := &configuration.AppConfig{
		AwsRegion: os.Getenv("aws_region"),
	}
	configurationService := configuration.NewConfigurationService(appConfig)
	directoryRepository := repositories.NewDirectoryRepository(appConfig, configurationService)
	messageHub := messaging.NewMessageHub(appConfig)

	switch strings.ToLower(*objectArgument) {
	case "user":
		err := orchestration.ExtractUsers(*directoryArgument, configurationService, directoryRepository, messageHub)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalln("Unrecognized object")
	}
}
