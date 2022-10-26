package models

type Directory struct {
	Id                     string
	Name                   string
	Host                   string
	Base                   string
	Type                   string
	AuthenticationType     string
	ClientIdConfigName     string
	ClientSecretConfigName string
}
