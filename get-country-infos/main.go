package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type DefaultUrlConstructor interface {
	DefaultUrl() string
}

type CountryChecker interface {
	GetCountryInfos(countryName string)
}

type Countryer interface {
	AddName(name string)
	GetName() string
}

type DefaultHttpUrl struct{}

func (d DefaultHttpUrl) DefaultUrl() string {
	return "https://restcountries.com/v3.1"
}

type Country struct {
	name string
}

func (c *Country) AddName(name string) {
	c.name = name
}

func (c *Country) GetName() string {
	return c.name
}

func (c *Country) GetCountryInfos(countryName string) {
	// Creation of the base URL
	url := DefaultUrlConstructor(DefaultHttpUrl{})
	baseUrl := url.DefaultUrl()

	// path
	path := "/name/"

	// fullUrl
	fullUrl := fmt.Sprintf("%s%s%s", baseUrl, path, countryName)

	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func main() {
	// Add the country's name
	country := Countryer(&Country{})
	country.AddName("brazil")
	countryName := country.GetName()

	// Get the country details
	getCountry := CountryChecker(&Country{})
	getCountry.GetCountryInfos(countryName)

}
