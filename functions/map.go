package functions

import (
	"fmt"
	"tyk-register/dtos"
)

func Map(orgId string, appendix string, allowedOrigins []string, routes *[]RouteInfo) error {
	for _, route := range *routes {
		requestData := dtos.RequestData{
			Name:           appendix + "_" + route.MethodName,
			APIID:          appendix + "_" + route.MethodName,
			OrgID:          orgId,
			AllowedMethods: []string{route.Method},
			ListenPath:     "/" + appendix + "/api/v1" + route.Route,
			TargetURL:      "https://" + appendix + ".eagleeyes.ai" + route.GroupName + route.Route,
			AllowedOrigins: allowedOrigins,
		}

		//add a unique check for api id
		println(fmt.Sprintf("Request Data: %v", requestData))
	}
	return nil
}
