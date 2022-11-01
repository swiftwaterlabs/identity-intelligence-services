package models

type Group struct {
	Id            string
	DirectoryName string
	ObjectType    string
	Location      string
	Type          string
	Name          string
	Members       []string
}
