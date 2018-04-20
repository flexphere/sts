package container

import (
	"bytes"
	"encoding/base64"
	"io"
	"strings"

	"github.com/flexphere/sts/lib/aes"
	"github.com/flexphere/sts/lib/response"
	"github.com/flexphere/sts/lib/s3"
	"github.com/gin-gonic/gin"
)

// Download object from s3
func Download(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	id := ctx.PostForm("id")
	password := ctx.PostForm("password")

	downloader := s3.New()
	result, err := downloader.Download(id)
	if err != nil {
		ctx.JSON(response.NewNotFound())
		panic(err)
	}

	encKey := password + "00000000"
	encFilenameB64 := strings.Replace(id, "-", "/", -1)
	encFilename, _ := base64.StdEncoding.DecodeString(encFilenameB64)
	filename, err := aes.Decrypt([]byte(encKey), []byte(encFilename))
	if err != nil {
		ctx.JSON(response.NewInternalServerError())
		panic(err)
	}

	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename=\""+filename+"\"")
	ctx.Header("Content-Length", string(len(result)))
	io.Copy(ctx.Writer, bytes.NewReader(result))
}
