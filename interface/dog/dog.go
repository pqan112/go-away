package dog

import (
	"errors"
	"strings"
)

type Dog struct {
	Name string `json:"name"`
}

func New(name string) (*Dog, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if len(name) > 50 {
		return nil, errors.New("name is too long (nax 50 characters)")
	}

	return &Dog{
		Name: name,
	}, nil
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) Speak() string {
	return "Gau Gau"
}
