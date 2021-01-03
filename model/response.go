package model

type ResponseData struct {
	StatusCode int
	Error      error
	Data       []byte
}

type NotifyData struct {
	StatusCode int
	Error      error
	Data       string
}
