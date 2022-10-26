package configuration

import "os"

type ConfigurationService interface {
	GetValue(name string) string
	GetSecret(name string) string
}

func NewConfigurationService() ConfigurationService {
	return &environmentConfigurationService{}
}

type environmentConfigurationService struct {
}

func (s *environmentConfigurationService) GetValue(name string) string {
	return os.Getenv(name)
}

func (s *environmentConfigurationService) GetSecret(name string) string {
	return os.Getenv(name)
}
