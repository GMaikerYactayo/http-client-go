package model

type GeneralResponse struct {
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
}

type LoginResponse struct {
	GeneralResponse
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}
