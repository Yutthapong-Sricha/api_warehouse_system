package models

type Sessions struct {
	Session_id             string `json:"session_id" db:"session_id"`
	Session_start          string `json:"session_start" db:"session_start"`
	Session_last_activity  string `json:"session_last_activity" db:"session_last_activity"`
	Session_ip_address     string `json:"session_ip_address" db:"session_ip_address"`
	Session_user_agent     string `json:"session_user_agent" db:"session_user_agent"`
	Session_data           string `json:"session_data" db:"session_data"`
	Id_user                string `json:"id_user" db:"id_user"`
	Email_login            string `json:"email_login" db:"email_login"`
	Is_success1_fail2_flag string `json:"is_success1_fail2_flag" db:"is_success1_fail2_flag"`
	Login_time             string `json:"login_time" db:"login_time"`
	Logout_time            string `json:"logout_time" db:"logout_time"`
}
