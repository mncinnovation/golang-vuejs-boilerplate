package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/buger/jsonparser"
	"github.com/labstack/echo/v4"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

type renderData struct {
	title string
	data  interface{}
}

func MainPage(c echo.Context) error {
	resp := new(http.Response)
	req, err := http.NewRequest("GET", "http://localhost:1323/frameworks", nil)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	body := string(bodyBytes)

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		fmt.Println(err)
		return c.Render(resp.StatusCode, "error", body)
	}

	val, _, _, err := jsonparser.Get([]byte(body))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	var frameworks []map[string]interface{}
	err = json.Unmarshal(val, &frameworks)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(val)

	fmt.Println(frameworks)
	data := make(echo.Map)
	data["title"] = "Golang + Vue.js Boilerplate"
	data["frameworks"] = frameworks
	// for _, dt := range bodyBytes
	return c.Render(resp.StatusCode, "main", data)
}
