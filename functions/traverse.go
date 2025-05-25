package functions

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Traverse(orgId string, gatewayUrl string, baseDir string) {

	//baseDir := BASE_PATH

	// Walk through the directory
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Print("line 18: ", err)
			return err
		}

		// Check if it's a directory
		if info.IsDir() {
			// Construct the path to router.go
			routerPath := filepath.Join(path, "apiHandlers", "router.go")

			// Check if router.go exists
			if _, err := os.Stat(routerPath); err == nil {
				// Ask the user if they want to skip this microservice
				var skip string
				fmt.Printf("Do you want to skip the microservice at %s? (yes/no): ", path)
				fmt.Scanln(&skip)
				if skip == "yes" {
					fmt.Println("Skipping microservice at:", path)
					return nil
				}

				// Read and print the content of router.go
				content, err := ioutil.ReadFile(routerPath)
				if err != nil {
					fmt.Print("line 40:", err)
					return err
				}

				//print the file path currently being explored
				fmt.Println(path)
				var appendix, groupname string
				fmt.Println("What do you want the group name to be for this microservice (ex: '/ac_user_management/api'?: ")
				fmt.Scanln(&groupname)
				fmt.Println("What do you want the appendix to be for this microservice [] (ex: LotteryDrawSchedule-app1954)?: ")
				//take the input from user into a variable called appendix
				fmt.Scanln(&appendix)

				//fmt.Println("What do you want the allowed origins to be for this microservice (ex: https://admin-console.eagleeyes.ai/)? You can provide multiple origins separated by commas:")
				//
				//// Use a buffered reader to allow multiple comma-separated inputs
				//reader := bufio.NewReader(os.Stdin)
				//originsInput, _ := reader.ReadString('\n')
				//
				//// Split the input into a slice of strings and trim any extra spaces
				//allowedOrigins := strings.Split(strings.TrimSpace(originsInput), ",")

				allowedOrigins := []string{"*"}

				//extract the routes from the router.go file
				data := Extract(string(content), groupname)

				_ = Map(orgId, appendix, allowedOrigins, data, gatewayUrl)

				fmt.Println("Completed processing the microservice at: ", path)

			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %v: %v", baseDir, err)
	}
}
