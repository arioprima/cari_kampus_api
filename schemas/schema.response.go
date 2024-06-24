package schemas

type SchemaResponses struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Auth    interface{} `json:"auth,omitempty"`
}
