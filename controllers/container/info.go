package container

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/flexphere/sts/lib/aes"
	"github.com/flexphere/sts/lib/response"
	"github.com/flexphere/sts/lib/s3"
	"github.com/gin-gonic/gin"
)

// Info (name, size) of s3 object
func Info(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	req := &struct {
		Body     string `json:"id"`
		Password string `json:"password"`
	}{}
	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(req); err != nil {
		ctx.JSON(response.NewInternalServerError())
		panic(err)
	}

	downloader := s3.New()
	result, err := downloader.Head(req.Body)
	if err != nil {
		ctx.JSON(response.NewNotFound())
		panic(err)
	}

	encKey := req.Password + "00000000"
	encFilenameB64 := strings.Replace(req.Body, "-", "/", -1)
	encFilename, _ := base64.StdEncoding.DecodeString(encFilenameB64)
	filename, err := aes.Decrypt([]byte(encKey), []byte(encFilename))
	if err != nil {
		ctx.JSON(response.NewInternalServerError())
		panic(err)
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(
		200,
		gin.H{
			"name": filename,
			"size": humanize.Bytes(uint64(*result.ContentLength)),
		},
	)
}
