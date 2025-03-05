package commons

type Response struct {
	Code    int               `json:"code,omitempty"`
	Message string            `json:"message,omitempty"`
	Result  interface{}       `json:"data,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}
