package services

import (
	"errors"
	"fmt"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/configuration"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
	"strings"
)

type DirectoryService interface {
	HandleUsers(action func([]*models.User)) error
	HandleGroups(action func([]*models.Group)) error
	Close()
}

func NewDirectoryService(directory *models.Directory, configurationService configuration.ConfigurationService) (DirectoryService, error) {
	switch strings.ToLower(directory.Type) {
	case "ldap":
		return NewLdapDirectoryService(directory, configurationService)
	default:
		return nil, errors.New(fmt.Sprintf("unrecognized directory type:%s", directory.Type))
	}
}
