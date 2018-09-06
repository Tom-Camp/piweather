package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	baseURL, err := url.Parse("https://api.openweathermap.org/data/2.5/weather")
	checkErr(err)

	params := url.Values{}
	params.Add("zip", "81428,us")
	params.Add("appid", "a5708c2191b6429113e3e22e2dcd3a63")
	params.Add("units", "imperial")
	baseURL.RawQuery = params.Encode()
	response, err := http.Get(baseURL.String())
	if err != nil {
		panic(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Printf(string(data))
	}
	create()
}

func insert() {
	db, err := sql.Open("mysql", "weatheruser:gr4c!3!2@192.168.1.13/weather")
	checkErr(err)
	defer db.Close()
	_, err = db.Exec("SELECT * FROM wdata WHERE 1")
}

func create() {
	db, err := sql.Open("mysql", "weatheruser:gr4c!3!2@192.168.1.13/weather")
	checkErr(err)
	defer db.Close()
	_, err = db.Exec("CREATE TABLE wdata ( id integer, data varchar(32) )")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
