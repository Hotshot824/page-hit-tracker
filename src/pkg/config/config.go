package config

import (
	"os"
	"github.com/goccy/go-json"
)

// Public struct for reading JSON
type readJson struct {
	TrackerURL string `json:"TrackerURL"`
}

// Private struct for config
type Config  struct {
	trackerURL string
}

func NewConfig() *Config  {
	c := &Config {}
	err := c.load("./config.json")
	if err != nil {
		panic(err)
	}
	return c
}

func (c *Config) load(path string) error {
	file, err := os.Open("./config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Read JSON
	var read readJson
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&read); err != nil {
		return err
	}

	c.trackerURL = read.TrackerURL
	return nil
}

func (c *Config ) GetURL() string {
	return c.trackerURL
}