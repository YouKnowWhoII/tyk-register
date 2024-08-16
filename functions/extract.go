package functions

import (
	"regexp"
	"strings"
)

func Extract(routerFile string) *[]RouteInfo {

	// Regular expression to match routes and their corresponding methods.
	re := regexp.MustCompile(`cg\.(\w+)\("([^"]+)",\s*api\.(\w+)\)`)

	// Find all matches in the routerFile string.
	matches := re.FindAllStringSubmatch(routerFile, -1)

	var routes []RouteInfo

	for _, match := range matches {
		method := match[1]
		route := match[2]
		methodName := match[3]
		groupName := "/ac_user_management/api" // Assuming this is the group for these routes.

		routes = append(routes, RouteInfo{
			Method:     strings.ToUpper(method),
			Route:      route,
			MethodName: methodName,
			GroupName:  groupName,
		})
	}

	// Convert routes to JSON.
	//routesJSON, err := json.MarshalIndent(routes, "", "  ")
	//if err != nil {
	//	fmt.Printf("Error marshalling routes to JSON: %v\n", err)
	//	return
	//}

	// Print the JSON.
	//fmt.Println(string(routesJSON))
	return &routes
}

type RouteInfo struct {
	Method     string `json:"method"`
	Route      string `json:"route"`
	MethodName string `json:"method_name"`
	GroupName  string `json:"group_name"`
}
