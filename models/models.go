package models

type ShortUrlModel struct {
	Url  string `json:"url"`
	Uuid string `json:"uuid"`
}

type ReceiverModel struct {
	Url string `json:"url"`
}
