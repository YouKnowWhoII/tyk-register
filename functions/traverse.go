package functions

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

				data := Extract(string(content))

				fmt.Println(*data)
				fmt.Println("--------------------------------------------------")
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %v: %v", baseDir, err)
	}
}
