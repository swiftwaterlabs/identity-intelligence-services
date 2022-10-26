package models

type Directory struct {
	Id                     string
	Name                   string
	Domain                 string
	Type                   string
	AuthenticationType     string
	ClientIdConfigName     string
	ClientSecretConfigName string
}
