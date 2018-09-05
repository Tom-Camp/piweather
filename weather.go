package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?zip=81428,us&appid=a5708c2191b6429113e3e22e2dcd3a63&units=imperial")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Printf(string(data))
	}
}
