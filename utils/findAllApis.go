package utils

import (
	"fmt"
	"io"
	"net/http"
)

func FindAllAPIs() ([]string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Set("x-tyk-authorization", "foo")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	return []string{string(response)}, err
}
