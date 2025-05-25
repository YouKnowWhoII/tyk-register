package functions

import (
	"fmt"
	"strings"
	"text/template"
	"tyk-register/dtos"
	"tyk-register/utils"
)

func Map(orgId string, appendix string, allowedOrigins []string, routes *[]dtos.RouteInfo, gatewayUrl string) error {
	for _, route := range *routes {
		requestData := dtos.RequestData{
			Name:           appendix + "_" + route.MethodName,
			APIID:          appendix + "_" + route.MethodName,
			OrgID:          orgId,
			AllowedMethods: []string{route.Method},
			ListenPath:     "/gateway/" + appendix + route.Route,
			TargetURL:      "https://" + appendix + ".demo.cgaas.ai" + route.GroupName + route.Route,
			AllowedOrigins: allowedOrigins,
		}

		//add a unique check for api id
		// utils.UniqueCheck(requestData.APIID)

		//parse to template here
		tmpl, err := template.New("request").Funcs(template.FuncMap{
			"join": strings.Join,
		}).ParseFiles("templates/request.template")
		if err != nil {
			fmt.Println("line 31", err)
			return err
		}

		buf := new(strings.Builder)

		err = tmpl.ExecuteTemplate(buf, "request.template", requestData)
		if err != nil {
			fmt.Println("line 39", err)
			return err
		}

		_ = createEndpoint(gatewayUrl, buf.String())
	}
	fmt.Println("Successfully registered the microservice with Tyk!")
	utils.ReloadGateway(gatewayUrl)
	return nil
}
