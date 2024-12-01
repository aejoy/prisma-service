package models

import "time"

type Photo struct {
	ID        string    `json:"id,omitempty"`
	Creator   string    `json:"creator,omitempty"`
	To        string    `json:"to,omitempty"`
	URL       string    `json:"url,omitempty"`
	BlurHash  string    `json:"blur_hash,omitempty"`
	Height    int       `json:"height,omitempty"`
	Width     int       `json:"width,omitempty"`
	Size      int       `json:"size,omitempty"`
	SizeInKB  float32   `json:"size_in_kb,omitempty"`
	SizeInMB  float32   `json:"size_in_mb,omitempty"`
	Published time.Time `json:"published,omitempty"`
	Updated   time.Time `json:"updated,omitempty"`
	Archived  time.Time `json:"archived,omitempty"`
}
