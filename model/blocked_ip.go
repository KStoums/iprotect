package model

import "time"

type BlockedIp struct {
	Ip        string    `json:"ip"`
	BlockedAt time.Time `json:"blockedAt"`
}
