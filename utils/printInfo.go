package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"tyk-register/dtos"
)

func PrintInfo() {
	// JSON data
	jsonData, err := FindAllAPIs()
	if err != nil {
		log.Fatalf("Error finding all APIs")
	}

	// Convert []string to []byte
	jsonBytes := []byte(jsonData[0])

	// Parse JSON data
	var responses []dtos.APIResponse
	if err := json.Unmarshal(jsonBytes, &responses); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Extract required fields
	for _, response := range responses {
		fmt.Printf("API Name: %s\n", response.Name)
		fmt.Printf("Org ID: %s\n", response.OrgID)
		fmt.Printf("Listen Path: %s\n", response.Proxy.ListenPath)
		fmt.Printf("Target URL: %s\n", response.Proxy.TargetURL)
		fmt.Printf("Allowed Origins: %v\n", response.CORS.AllowedOrigins)
		fmt.Printf("Allowed Methods: %v\n", response.CORS.AllowedMethods)
		fmt.Println()
	}
}
