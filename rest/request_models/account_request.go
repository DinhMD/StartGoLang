package request_models

type AccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
