package container

import (
	"bytes"
	"encoding/base64"
	"io"
	"strings"

	"github.com/flexphere/sts/lib/aes"
	"github.com/flexphere/sts/lib/key"
	"github.com/flexphere/sts/lib/response"
	"github.com/flexphere/sts/lib/s3"
	"github.com/gin-gonic/gin"
)

// Upload file to S3
func Upload(ctx *gin.Context) {
	// get file info
	file, _ := ctx.FormFile("file")

	//get fileData
	fileBuff := new(bytes.Buffer)
	f, err := file.Open()
	if err != nil {
		ctx.JSON(response.NewInternalServerError())
		panic(err)
	}
	io.Copy(fileBuff, f)
	fileData := fileBuff.Bytes()

	// get fileName & password
	password := key.RandStringBytesMaskImprSrc(8)
	encKey := password + "00000000"
	encFilename, err := aes.Encrypt([]byte(encKey), string(file.Filename))
	if err != nil {
		ctx.JSON(response.NewInternalServerError())
		panic(err)
	}
	encFilenameB64 := base64.StdEncoding.EncodeToString([]byte(encFilename))
	encFilenameB64 = strings.Replace(encFilenameB64, "/", "-", -1)

	uploader := s3.New()
	if err := uploader.Upload(encFilenameB64, fileData); err != nil {
		ctx.JSON(response.NewInternalServerError())
		panic(err)
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(
		200,
		response.New(
			map[string]interface{}{
				"id":       encFilenameB64,
				"password": password,
			},
		),
	)
}
