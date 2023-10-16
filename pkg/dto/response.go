package dto

type DefaultResponse struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Id      int64       `json:"id,omitempty"`
}
