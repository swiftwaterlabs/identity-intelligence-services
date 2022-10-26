package core

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/configuration"
)

func GetAwsSession(config configuration.ConfigurationService) *session.Session {

	region := config.GetValue("aws_region")

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewEnvCredentials(),
		},
	}))

	return session
}
