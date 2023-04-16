package hello

import (
	"error"
)

func Hello(name String) (string, error) {
	if name == "" {
		return name, error.New("empty name")
	}
}
