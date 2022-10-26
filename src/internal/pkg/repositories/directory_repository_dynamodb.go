package repositories

import (
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
)

func NewDynamoDbDirectoryRepository() *DynamoDbDirectoryRepository {
	return &DynamoDbDirectoryRepository{
		tableName: "identity-directories",
	}
}

type DynamoDbDirectoryRepository struct {
	tableName string
}

func (r *DynamoDbDirectoryRepository) GetAll() ([]*models.Directory, error) {
	return []*models.Directory{}, nil
}

func (r *DynamoDbDirectoryRepository) Get(identifier string) (*models.Directory, error) {

	return nil, nil
}
