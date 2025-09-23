package awssettings

import (
	"context"
	"fmt"
	"yadhronics-blog/settings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var AWSClient *s3.Client
var PreSignedURL *s3.PresignClient

func InitializeS3Client() {
	awsCfg, err := CreateAWSConfig()
	if err != nil {
		fmt.Printf("unable to load AWS config: %v", err)
	}

	AWSClient = s3.NewFromConfig(awsCfg)
	PreSignedURL = s3.NewPresignClient(AWSClient)
}

func CreateAWSConfig() (aws.Config, error) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(settings.Config.AWS.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			settings.Config.AWS.AccessKey,
			settings.Config.AWS.SecretKey,
			"",
		)),
	)
	if err != nil {
		return aws.Config{}, fmt.Errorf("unable to load AWS config: %w", err)
	}

	return awsCfg, nil
}
