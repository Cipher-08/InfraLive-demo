package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func main() {
	// Get user input
	var accessKey, secretKey, bucketName, region string
	fmt.Print("Enter AWS Access Key ID: ")
	fmt.Scanln(&accessKey)
	fmt.Print("Enter AWS Secret Access Key: ")
	fmt.Scanln(&secretKey)
	fmt.Print("Enter S3 Bucket Name: ")
	fmt.Scanln(&bucketName)
	fmt.Print("Enter AWS Region: ")
	fmt.Scanln(&region)

	// Load AWS configuration with user-provided credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		log.Fatalf("Failed to load AWS configuration: %v", err)
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// Create the S3 bucket
	if region == "us-east-1" {
		// For us-east-1, do not include CreateBucketConfiguration
		_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		})
	} else {
		// For all other regions, include CreateBucketConfiguration
		_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
			CreateBucketConfiguration: &types.CreateBucketConfiguration{
				LocationConstraint: types.BucketLocationConstraint(region),
			},
		})
	}

	if err != nil {
		log.Fatalf("Failed to create S3 bucket: %v", err)
	}

	fmt.Printf("Bucket '%s' created successfully in region '%s'.\n", bucketName, region)
}