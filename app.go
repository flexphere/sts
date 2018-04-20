package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"

	"github.com/flexphere/sts/controllers/container"
	"github.com/flexphere/sts/lib/response"
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
	router.POST("/info", container.Info)
	router.POST("/download", container.Download)
	router.POST("/upload", container.Upload)

	router.Run(":8080")
}
