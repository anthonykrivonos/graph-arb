package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config interface {
	// Gemini
	GeminiUrl() string
	GeminiKey() string
	GeminiSecret() string
}

type config struct {
	Config
	geminiUrl string
	geminiKey string
	geminiSecret string
}

func (c *config) GeminiUrl() string {
	return c.geminiUrl
}

func (c *config) GeminiKey() string {
	return c.geminiKey
}

func (c *config) GeminiSecret() string {
	return c.geminiSecret
}

func NewConfig() Config {
	// Load environment variables file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	c := new(config)

	// Retrieve environment variables
	c.geminiUrl = os.Getenv("GEMINI_URL")
	c.geminiKey = os.Getenv("GEMINI_KEY")
	c.geminiSecret = os.Getenv("GEMINI_SECRET")

	return c
}

// Assert config conforms to Config interface
var _ Config = &config{}
