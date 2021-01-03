package model

type PostModel struct {
	Url      string      `json:"url,omitempty"`
	DateTime string      `json:"date_time,omitempty"`
	Status   *NotifyData `json:"status,omitempty"`
	Message  string      `json:"message,omitempty"`
}
