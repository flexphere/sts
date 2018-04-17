package s3

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/flexphere/sts/settings"
)

func (s *S3Client) Upload(key string, bin []byte) error {
	_, err := s.svc.PutObject(&s3.PutObjectInput{
		Bucket: settings.Settings.S3_BUCKET_NAME,
		Key:    aws.String(key),
		Body:   bytes.NewReader(bin),
	})
	if err != nil {
		return err
	}
	return nil
}
