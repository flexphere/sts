package s3

import (
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/flexphere/sts/settings"
)

func (s *S3Client) Download(key string) ([]byte, error) {
	result, err := s.svc.GetObject(&s3.GetObjectInput{
		Bucket: settings.Settings.S3_BUCKET_NAME,
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	bin, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := result.Body.Close()
		if err != nil {
			panic(err)
		}
	}()
	return bin, nil
}
