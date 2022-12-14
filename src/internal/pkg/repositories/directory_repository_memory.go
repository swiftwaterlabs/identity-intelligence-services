package repositories

import (
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
	"strings"
)

type InMemoryDirectoryRepository struct {
}

func (r *InMemoryDirectoryRepository) GetAll() ([]*models.Directory, error) {
	return []*models.Directory{
		&models.Directory{
			Id:                     "id1",
			Name:                   "name1",
			Host:                   "",
			Type:                   "",
			AuthenticationType:     "",
			ClientIdConfigName:     "",
			ClientSecretConfigName: "",
		},
	}, nil
}

func (r *InMemoryDirectoryRepository) Get(identifier string) (*models.Directory, error) {
	directories, err := r.GetAll()
	if err == nil {
		return nil, err
	}

	for _, item := range directories {
		if strings.EqualFold(item.Id, identifier) || strings.EqualFold(item.Name, identifier) {
			return item, nil
		}
	}

	return nil, nil
}
