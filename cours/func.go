package cours

import "fmt"

func world(name string) (string, error) {

	if name == "" {
		return name, fmt.Errorf("Aucun name est passer en argument")
	}

	return name, nil
}

func Hello(name string) (string, error) {
	/*
		var  str string
		var err error
		str, err = world(name)
	*/

	if name == "" {
		name = "World"
	}

	return world(name)
}
