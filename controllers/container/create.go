package container

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"

	"github.com/flexphere/sts/lib/key"
	"github.com/flexphere/sts/lib/response"
	"github.com/flexphere/sts/lib/s3"
	"github.com/gin-gonic/gin"
)

// Create file to S3
func Create(ctx *gin.Context) {
	req, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(response.NewInternalServerError())
		panic(err)
	}
	defer func() {
		err := ctx.Request.Body.Close()
		if err == nil {
			return
		}
		ctx.JSON(response.NewInternalServerError())
		panic(err)
	}()

	password := key.RandStringBytesMaskImprSrc(8)
	key := fmt.Sprintf("%X", sha512.Sum512(req))

	uploader := s3.New()
	if err := uploader.Upload(key, req); err != nil {
		ctx.JSON(response.NewInternalServerError())
		panic(err)
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(
		200,
		response.New(
			map[string]interface{}{
				"id":       key,
				"password": password,
			},
		),
	)
}
