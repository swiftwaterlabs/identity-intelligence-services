package configuration

import "os"

type environmentConfigurationService struct {
}

func (s *environmentConfigurationService) GetValue(name string) string {
	return os.Getenv(name)
}

func (s *environmentConfigurationService) GetSecret(name string) string {
	return os.Getenv(name)
}
