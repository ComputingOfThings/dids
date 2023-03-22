package test

import "embed"

var (
	//go:embed data
	testVector embed.FS
)

func GetTestResource(name string) ([]byte, error) {
	return testVector.ReadFile("data/" + name)
}
