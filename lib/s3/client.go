package s3

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/flexphere/sts/settings"
)

type S3Client struct {
	svc s3iface.S3API
}

func New() *S3Client {
	return &S3Client{
		svc: s3.New(settings.Settings.AWS_SESSION, settings.Settings.AWS_CONFIG),
	}
}
