package container

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/flexphere/sts/lib/directory"
	"github.com/flexphere/sts/lib/key"
	"github.com/flexphere/sts/lib/response"
	"github.com/flexphere/sts/lib/s3"
	"github.com/flexphere/sts/lib/storage"
	"github.com/flexphere/sts/settings"
	"github.com/gin-gonic/gin"
)

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

	encStart := time.Now().UTC()

	k := key.RandStringBytesMaskImprSrc(16)
	// encryptedStr, err := aes.Encrypt([]byte(k), string(req))
	// if err != nil {
	// 	ctx.JSON(response.NewInternalServerError())
	// 	panic(err)
	// }
	key := fmt.Sprintf("%X", sha512.Sum512(req))

	fmt.Println("enc done in: ", time.Since(encStart))

	uploadStart := time.Now().UTC()

	var uploader storage.Storage
	if *settings.Settings.S3_BUCKET_NAME == "" {
		uploader = directory.New()
	} else {
		uploader = s3.New()
	}

	if err := uploader.Upload(key, req); err != nil {
		ctx.JSON(response.NewInternalServerError())
		panic(err)
	}

	fmt.Println("uploaded done in: ", time.Since(uploadStart))

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(
		200,
		response.New(
			map[string]interface{}{
				"password": k,
				"id":       key,
			},
		),
	)
}
