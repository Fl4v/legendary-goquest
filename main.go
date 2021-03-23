package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type data struct {
	Id              string
	Employee_name   string
	Employee_salary string
	Employee_age    string
	Profile_image   string
}

type all struct {
	Status string
	Data   []data
}

func main() {
	var a all

	response, err := http.Get("http://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		fmt.Println("Oh No")
	}

	defer response.Body.Close()

	fmt.Println(response.Status)

	body, err := io.ReadAll(response.Body)

	err = json.Unmarshal(body, &a)

	if err != nil {
		panic(err)
	}

	fmt.Println(a.Data)
}
