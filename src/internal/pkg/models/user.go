package models

import "time"

type User struct {
	Id              string
	ObjectType      string
	Location        string
	Upn             string
	Name            string
	Email           string
	GivenName       string
	Surname         string
	Manager         string
	Type            string
	Company         string
	Department      string
	Status          string
	Title           string
	CredentialInfo  *UserCredentialInfo
	GroupMembership []string

	CreatedAt     time.Time
	LastUpdatedAt time.Time
}

type UserCredentialInfo struct {
	FailedLoginAttempts    int
	LastFailedLoginAttempt time.Time
	LastLogin              time.Time
	LoginCount             int
	PasswordLastSet        time.Time
}
