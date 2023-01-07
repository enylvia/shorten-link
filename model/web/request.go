package web

import "time"

type URLRequest struct {
	URL         string        `json:"url" validate:"required,url"`
	CustomShort string        `json:"custom_short"`
	ExpiryDate  time.Duration `json:"expiry_date"`
}
