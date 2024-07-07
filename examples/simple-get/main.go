package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status got %v", resp.Status))
	}

	fmt.Println(resp.Header.Get("Content-Type"))

	var body struct {
		UserId int `json:"userId"`
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", body)
}
