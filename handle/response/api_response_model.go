package response

type APIError struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
}

func (e *APIError) Error() string {
	return e.Message
}

type APIResponse struct {
	Message any `json:"message"`
}
