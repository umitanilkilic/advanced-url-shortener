package model

import (
	"time"
)

type ActivityType uint8

const (
	ActivityTypeClick ActivityType = iota
	ActivityTypeVisit
	ActivityTypeRedirect
)

type UrlStats struct {
	UrlID        uint32       `json:"short_url_id"`
	ActivityTime time.Time    `json:"activity_time"`
	ActivityType ActivityType `json:"activity_type"`
	TimeZone     string       `json:"time_zone"`
	IpAddress    string       `json:"ip_address"`
	Location     string       `json:"location"`
	UserAgent    string       `json:"user_agent"`
}
