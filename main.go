package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Common headers
	commonHeaders := func(c *fiber.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Set("Content-Disposition", "inline")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		commonHeaders(c)
		return c.SendString(`{"status":"ok"}`)
	})

	app.Get("/ip", func(c *fiber.Ctx) error {
		commonHeaders(c)
		return c.JSON(fiber.Map{
			"remote_addr":          c.IP(),
			"http_x_forwarded_for": c.Get("X-Forwarded-For"),
			"http_x_real_ip":       c.Get("X-Real-IP"),
		})
	})

	app.Get("/nocache", func(c *fiber.Ctx) error {
		commonHeaders(c)
		c.Set("Cache-Control", "no-cache")
		return c.JSON(fiber.Map{
			"date": time.Now().Format(time.RFC3339),
		})
	})

	app.Get("/cache", func(c *fiber.Ctx) error {
		commonHeaders(c)
		c.Set("Cache-Control", "public, max-age=60")
		return c.JSON(fiber.Map{
			"date": time.Now().Format(time.RFC3339),
		})
	})

	app.Get("/cache-swr", func(c *fiber.Ctx) error {
		commonHeaders(c)
		c.Set("Cache-Control", "public, max-age=60, stale-while-revalidate=300")
		return c.JSON(fiber.Map{
			"date": time.Now().Format(time.RFC3339),
		})
	})

	app.Get("/cdnnocache", func(c *fiber.Ctx) error {
		commonHeaders(c)
		c.Set("CDN-Cache-Control", "no-cache")
		return c.JSON(fiber.Map{
			"date": time.Now().Format(time.RFC3339),
		})
	})

	app.Get("/cdncache", func(c *fiber.Ctx) error {
		commonHeaders(c)
		c.Set("CDN-Cache-Control", "public, max-age=60")
		return c.JSON(fiber.Map{
			"date": time.Now().Format(time.RFC3339),
		})
	})

	app.Get("/cdncache-swr", func(c *fiber.Ctx) error {
		commonHeaders(c)
		c.Set("CDN-Cache-Control", "public, max-age=60, stale-while-revalidate=300")
		return c.JSON(fiber.Map{
			"date": time.Now().Format(time.RFC3339),
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := app.Listen("0.0.0.0:" + port)
	if err != nil {
		log.Fatal(err)
	}
}
