package api

// This file contains classes to represent request parameters against the deconz api

type AcquireKey_Req struct {
	DeviceType string `json:"devicetype"`
	UserName   string `json:"username"`
}
