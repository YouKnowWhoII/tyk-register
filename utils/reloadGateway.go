package utils

import (
	"fmt"
	"io"
	"net/http"
)

func ReloadGateway(gatewayurl string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", gatewayurl+"/tyk/reload/group", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Set("x-tyk-authorization", "foo")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}
	fmt.Println("Gateway reloaded")
}
