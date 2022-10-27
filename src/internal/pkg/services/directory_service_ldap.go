package services

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/configuration"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
	"log"
)

type ldapDirectoryService struct {
	baseDN     string
	connection *ldap.Conn
}

func NewLdapDirectoryService(directory *models.Directory, configurationService configuration.ConfigurationService) (DirectoryService, error) {
	userName := configurationService.GetSecret(directory.ClientIdConfigName)
	password := configurationService.GetSecret(directory.ClientSecretConfigName)

	return &ldapDirectoryService{
		baseDN:     directory.Base,
		connection: getLdapConnection(directory.Host, userName, password),
	}, nil
}

func getLdapConnection(host string, userName string, password string) *ldap.Conn {
	address := fmt.Sprintf("ldaps://%v", host)
	conn, err := ldap.DialURL(address)
	if err != nil {
		log.Fatalf("Failed to connect: %s\n", err)
	}

	err = conn.Bind(userName, password)
	if err != nil {
		log.Fatalf("Failed to bind: %s\n", err)
	}

	return conn
}

func (s *ldapDirectoryService) HandleUsers(action func([]*models.User)) error {
	return nil
}
func (s *ldapDirectoryService) HandleGroups(action func([]*models.Group)) error {
	return nil
}
func (s *ldapDirectoryService) Close() {

}
