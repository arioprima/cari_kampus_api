package schemas

type SchemaDatabaseError struct {
	Code int    `json:"code"`
	Type string `json:"type"`
}

type SchemaErrorResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
}

type SchemaUnauthorizedResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
