package main

import (
	"fmt"
	"tyk-register/functions"
)

// RouteInfo holds information about each route.
type RouteInfo struct {
	Method     string `json:"method"`
	Route      string `json:"route"`
	MethodName string `json:"method_name"`
	GroupName  string `json:"group_name"`
}

func main() {
	var orgId, gatewayUrl, baseDir string
	fmt.Println("What do you want the organization id to be the microservices (ex: silkworm_ac)?: ")
	fmt.Scanln(&orgId)
	fmt.Println("What is the file path for your organization (folder in which the microservices are, ex: '/Users/nadulj/Documents/Evolza/Silkworm/global_console')")
	fmt.Scanln(&baseDir)
	fmt.Println("What is the gateway url (ex: 'https://tyk-apim.cgaas.ai/tyk/apis' ?: ")
	fmt.Scanln(&gatewayUrl)
	//orgId = "silkworm_gc"
	//gatewayUrl = "http://localhost:8080/tyk/apis"
	functions.Traverse(orgId, gatewayUrl, baseDir)
	//utils.PrintInfo()
}

//func main() {
//	routerFile := `
//	import (
//		"ac_user_management/api"
//		"strings"
//		"github.com/gofiber/fiber/v2"
//		"github.com/gofiber/fiber/v2/middleware/cors"
//		"github.com/gofiber/fiber/v2/middleware/logger"
//		"github.com/gofiber/fiber/v2/middleware/recover"
//	)
//
//	func Router(app *fiber.App) {
//		app.Use(cors.New())
//		app.Use(logger.New())
//		app.Use(recover.New())
//
//		group := app.Group("/ac_user_management/api")
//		check := app.Group("/")
//		app.Use(cors.New(cors.Config{
//			AllowOrigins: "*",
//			AllowMethods: strings.Join([]string{
//				fiber.MethodGet,
//				fiber.MethodPost,
//				fiber.MethodHead,
//				fiber.MethodPut,
//				fiber.MethodDelete,
//				fiber.MethodPatch,
//			}, ","),
//			AllowHeaders: "Origin, Content-Type, Accept",
//		}))
//
//		group.Use(ValidateToken)
//
//		DefaultMappings(check)
//		RouteMappings(group)
//	}
//
//	func RouteMappings(cg fiber.Router) {
//		cg.Post("/verify/email", api.VerifyEmailApi)
//		cg.Post("/resend/otp", api.ResendOTPApi)
//		cg.Post("/login", api.LoginApi)
//		cg.Post("/login/mfa", api.LoginMFAApi)
//		cg.Post("/validate/otp", api.ValidateOtpApi)
//		cg.Get("/find/user", api.FindUserApi)
//		cg.Get("/forgotPassword", api.ForgotPasswordApi)
//		cg.Post("/create/userRecordCP", api.CreateUserRecordCPApi)
//
//		cg.Post("/CreateUser", api.CreateUserApi)
//		cg.Put("/UpdateUser", api.UpdateUserApi)
//		cg.Delete("/DeleteUser", api.DeleteUserApi)
//		cg.Get("/FindallUser", api.FindallUserApi)
//		cg.Get("/FindUsersByEmails", api.FindallUsersByEmailsApi)
//		cg.Post("/changePassword", api.ChangePasswordApi)
//		cg.Post("/validate/changePassword/otp", api.ValidatePasswordChangeOtpApi)
//
//		cg.Post("/CreateUserRole", api.CreateUserRoleApi)
//		cg.Get("/FindallUserRoles", api.FindallUserRolesApi)
//		cg.Delete("/DeleteUserRole", api.DeleteUserRoleApi)
//		cg.Put("/UpdateUserRole", api.UpdateUserRoleApi)
//		cg.Get("/findall/emails", api.FindAllEmailsApi)
//
//		cg.Get("/FindAllIfUsers/walletId", api.FindAllIfUsersByWalletIdApi)
//		cg.Get("/get/user/mobileNumber/userId", api.GetUserMobileNumberApi)
//		cg.Get("/FindUserAnalytics", api.FindUserAnalyticsCPApi)
//		cg.Get("/FindUserCount/summary", api.FindUserCountSummaryApi)
//		cg.Get("/send/mobileNumber/verification/otp", api.SendMobileNumberVerificationOtpApi)
//		cg.Post("/changeMobileNumber/validate/otp", api.ValidateChangeMobileNumberOtpApi)
//		cg.Post("/logout/activity", api.LogoutActivity)
//	}
//
//	func DefaultMappings(cg fiber.Router) {
//		cg.Get("/", func(c *fiber.Ctx) error {
//			return c.JSON(map[string]string{"message": "ac_user_management service is up and running", "version": "1.0"})
//		})
//	}
//	`
//
//	// Regular expression to match group definitions
//	groupRe := regexp.MustCompile(`(\w+)\s*:=\s*app\.Group\("([^"]*)"\)`)
//
//	// Regular expression to match routes and their corresponding methods.
//	routeRe := regexp.MustCompile(`(\w+)\.(\w+)\("([^"]+)",\s*api\.(\w+)\)`)
//
//	// Find all group matches in the routerFile string.
//	groupMatches := groupRe.FindAllStringSubmatch(routerFile, -1)
//
//	// Create a map to associate group variables with group names.
//	groupNames := make(map[string]string)
//	for _, match := range groupMatches {
//		groupVar := match[1]  // Group variable name
//		groupName := match[2] // Group name (route prefix)
//		groupNames[groupVar] = groupName
//	}
//
//	var routes []RouteInfo
//
//	// Find all route matches in the routerFile string.
//	routeMatches := routeRe.FindAllStringSubmatch(routerFile, -1)
//
//	for _, match := range routeMatches {
//		groupVar := match[1]
//		method := match[2]
//		route := match[3]
//		methodName := match[4]
//		groupName, exists := groupNames[groupVar]
//		if !exists {
//			groupName = "unknown" // If a group name is not found
//		}
//
//		routes = append(routes, RouteInfo{
//			Method:     strings.ToUpper(method),
//			Route:      route,
//			MethodName: methodName,
//			GroupName:  groupName,
//		})
//	}
//
//	// Convert routes to JSON.
//	routesJSON, err := json.MarshalIndent(routes, "", "  ")
//	if err != nil {
//		fmt.Printf("Error marshalling routes to JSON: %v\n", err)
//		return
//	}
//
//	// Print the JSON.
//	fmt.Println(string(routesJSON))
//}
