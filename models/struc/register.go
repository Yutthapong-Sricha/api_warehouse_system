package models

// Binding from JSON
type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

type Login struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Id_branch int    `json:"id_branch" binding:"required"`
}

type LogLogin struct {
	Id_user                int    `json:"id_user" db:"id_user"`
	Staff_fullname         string `json:"staff_fullname" db:"staff_fullname"`
	Tmp_position_name      string `json:"tmp_position_name" db:"tmp_position_name"`
	Id_branch              int    `json:"id_branch" db:"id_branch"`
	Tmp_branch_name        string `json:"tmp_branch_name" db:"tmp_branch_name"`
	Email_login            string `json:"email_login" db:"email_login"`
	Is_success1_fail2_flag int    `json:"is_success1_fail2_flag" db:"is_success1_fail2_flag"`
	Login_time             int    `json:"login_time" db:"login_time"`
	Logout_time            int    `json:"logout_time" db:"logout_time"`
}
