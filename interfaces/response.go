package interfaces

type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}
