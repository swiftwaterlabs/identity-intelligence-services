package services

import (
	"errors"
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
	if action == nil {
		return errors.New("no action defined")
	}

	filterCriteria := "(&(objectClass=user))"
	fields := getUserAttributes()

	processor := func(items []*ldap.Entry) {
		data := make([]*models.User, 0)
		for _, item := range items {
			user := MapSearchResultToUser(item)
			data = append(data, user)
		}
		action(data)
	}

	return s.searchWithAction(filterCriteria, fields, processor)
}
func (s *ldapDirectoryService) HandleGroups(action func([]*models.Group)) error {
	return nil
}
func (s *ldapDirectoryService) Close() {
	s.connection.Close()
}

func (s *ldapDirectoryService) searchWithAction(filter string, fields []string, action func([]*ldap.Entry)) error {

	pagingControl := &ldap.ControlPaging{PagingSize: 100}
	searchRequest := ldap.NewSearchRequest(
		s.baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter, // The filter to apply
		fields, // A list attributes to retrieve
		[]ldap.Control{pagingControl},
	)

	for {
		result, err := s.connection.Search(searchRequest)
		if err != nil {
			return err
		}
		if result == nil {
			return ldap.NewError(ldap.ErrorNetwork, errors.New("ldap: packet not received"))
		}

		action(result.Entries)

		pagingResult := ldap.FindControl(result.Controls, ldap.ControlTypePaging)
		if pagingResult == nil {
			pagingControl = nil
			break
		}

		cookie := pagingResult.(*ldap.ControlPaging).Cookie
		if len(cookie) == 0 {
			pagingControl = nil
			break
		}
		pagingControl.SetCookie(cookie)
	}

	return nil
}
