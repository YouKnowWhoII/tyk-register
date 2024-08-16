package dtos

type RouteInfo struct {
	Method     string `json:"method"`
	Route      string `json:"route"`
	MethodName string `json:"method_name"`
	GroupName  string `json:"group_name"`
}

//{POST /logout/activity LogoutActivity /ac_user_management/api}
