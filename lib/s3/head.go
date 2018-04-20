package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/flexphere/sts/settings"
)

func (s *S3Client) Head(key string) (*s3.HeadObjectOutput, error) {
	result, err := s.svc.HeadObject(&s3.HeadObjectInput{
		Bucket: settings.Settings.S3_BUCKET_NAME,
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
