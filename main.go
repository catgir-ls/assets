package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/catgir-ls/assets/utils"
	"github.com/gofiber/fiber/v2"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	err := utils.LoadConfig("config.toml")

	if err != nil {
		log.Fatalln("Unable to load config")
	}

	config := utils.GetConfig()

	app := fiber.New()

	client, err := minio.New(config.S3.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.S3.AccessKey, config.S3.SecretKey, ""),
		Secure: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  200,
			"message": "Welcome to the catgir.ls assets proxy!",
			"data":    fiber.Map{},
		})
	})

	app.Get("/:file", func(c *fiber.Ctx) error {
		obj, err := client.GetObject(context.Background(), config.S3.Bucket, c.Params("file"), minio.GetObjectOptions{})

		if err != nil {
			return c.SendStatus(404)
		}

		defer obj.Close()

		data, err := io.ReadAll(obj)

		if err != nil {
			return c.SendStatus(404)
		}

		return c.Send(data)
	})

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.App.Port)))
}
