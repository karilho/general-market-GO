package cloud

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"path/filepath"
)

type S3StorageService struct {
	s3Client  *s3.S3
	sqsClient *sqs.SQS
}

func NewS3StorageService() *S3StorageService {
	log.Info("Connecting to AWS")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to AWS: %s", err.Error())
	}
	log.Info("Connection to AWS sucessfully")
	return &S3StorageService{
		s3Client:  s3.New(sess),
		sqsClient: sqs.New(sess),
	}
}

func (s *S3StorageService) CreateBucket(bucketName string) error {

	log.Info("Starting bucket creation")

	_, err := s.s3Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err == nil {
		log.Error("Bucket already exists, skipping creation")
		return nil
	}

	_, err = s.s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatalf("Failed to create bucket: %s", err.Error())
	}

	err = s.s3Client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatalf("Failed to wait for bucket to exist: %s", err.Error())
	}

	log.Info("Bucket created successfully")
	return nil
}

func (s *S3StorageService) UploadFileToBucket(bucketName, fileName string) error {
	log.Info("Opening file", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	os.Remove(fileName)

	log.Info("File: ", file, " opened successfully")

	// Imprime o formato do arquivo
	log.Debug("File format: ", filepath.Ext(fileName))
	fmt.Println("File format: ", filepath.Ext(fileName))

	log.Info("Uploading file to bucket")
	_, err = s.s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	log.Info("File uploaded successfully")
	return err
}

func (s *S3StorageService) CreateQueue(queueName string) error {
	log.Info("Starting queue creation")

	params := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("60"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	}
	_, err := s.sqsClient.CreateQueue(params)
	if err != nil {
		log.Fatalf("Failed to create queue: %s", err.Error())
		return err
	}

	log.Info("Queue created successfully")
	return nil
}

func (s *S3StorageService) SendMessage(queueName string, messageBody string) error {
	log.Info("Starting to send message")

	sendMessageInput := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueName),
		MessageBody: aws.String(messageBody),
	}

	_, err := s.sqsClient.SendMessage(sendMessageInput)
	if err != nil {
		log.Fatalf("Failed to send message: %s", err.Error())
		return err
	}

	log.Info("Message sent successfully")
	return nil
}

func (s *S3StorageService) ListMessages(queueName string) error {
	log.Info("Starting to list messages")

	receiveMessageInput := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueName),
		MaxNumberOfMessages: aws.Int64(10), // You can adjust this value
		VisibilityTimeout:   aws.Int64(0),  // This ensures messages are not removed
	}

	result, err := s.sqsClient.ReceiveMessage(receiveMessageInput)
	if err != nil {
		log.Fatalf("Failed to list messages: %s", err.Error())
		return err
	}

	for _, message := range result.Messages {
		log.Info("Message: ", *message.Body)
	}

	log.Info("Messages listed successfully")
	return nil
}
