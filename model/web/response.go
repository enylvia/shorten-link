package web

import "time"

type URLResponse struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"custom_short"`
	ExpiryDate     time.Duration `json:"expiry_date"`
	XRateRemaining int           `json:"x_rate_remaining"`
	XrateLimitRest time.Duration `json:"x_rate_limit_rest"`
}
