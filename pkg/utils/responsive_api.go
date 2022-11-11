package utils

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"err_message"`
	Data       interface{} `json:"data"`
}
