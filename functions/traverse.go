package functions

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Traverse() {

	baseDir := BASE_PATH

	// Walk through the directory
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a directory
		if info.IsDir() {
			// Construct the path to router.go
			routerPath := filepath.Join(path, "apiHandlers", "router.go")

			// Check if router.go exists
			if _, err := os.Stat(routerPath); err == nil {
				// Read and print the content of router.go
				content, err := ioutil.ReadFile(routerPath)
				if err != nil {
					return err
				}

				//print the file path currently being explored
				fmt.Println(path)
				var appendix, orgId string
				fmt.Println("What do you want the appendix to be for this microservice (ex: accvm)?: ")
				//take the input from user into a variable called appendix
				fmt.Scanln(&appendix)
				fmt.Println("What do you want the organization id to be this microservice (ex: silkworm_ac)?: ")
				fmt.Scanln(&orgId)

				fmt.Println("What do you want the allowed origins to be for this microservice (ex: https://admin-console.eagleeyes.ai/)? You can provide multiple origins separated by commas:")

				// Use a buffered reader to allow multiple comma-separated inputs
				reader := bufio.NewReader(os.Stdin)
				originsInput, _ := reader.ReadString('\n')

				// Split the input into a slice of strings and trim any extra spaces
				allowedOrigins := strings.Split(strings.TrimSpace(originsInput), ",")

				//extract the routes from the router.go file
				data := Extract(string(content))

				_ = Map(orgId, appendix, allowedOrigins, data)
				fmt.Println("Successfully mapped the routes")
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %v: %v", baseDir, err)
	}
}
