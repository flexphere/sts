package settings

import (
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var Settings = &struct {
	AWS_REGION     *string
	S3_BUCKET_NAME *string
	AWS_CONFIG     *aws.Config
	AWS_SESSION    *session.Session
	TMP_DIRECTORY  *string
}{}

func InitSettings() {
	Settings.AWS_REGION = aws.String(os.Getenv("REGION"))
	Settings.S3_BUCKET_NAME = aws.String(os.Getenv("S3_BUCKET_NAME"))
	Settings.AWS_CONFIG = &aws.Config{
		Region: Settings.AWS_REGION,
	}
	Settings.AWS_SESSION = session.New(Settings.AWS_CONFIG)

	tempDir, err := ioutil.TempDir(os.TempDir(), "sts")
	if err != nil {
		panic(err)
	}
	tempDir += "/"

	Settings.TMP_DIRECTORY = &tempDir
}
