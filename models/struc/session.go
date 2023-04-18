package models

import "time"

type SessionStruc struct {
	Session_id            string `json:"session_id" db:"session_id"`
	Session_start         int    `json:"session_start" db:"session_start"`
	Session_last_activity int    `json:"session_last_activity" db:"session_last_activity"`
	Session_ip_address    string `json:"session_ip_address" db:"session_ip_address"`
	Session_user_agent    string `json:"session_user_agent" db:"session_user_agent"`
	Session_data          string `json:"session_data" db:"session_data"`
}

// each session contains the username of the user and the time at which it expires
type session struct {
	username string
	expiry   time.Time
}
