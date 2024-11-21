package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/catgir-ls/assets/config"
	"github.com/catgir-ls/assets/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// Initialize Config
	config, err := config.Load("config.toml")

	if err != nil {
		log.Fatalln("Unable to load config")
	}

	logger.Log("Succesfully loaded items into the config!")

	// Variables
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               false,
	})

	// Initialize MinIO
	client, err := minio.New(config.S3.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.S3.AccessKey, config.S3.SecretKey, ""),
		Secure: config.S3.SSL,
	})

	if err != nil {
		logger.Error("Unable to initialize MinIO client!", true)
	}

	logger.Log("Succesfully initialized MinIO client!")

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  200,
			"message": "Welcome to the catgir.ls assets proxy!",
			"data":    fiber.Map{},
		})
	})

	app.Get("/:file", func(c *fiber.Ctx) error {
		file := c.Params("file")

		logger.Log(fmt.Sprintf("Fetching %s", file))

		obj, err := client.GetObject(
			context.Background(),
			config.S3.Bucket,
			file,
			minio.GetObjectOptions{},
		)

		if err != nil {
			logger.Error(fmt.Sprintf("Unable to fetch %s -> %s", file, err.Error()), false)
			return c.SendStatus(404)
		}

		logger.Log(fmt.Sprintf("Fetched %s", file))

		defer obj.Close()

		data, err := io.ReadAll(obj)

		if err != nil {
			logger.Error(fmt.Sprintf("Failed to read %s -> %s", file, err.Error()), false)
			return c.SendStatus(404)
		}

		stat, err := obj.Stat()

		if err != nil {
			logger.Error(fmt.Sprintf("Failed to stat %s -> %s", file, err.Error()), false)
			return c.SendStatus(404)
		}

		c.Set("Content-Type", stat.ContentType)
		return c.Send(data)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.App.Port)))
}
