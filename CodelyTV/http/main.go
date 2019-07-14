package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Beer struct {
	Id          int
	Name        string
	Tagline     string
	FirstBrewed string `json:"first_brewed"`
	Description string
	Image_url   string
}

func main() {
	fmt.Println("hello")

	allBeersResponse, err := http.Get("https://api.punkapi.com/v2/beers")
	if err != nil {
		fmt.Errorf(err.Error(), nil)
	}

	fmt.Println(allBeersResponse)

	contents, err := ioutil.ReadAll(allBeersResponse.Body)
	if err != nil {
		fmt.Errorf(err.Error(), nil)
	}

	fmt.Println(contents)
	var beers []Beer
	err = json.Unmarshal(contents, &beers)

	fmt.Println(beers)

	fmt.Println("bye")
}
