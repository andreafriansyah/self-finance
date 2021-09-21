package shared

// JSONResponse for response api
type JSONResponseValidate struct {
	Code    int                      `json:"code"`
	Status  string                   `json:"status"`
	Message []map[string]interface{} `json:"message"`
}

type JSONResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// JsonResponse for checking employee data on project (11-02-21)
type JSONResponseExist struct {
	Code        int      `json:"code"`
	Status      string   `json:"status"`
	Message     string   `json:"message"`
	IsExist     bool     `json:"is_exist"`
	ProjectData []string `json:"project_data"`
}
