package model

type ConfigFile struct {
	WatchApi  string `json:"watch_api"`
	NotifyApi string `json:"notify_api"`
	Interval  int    `json:"interval"`
	Verbose   bool   `json:"verbose"`
}
