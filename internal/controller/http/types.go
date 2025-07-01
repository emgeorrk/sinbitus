package http

import "time"

type EchoRequest struct {
	Message string `json:"message" validate:"required"`
}

type EchoResponse struct {
	Message string    `json:"message"`
	Echoed  bool      `json:"echoed"`
	Time    time.Time `json:"time"`
}

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
}

type ErrorResponse struct {
	Error   string    `json:"error"`
	Code    int       `json:"code"`
	Time    time.Time `json:"time"`
	Request string    `json:"request_id,omitempty"`
}

type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Time    time.Time   `json:"time"`
}
