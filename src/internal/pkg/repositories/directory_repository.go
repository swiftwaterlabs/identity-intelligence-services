package repositories

import (
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/configuration"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
)

type DirectoryRepository interface {
	GetAll() ([]*models.Directory, error)
	Get(identifier string) (*models.Directory, error)
}

func NewDirectoryRepository(configurationService configuration.ConfigurationService) DirectoryRepository {
	return NewDynamoDbDirectoryRepository(configurationService)
}
