package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type Config interface {
	// Gemini
	GeminiUrl() string
	GeminiKey() string
	GeminiSecret() string
	// Binance
	BinanceUrl() string
	BinanceKey() string
	BinanceSecret() string
}

type config struct {
	Config
	// Gemini
	geminiUrl string
	geminiKey string
	geminiSecret string
	// Binance
	binanceUrl string
	binanceKey string
	binanceSecret string
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

func (c *config) BinanceUrl() string {
	return c.binanceUrl
}

func (c *config) BinanceKey() string {
	return c.binanceKey
}

func (c *config) BinanceSecret() string {
	return c.binanceSecret
}

func NewConfig() Config {
	// Load environment variables file
	err := godotenv.Load(path.Join(rootDir(), ".env"))
	if err != nil {
		log.Fatal(err)
	}
	c := new(config)

	// Retrieve environment variables
	c.geminiUrl = os.Getenv("GEMINI_URL")
	c.geminiKey = os.Getenv("GEMINI_KEY")
	c.geminiSecret = os.Getenv("GEMINI_SECRET")
	c.binanceUrl = os.Getenv("BINANCE_URL")
	c.binanceKey = os.Getenv("BINANCE_KEY")
	c.binanceSecret = os.Getenv("BINANCE_SECRET")

	return c
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

// Assert config conforms to Config interface
var _ Config = &config{}
