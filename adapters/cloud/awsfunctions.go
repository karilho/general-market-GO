package cloud

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2/log"
	"os"
)

type S3StorageService struct {
	s3Client *s3.S3
}

func NewS3StorageService() *S3StorageService {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to AWS: %s", err.Error())
	}
	return &S3StorageService{
		s3Client: s3.New(sess),
	}

}

func (s *S3StorageService) CreateBucket(bucketName string) error {

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
	log.Info("Opening file")
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	log.Info("File: ", file, " opened successfully")

	log.Info("Uploading file to bucket")
	_, err = s.s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	log.Info("File uploaded successfully")
	return err
}
