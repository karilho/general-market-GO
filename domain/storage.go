package domain

type StorageService interface {
	CreateBucket(bucketName string) error
	UploadFileToBucket(bucketName, fileName string) error
}
