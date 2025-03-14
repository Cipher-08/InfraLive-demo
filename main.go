package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Request struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
	Region    string `json:"region"`
}

func createBucket(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Load AWS configuration with user-provided credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(req.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(req.AccessKey, req.SecretKey, "")),
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to load AWS configuration: %v", err), http.StatusInternalServerError)
		return
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// Create the S3 bucket
	if req.Region == "us-east-1" {
		// For us-east-1, do not include CreateBucketConfiguration
		_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
			Bucket: aws.String(req.Bucket),
		})
	} else {
		// For all other regions, include CreateBucketConfiguration
		_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
			Bucket: aws.String(req.Bucket),
			CreateBucketConfiguration: &types.CreateBucketConfiguration{
				LocationConstraint: types.BucketLocationConstraint(req.Region),
			},
		})
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create S3 bucket: %v", err), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Bucket '%s' created successfully in region '%s'.", req.Bucket, req.Region)))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Define the API endpoint with CORS enabled
	http.Handle("/create-bucket", enableCORS(http.HandlerFunc(createBucket)))

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}