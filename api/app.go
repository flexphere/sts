package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"

	"github.com/flexphere/sts/controllers/container"
	"github.com/flexphere/sts/lib/aes"
	"github.com/flexphere/sts/lib/directory"
	"github.com/flexphere/sts/lib/response"
	"github.com/flexphere/sts/lib/s3"
	"github.com/flexphere/sts/lib/storage"
	"github.com/flexphere/sts/settings"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
	settings.InitSettings()
}

func favicon(ctx *gin.Context) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	bin, err := ioutil.ReadFile(fmt.Sprintf("%v/icon/favicon.ico", wd))
	if err != nil {
		panic(err)
	}
	ctx.Writer.Write(bin)
}

func ping(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(
		200,
		response.New(nil),
	)
}

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Kill, os.Interrupt)

	go func() {
		<-signals
		err := os.RemoveAll(*settings.Settings.TMP_DIRECTORY)
		if err != nil {
			panic(err)
		}
		fmt.Println(*settings.Settings.TMP_DIRECTORY)
		os.Exit(0)
	}()

	router := gin.Default()

	router.GET("/favicon.ico", favicon)
	router.GET("/ping", ping)
	router.POST("/ping", ping)
	router.POST("/set", container.Create)
	router.POST("/get", func(ctx *gin.Context) {
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

		var downloader storage.Storage
		if *settings.Settings.S3_BUCKET_NAME == "" {
			downloader = directory.New()
		} else {
			downloader = s3.New()
		}

		encryptedStr, err := downloader.Download(req.Body)
		if err != nil {
			ctx.JSON(response.NewNotFound())
			panic(err)
		}

		plain, err := aes.Decrypt([]byte(req.Password), encryptedStr)
		if err != nil {
			ctx.JSON(response.NewNotFound())
			panic(err)
		}
		ctx.Writer.Write([]byte(plain))
	})

	router.Run(":8080")
}
