package utils

import (
	"context"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// BucketBasics encapsulates the Amazon Simple Storage Service (Amazon S3) actions
// used in the examples.
// It contains S3Client, an Amazon S3 service client that is used to perform bucket
// and object actions.
type StemexS3Client struct {
	bucketName string
	s3Client   *s3.Client
}

// UploadFile reads from a file and puts the data into an object in a bucket.
func (b StemexS3Client) UploadFile(objectKey string, file io.Reader) error {
	// file, err := os.Open(fileName)
	// if err != nil {
	// 	log.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
	// } else {
	// 	defer file.Close()
	_, err := b.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(b.bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		// log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
		// 	fileName, b.bucketName, objectKey, err)
		return err
	}
	//}
	return nil
}

// ListObjects lists the objects in a bucket.
func (b StemexS3Client) ListObjects() ([]types.Object, error) {
	result, err := b.s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(b.bucketName),
	})
	var contents []types.Object
	if err != nil {
		log.Printf("Couldn't list objects in bucket %v. Here's why: %v\n", b.bucketName, err)
	} else {
		contents = result.Contents
	}
	return contents, err
}

// DownloadFile gets an object from a bucket and stores it in a local file.
func (b StemexS3Client) DownloadFile(objectKey string) ([]byte, error) {
	result, err := b.s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(b.bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", b.bucketName, objectKey, err)
		return nil, err
	}
	defer result.Body.Close()
	return io.ReadAll(result.Body)

	// file, err := os.Create(fileName)
	// if err != nil {
	// 	log.Printf("Couldn't create file %v. Here's why: %v\n", fileName, err)
	// 	return err
	// }
	// defer file.Close()
	// body, err := io.ReadAll(result.Body)
	// if err != nil {
	// 	log.Printf("Couldn't read object body from %v. Here's why: %v\n", objectKey, err)
	// }
	// _, err = file.Write(body)
	// return err
}

func NewStemexS3Client() *StemexS3Client {
	opts := [](func(*config.LoadOptions) error){
		config.WithRegion("ap-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("AKIAXTXB2EEWGOK6IDNV", "gZy0G+tR61m3rDr4tgC0iBigGYQPeQ6LRKIj94PI", "")),
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), opts...)
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	return &StemexS3Client{bucketName: "stemexhub-files", s3Client: s3.NewFromConfig(cfg)}
}
