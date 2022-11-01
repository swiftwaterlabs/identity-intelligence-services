package services

import (
	"strconv"
	"time"

	"github.com/go-ldap/ldap/v3"
	"github.com/google/uuid"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
)

func getUserAttributes() []string {
	return []string{
		"objectGUID",
		"sAMAccountName",
		"mail",
		"userPrincipalName",
		"givenName",
		"sn",
		"distinguishedName",
		"manager",
		"sAMAccountType",
		"company",
		"department",
		"whenCreated",
		"whenChanged",
		"logonCount",
		"badPwdCount",
		"badPasswordTime",
		"pwdLastSet",
		"lastLogon",
		"lastLogonTimestamp",
		"userAccountControl",
		"title",
		"memberOf",
	}
}

func getGroupAttributes() []string {
	return []string{
		"objectGUID",
		"sAMAccountName",
		"groupType",
		"distinguishedName",
		"member",
	}
}

func mapSearchResultToUser(directoryName string, result *ldap.Entry) *models.User {
	accountTypeRaw := result.GetAttributeValue("sAMAccountType")

	return &models.User{
		Id:            getObjectGuid(result).String(),
		Directory:     directoryName,
		ObjectType:    "User",
		Location:      result.GetAttributeValue("distinguishedName"),
		Upn:           result.GetAttributeValue("userPrincipalName"),
		Email:         result.GetAttributeValue("mail"),
		Name:          result.GetAttributeValue("sAMAccountName"),
		GivenName:     result.GetAttributeValue("givenName"),
		Surname:       result.GetAttributeValue("sn"),
		Manager:       result.GetAttributeValue("manager"),
		Type:          mapAccountTypeToDescription(accountTypeRaw),
		Company:       result.GetAttributeValue("company"),
		Department:    result.GetAttributeValue("department"),
		Status:        result.GetAttributeValue("userAccountControl"),
		Title:         result.GetAttributeValue("title"),
		CreatedAt:     getAttributeTimestamp(result, "whenCreated"),
		LastUpdatedAt: getAttributeTimestamp(result, "whenChanged"),
		CredentialInfo: &models.UserCredentialInfo{
			FailedLoginAttempts:    getAttributeInt(result, "badPwdCount"),
			LastFailedLoginAttempt: getAttributeDate(result, "badPasswordTime"),
			LastLogin:              getlastLogon(result),
			LoginCount:             getAttributeInt(result, "logonCount"),
			PasswordLastSet:        getAttributeDate(result, "pwdLastSet"),
		},
		GroupMembership: result.GetAttributeValues("memberOf"),
	}

}

func mapAccountTypeToDescription(accountType string) string {
	knownTypes := map[string]string{
		"268435456":  "SAM_GROUP_OBJECT",
		"268435457":  "SAM_NON_SECURITY_GROUP_OBJECT",
		"536870912":  "SAM_ALIAS_OBJECT",
		"536870913":  "SAM_NON_SECURITY_ALIAS_OBJECT",
		"805306368":  "SAM_NORMAL_USER_ACCOUNT",
		"805306369":  "SAM_MACHINE_ACCOUNT",
		"805306370":  "SAM_TRUST_ACCOUNT",
		"1073741824": "SAM_APP_BASIC_GROUP",
		"1073741825": "SAM_APP_QUERY_GROUP",
		"2147483647": "SAM_ACCOUNT_TYPE_MAX",
	}

	description, exists := knownTypes[accountType]
	if exists {
		return description
	}
	return accountType
}

func mapSearchResultToGroup(directoryName string, result *ldap.Entry) *models.Group {
	mappedType := mapGroupTypeToDescription(result.GetAttributeValue("groupType"))
	return &models.Group{
		Id:         getObjectGuid(result).String(),
		Directory:  directoryName,
		ObjectType: "Group",
		Location:   result.GetAttributeValue("distinguishedName"),
		Name:       result.GetAttributeValue("sAMAccountName"),
		Type:       mappedType,
		Members:    result.GetAttributeValues("member"),
	}
}

func mapGroupTypeToDescription(value string) string {
	knownTypes := map[string]string{
		"2":           "Global distribution group",
		"4":           "Domain Local distribution group",
		"8":           "Universal distribution group",
		"-2147483646": "Global security group",
		"-2147483644": "Domain Local security group",
		"-2147483640": "Universal security group",
	}

	description, exists := knownTypes[value]
	if exists {
		return description
	}
	return value
}

func getAttributeTimestamp(entry *ldap.Entry, name string) time.Time {
	value := entry.GetAttributeValue(name)
	date, _ := time.Parse("20060102150405.0Z07", value)

	return date
}

func getAttributeInt(entry *ldap.Entry, name string) int {
	value := entry.GetAttributeValue(name)
	parsed, _ := strconv.Atoi(value)

	return parsed
}

func getAttributeDate(entry *ldap.Entry, name string) time.Time {
	value := getAttributeInt(entry, name)
	seconds := value / 10000000            //Convert to seconds
	unixTimeStamp := seconds - 11644473600 // 1.1.1600 -> 1.1.1970 difference in seconds

	date := time.Unix(int64(unixTimeStamp), 0)
	return date
}

func getlastLogon(entry *ldap.Entry) time.Time {
	lastLogon := getAttributeDate(entry, "lastLogon")
	lastLogonTimeStamp := getAttributeDate(entry, "lastLogonTimestamp")

	if lastLogon.After(lastLogonTimeStamp) {
		return lastLogon
	}

	return lastLogonTimeStamp
}

func getObjectGuid(entry *ldap.Entry) uuid.UUID {
	value, _ := uuid.FromBytes(entry.GetRawAttributeValue("objectGUID"))

	return value
}
