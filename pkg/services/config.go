package services

import "os"

type ConfigService struct {
	zincsearchUrl string
}

func (c *ConfigService) GetZincsearchUrl() string {
	return c.zincsearchUrl
}

func (c *ConfigService) SetUrlsFromEnv() {
	zincsearchUrlFromEnv, existsZincsearchUrlFromEnv := os.LookupEnv("ZINCSEARCH_URL")

	if !existsZincsearchUrlFromEnv {
		zincsearchUrlFromEnv = "http://localhost:4080"
	}

	c.zincsearchUrl = zincsearchUrlFromEnv
}
