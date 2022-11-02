package orchestration

import (
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/repositories"
	"strings"
)

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
