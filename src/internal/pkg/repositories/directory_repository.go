package repositories

import (
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
)

type DirectoryRepository interface {
	GetAll() ([]*models.Directory, error)
	Get(identifier string) (*models.Directory, error)
}

func NewDirectoryRepository() DirectoryRepository {
	return &InMemoryDirectoryRepository{}
}
