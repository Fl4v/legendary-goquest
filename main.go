package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type data struct {
	id              string
	employee_name   string
	employee_salary string
	employee_age    string
	profile_image   string
}

type APIResponse struct {
	status    string
	data_list []data
}

func main() {
	resp, err := http.Get("http://dummy.restapiexample.com/api/v1/employees")

	if err != nil {
		fmt.Println("Oh No")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	s, err := getStations([]byte(body))

	fmt.Println(s)
}

func getStations(body []byte) (*APIResponse, error) {
	var s = new(APIResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
