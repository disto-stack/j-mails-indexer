package services

import "os"


type Config	struct {
	zincsearchUrl	string
}

func (c *Config) GetZincsearchUrl() string  {
	return c.zincsearchUrl;
}

func (c *Config) SetUrlsFromEnv() {
	zincsearchUrlFromEnv, existsZincsearchUrlFromEnv := os.LookupEnv("ZINCSEARCH_URL")

	if !existsZincsearchUrlFromEnv {
		zincsearchUrlFromEnv = "http://localhost:4080"
	}

	c.zincsearchUrl = zincsearchUrlFromEnv;
}