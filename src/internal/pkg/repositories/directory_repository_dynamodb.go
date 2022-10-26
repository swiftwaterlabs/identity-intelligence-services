package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/configuration"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/core"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
)

func NewDynamoDbDirectoryRepository(configurationService configuration.ConfigurationService) *DynamoDbDirectoryRepository {
	session := core.GetAwsSession(configurationService)
	client := dynamodb.New(session)

	return &DynamoDbDirectoryRepository{
		tableName: "identity-directories",
		client:    client,
	}
}

type DynamoDbDirectoryRepository struct {
	tableName string
	client    *dynamodb.DynamoDB
}

func (r *DynamoDbDirectoryRepository) GetAll() ([]*models.Directory, error) {
	result := make([]*models.Directory, 0)
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}
	queryResult, err := r.client.Scan(scanInput)
	if err != nil {
		return result, err
	}

	for _, item := range queryResult.Items {
		directory := r.mapItemToDirectory(item)
		result = append(result, directory)
	}

	return result, nil
}

func (r *DynamoDbDirectoryRepository) Get(identifier string) (*models.Directory, error) {
	itemInput := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(identifier),
			},
		},
		TableName: aws.String(r.tableName),
	}
	queryResult, err := r.client.GetItem(itemInput)
	if err != nil {
		return nil, err
	}

	result := r.mapItemToDirectory(queryResult.Item)
	return result, nil
}

func (r *DynamoDbDirectoryRepository) mapItemToDirectory(item map[string]*dynamodb.AttributeValue) *models.Directory {
	return &models.Directory{
		Id:                     item["Id"].String(),
		Name:                   item["Name"].String(),
		Host:                   item["Host"].String(),
		Base:                   item["Base"].String(),
		Type:                   item["Type"].String(),
		AuthenticationType:     item["AuthenticationType"].String(),
		ClientIdConfigName:     item["ClientIdConfigName"].String(),
		ClientSecretConfigName: item["ClientSecretConfigName"].String(),
	}
}
