package schema

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type NewCallbackRequest struct {
	Name string `json:"name"`
}

type NewCallbackResponse struct {
	Response
	Name string `json:"name"`
}
