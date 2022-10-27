package models

type Group struct {
	Id         string
	ObjectType string
	Location   string
	Type       string
	Name       string
	Members    []string
}
