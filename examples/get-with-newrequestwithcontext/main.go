package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-Go-Example", "NewRequestWithContext")

	resp, err := client.Do(req)
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
