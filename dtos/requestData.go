package dtos

type RequestData struct {
	Name           string
	APIID          string
	OrgID          string
	ListenPath     string
	TargetURL      string
	AllowedMethods []string
	AllowedOrigins []string
}
