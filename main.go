package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Replace this with the path to the directory containing your microservices
	baseDir := "/Users/nadulj/Documents/Evolza/Silkworm/admin_console"

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

				fmt.Printf("Contents of %s:\n%s\n", routerPath, content)
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %v: %v", baseDir, err)
	}
}
