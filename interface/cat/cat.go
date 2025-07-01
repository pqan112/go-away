package cat

import (
	"errors"
	"strings"
)

type Cat struct {
	Name string `json:"name"`
}

func New(name string) (*Cat, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if len(name) > 50 {
		return nil, errors.New("name is too long (nax 50 characters)")
	}

	return &Cat{
		Name: name,
	}, nil
}

func (c *Cat) GetName() string {
	return c.Name
}

func (c *Cat) Speak() string {
	return "Mèo meo"
}
