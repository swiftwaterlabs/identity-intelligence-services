package repositories

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/configuration"
	"github.com/swiftwaterlabs/identity-intelligence-services/internal/pkg/models"
	"log"
	"strings"
)

type S3DirectoryObjectRepository struct {
	bucketName   string
	awsRegion    string
	fileUploader *s3manager.Uploader
}

func (r *S3DirectoryObjectRepository) init(config configuration.ConfigurationService) {
	r.bucketName = config.GetValue("identity_intelligence_prd_blob_store")
	r.awsRegion = config.GetValue("aws_region")
	r.fileUploader = r.initS3Uploader()
}

func (r *S3DirectoryObjectRepository) Save(item *models.DirectoryObject) error {

	input := &s3manager.UploadInput{
		Bucket:      aws.String(r.bucketName),
		Key:         aws.String(r.getFileKey(item)),
		Body:        strings.NewReader(item.Data),
		ContentType: aws.String("application/json"),
	}
	_, err := r.fileUploader.UploadWithContext(context.Background(), input)
	if err != nil {
		return err
	}

	return nil
}

func (r *S3DirectoryObjectRepository) getFileKey(item *models.DirectoryObject) string {
	path := r.resolveItemPath(item)
	identifier := r.resolveItemId(item)

	return fmt.Sprintf("%v/%v", path, identifier)
}

func (r *S3DirectoryObjectRepository) resolveItemPath(item *models.DirectoryObject) string {
	if item == nil || item.ObjectType == "" {
		return "unknown"
	}

	return strings.ToLower(item.ObjectType)
}

func (r *S3DirectoryObjectRepository) resolveItemId(item *models.DirectoryObject) string {
	if item == nil || item.Id == "" {
		return uuid.New().String()
	}

	return strings.ToLower(item.Id)
}

func (r *S3DirectoryObjectRepository) initS3Uploader() *s3manager.Uploader {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(r.awsRegion)},
	)
	if err != nil {
		log.Fatalf("failed to create AWS session, %v", err)
	}

	uploader := s3manager.NewUploader(sess)
	return uploader
}

func (r *S3DirectoryObjectRepository) Destroy() {

}
