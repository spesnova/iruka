package schema

type VulcandBackend struct {
	Type string `json:"Type"`
}

type VulcandServer struct {
	URL string `json:"URL"`
}

type VulcandFrontend struct {
	Type      string `json:"Type"`
	BackendID string `json:"BackendId"`
	Route     string `json:"Route"`
}
