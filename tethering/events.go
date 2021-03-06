package tethering

// Code generated by chromedp-gen. DO NOT EDIT.

// EventAccepted informs that port was successfully bound and got a specified
// connection id.
type EventAccepted struct {
	Port         int64  `json:"port"`         // Port number that was successfully bound.
	ConnectionID string `json:"connectionId"` // Connection id to be used.
}
