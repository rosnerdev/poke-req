package main

import (
	. "fmt"
	. "io"
	"net/http"

	"github.com/bytedance/sonic"
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
	sonic.Unmarshal(resp, &Name)

	Println("This pokemon's name is:", Name.Name)
}
