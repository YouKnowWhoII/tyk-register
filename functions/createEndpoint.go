package functions

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func createEndpoint(gatewayUrl string, json string) error {

	client := &http.Client{}
	var data = strings.NewReader(json)

	req, err := http.NewRequest("POST", gatewayUrl, data)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("x-tyk-authorization", "foo")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}
	fmt.Println(string(bodyText))

	return nil
}
