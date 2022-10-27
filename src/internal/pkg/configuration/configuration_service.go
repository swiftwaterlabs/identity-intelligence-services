package configuration

type ConfigurationService interface {
	GetValue(name string) string
	GetSecret(name string) string
}

func NewConfigurationService() ConfigurationService {
	return &environmentConfigurationService{}
}
