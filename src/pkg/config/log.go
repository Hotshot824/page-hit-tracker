package config

import (
	"os"
	"io"
	"github.com/gin-gonic/gin"
)

func (c *Config) Initlog () {
	gin.DisableConsoleColor()

    a, err := os.OpenFile("./logs/access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(a, os.Stdout)

    e, err := os.OpenFile("./logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	gin.DefaultErrorWriter = io.MultiWriter(e, os.Stderr)
}