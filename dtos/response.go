package dtos

type APIResponse struct {
	Name  string `json:"name"`
	OrgID string `json:"org_id"`
	Proxy Proxy  `json:"proxy"`
	CORS  CORS   `json:"CORS"`
}

type Proxy struct {
	ListenPath string `json:"listen_path"`
	TargetURL  string `json:"target_url"`
}

type CORS struct {
	Enable         bool     `json:"enable"`
	AllowedOrigins []string `json:"allowed_origins"`
	AllowedMethods []string `json:"allowed_methods"`
}
