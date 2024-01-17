package services

import (
	"os"
)

type ConfigService struct {
	zincsearchUrl      string
	zincSearchUser     string
	zincSearchPassword string
}

func (c *ConfigService) GetZincsearchUrl() string {
	return c.zincsearchUrl
}

func (c *ConfigService) SetUrlsFromEnv() {
	zincsearchUrlFromEnv, existsZincsearchUrlFromEnv := os.LookupEnv("ZINCSEARCH_URL")
	zincSearchUser, existsZincsearchUser := os.LookupEnv("ZINCSEARCH_USER")
	zincSearchPassword, existsZincsearchPassword := os.LookupEnv("ZINCSEARCH_PASSWORD")

	if !existsZincsearchUrlFromEnv {
		zincsearchUrlFromEnv = "http://localhost:4080"
	}

	if !existsZincsearchUser {
		zincSearchUser = ""
	}

	if !existsZincsearchPassword {
		zincSearchPassword = ""
	}

	c.zincsearchUrl = zincsearchUrlFromEnv
	c.zincSearchUser = zincSearchUser
	c.zincSearchPassword = zincSearchPassword
}
