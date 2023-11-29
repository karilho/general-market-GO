package domain

type StorageService interface {
	CreateBucket(bucketName string) error
}
