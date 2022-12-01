package main

import (
	"encoding/json"
	. "fmt"
	. "io"
	"net/http"
)

type Name struct {
	Name string `json:"name"`
}

func main() {
	req, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		Println("There was an error")
	}

	defer req.Body.Close()

	resp, _ := ReadAll(req.Body)

	Name := Name{}
	json.Unmarshal(resp, &Name)

	Printf("This pokemon's name is: %s\n", Name.Name)
}
