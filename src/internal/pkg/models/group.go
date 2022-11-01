package models

type Group struct {
	Id         string
	Directory  string
	ObjectType string
	Location   string
	Type       string
	Name       string
	Members    []string
}
