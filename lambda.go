package main

/*
import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"io"
	"log"
)

/// THIS FILE IS JUST IN LAMBDA TO TEST THE FUNCTIONALITY.




type Message struct {
	BuyerID      int  `json:"buyer_id"`
	UserDataID   int  `json:"user_data_id"`
	HasPurchased bool `json:"has_purchased"`
}

func testLambda(ctx context.Context, s3Event events.S3Event) error {
	log.Println("Starting testLambda function")

	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("Failed to load SDK config: %s", err.Error())
		return err
	}

	s3Client := s3.NewFromConfig(sdkConfig)
	sqsClient := sqs.NewFromConfig(sdkConfig)

	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		log.Println("Getting object from bucket")
		resp, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &key,
		})
		if err != nil {
			log.Printf("Failed to get object: %s", err.Error())
			return err
		}
		defer resp.Body.Close()

		log.Println("Reading object data")
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read object data: %s", err.Error())
			return err
		}

		log.Println("Unmarshalling JSON data")
		var message Message
		err = json.Unmarshal(body, &message)
		if err != nil {
			log.Printf("Failed to unmarshal JSON: %s", err.Error())
			return err
		}

		log.Printf("JSON data: %v\n", message)

		if !message.HasPurchased {
			queueURL := "GET THEM FROM SQS"
			log.Println("Sending message to SQS queue")
			_, err = sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
				QueueUrl:    &queueURL,
				MessageBody: aws.String(string(body)),
			})
			if err != nil {
				log.Printf("Failed to send message to SQS queue: %s", err.Error())
				return err
			}
		}
	}

	log.Println("Finished testLambda function")
	return nil
}

func main() {
	lambda.Start(testLambda)
}

*/
