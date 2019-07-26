package gokit

import (
	"testing"
)

type Person struct {
	Name   *string
	Age    int
	School map[string]*string
	Hobbit []string
}

func TestPrettify(t *testing.T) {
	school := make(map[string]*string)
	school["junior"] = String("junior school")
	school["high"] = String("high school")
	person := Person{
		Age:    1,
		Name:   String("jess"),
		School: school,
		Hobbit: []string{"swim", "game"},
	}

	t.Log(Prettify(person))
}

func TestPrettifyJson(t *testing.T) {
	school := make(map[string]*string)
	school["junior"] = String("junior school")
	school["high"] = String("high school")
	person := Person{
		Age:    1,
		Name:   String("jess"),
		School: school,
		Hobbit: []string{"swim", "game"},
	}
	t.Log(PrettifyJson(person, false))

	t.Log(PrettifyJson([]*string{String("book"), String("bed")}, true))
}
