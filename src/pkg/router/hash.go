package router

import (
	"fmt"
	"github.com/spaolacci/murmur3"
)

func murmur3Hasher(href string) (string) {
	hasher := murmur3.New128()
	hasher.Write([]byte(href))
	h1, h2 := hasher.Sum128()
	return fmt.Sprintf("%x%x", h1, h2)
}