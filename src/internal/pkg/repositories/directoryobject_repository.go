package repositories

import (
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/configuration"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
)

type DirectoryObjectRepository interface {
	Save(item *models.DirectoryObject) error
	Destroy()
}

func NewDirectoryObjectRepository(config configuration.ConfigurationService) DirectoryObjectRepository {
	instance := &S3DirectoryObjectRepository{}
	instance.init(config)

	return instance
}
