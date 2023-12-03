package domain

type StorageService interface {
	CreateBucket(bucketName string) error
	UploadFileToBucket(bucketName, fileName string) error
	CreateQueue(queueName string) error
	SendMessage(queueName string, messageBody string) error
	ListMessages(queueName string) error
}
