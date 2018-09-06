package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	baseURL, err := url.Parse("https://api.openweathermap.org/data/2.5/weather")
	if err != nil {
		log.Fatal(err)
	}
	params := url.Values{}
	params.Add("zip", "81428,us")
	params.Add("appid", "a5708c2191b6429113e3e22e2dcd3a63")
	params.Add("units", "imperial")
	baseURL.RawQuery = params.Encode()
	response, err := http.Get(baseURL.String())
	if err != nil {
		log.Fatal(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Printf(string(data))
	}
}
