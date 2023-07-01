package config

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		EnableTrustedProxyCheck: true,
		ReadTimeout:             time.Second * time.Duration(readTimeoutSecondsCount),
		JSONEncoder:             json.Marshal,
		JSONDecoder:             json.Unmarshal,
	}
}
