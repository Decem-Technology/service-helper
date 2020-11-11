package bootstrap

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var s3session *session.Session

// S3 aws s3 helper Please enter AWS_REGION, AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY value to env
type S3 struct {
}

// CreateS3Connection make s3 connection
func CreateS3Connection() {
	var err error = nil
	s3session, err = session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		panic(fmt.Sprintf("[S3] s3 connection error : %s", err))
	}
	fmt.Println("[S3] connected")
}

// Session get s3 connection session
func (ctl *S3) Session() *session.Session {
	return s3session
}

// Uploader get uploader helper
func (ctl *S3) Uploader() *s3manager.Uploader {
	return s3manager.NewUploader(s3session)
}

// Downloader get downloer helper
func (ctl *S3) Downloader() *s3manager.Downloader {
	return s3manager.NewDownloader(s3session)
}

// Service get s3 service
func (ctl *S3) Service() *s3.S3 {
	return s3.New(s3session)
}
