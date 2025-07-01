package service

type AnimalName interface {
	GetName() string
}

type Animal interface {
	AnimalName
	Speak() string
}
